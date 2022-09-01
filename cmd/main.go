package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/chzyer/readline"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/urfave/cli/v2"

	"github.com/st0rrer/proxx/internal"
)

func main() {
	app := &cli.App{
		Name:  "proxx",
		Usage: "PROXX — a game of proximity",
		Flags: []cli.Flag{
			&cli.UintFlag{
				Name:    "board_size",
				Usage:   "Build a board within size (SxS)",
				Aliases: []string{"s"},
				Value:   10,
			},
			&cli.IntFlag{
				Name:    "bombs",
				Usage:   "Number of bombs",
				Aliases: []string{"b"},
				Value:   5,
			},
		},
		Action: func(ctx *cli.Context) error {
			rand.Seed(time.Now().UnixNano())
			log.Println("Let the games begin!")

			boardSize := ctx.Int("board_size")
			bombs := ctx.Int("bombs")

			err := validation.Errors{
				"board_size": validation.Validate(boardSize, validation.Min(5), validation.Max(20)),
				"bombs":      validation.Validate(bombs, validation.Min(1), validation.Max(boardSize*boardSize-9)),
			}.Filter()
			if err != nil {
				log.Printf("ivnvalid arguments: %s", err)
				return nil
			}

			board := internal.NewBoard(boardSize, bombs)

			printBoard := func(openAll bool) error {
				if _, err := fmt.Fprintf(os.Stdout, "\033[2J"); err != nil {
					return fmt.Errorf("failed to clear terminal window: %s", err)
				}
				encodedBoard, err := internal.NewEncoder(openAll).Encode(board)
				if err != nil {
					return fmt.Errorf("failed to encode board: %w", err)
				}

				if _, err := fmt.Fprintf(os.Stdout, encodedBoard); err != nil {
					return fmt.Errorf("failed to print board: %w", err)
				}
				return nil
			}
			if err := printBoard(false); err != nil {
				return err
			}

			proxxTerminal := &cli.App{
				CommandNotFound: func(context *cli.Context, s string) {
					log.Printf("Command `%s` does not exist. Please try `help` command", s)
				},
				Commands: []*cli.Command{
					{
						Name:  "open",
						Usage: "Open the square within the x and y coordinates",
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:     "x",
								Usage:    "x axis",
								Required: true,
							},
							&cli.IntFlag{
								Name:     "y",
								Usage:    "y axis",
								Required: true,
							},
						},
						Action: func(ctx *cli.Context) error {
							x := ctx.Int("x")
							y := ctx.Int("y")

							err := validation.Errors{
								"x": validation.Validate(x, validation.Min(0), validation.Max(boardSize).Exclusive()),
								"y": validation.Validate(y, validation.Min(0), validation.Max(boardSize).Exclusive()),
							}.Filter()
							if err != nil {
								log.Printf("%s: ivnvalid arguments: %s", ctx.Command.Name, err)
								return nil
							}

							if isBlackHole := board.IsBomb(x, y); isBlackHole {
								if err := printBoard(true); err != nil {
									log.Printf("failed to print board: %v", err)
								}
								return fmt.Errorf("game over {%d;%d} is BOMB", x, y)
							}
							remainingCells := board.Open(x, y)

							if remainingCells != 0 {
								return printBoard(false)
							}
							if err := printBoard(true); err != nil {
								return err
							}
							if _, err := fmt.Fprintln(os.Stdout, "WIN WIN"); err != nil {
								return err
							}
							os.Exit(0)
							return nil
						},
					},
				},
			}

			rl, err := readline.NewEx(&readline.Config{
				Prompt: "> ",
			})
			if err != nil {
				return err
			}
			defer rl.Close()

			for {
				line, err := rl.Readline()
				if err == io.EOF {
					return nil
				}
				if err != nil {
					return err
				}

				rl.SaveHistory(line)
				if err := proxxTerminal.Run(strings.Fields("cmd " + line)); err != nil {
					return err
				}
			}
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
