# Proxx

##### Parts

- [Part I](#part_1)
- [Part II](#part_2)
- [Part III & IV](#part_3_4)
- [Launch](#launch)

## Part I <a name="part_1"/>

To store squares, I created a two-dimensional array.
Each element of array has the following type: `type Cell byte`.
This allows us to use fewer memory allocations than for example if we would use `int`.

Also, there are several predefined `Cell` const:

* `OpenCell` - `byte(9)`
* `ClosedCell` - `byte(17)`
* `BombCell` - `byte(33)`
* `OpenBombCell` - `byte(41)`

If the square has adjacent bombs, we count all surrounding bombs and assign value to the square cell.
The value must be less or equal `8`. If there aren't any adjacent bombs, just set `OpenCell`.

## Part II <a name="part_2"/>

User can set a size of board and a number of bombs.
Number of the bombs should be less than `board_size^2 - 9`

To uniform distribution, I created a set of random numbers by `rand.Intn` with range `[0, board_size^2)`.
And then for each number of the set I transformed to two-dimensional array coordinates due to the following formula:

```
bombX := rand_number / board_size
bombY := rand_number % board_size
```

## Part III <a name="part_3_4"/>

For traversal across all the adjacent squares I used the `Breadth-First search`. For this purpose I implemented the
queue:

```go 
// internal/queue.go
type Queue struct {
	arr []int
}
```

## Launch <a name="launch"/>

### CLI Arguments

```NAME:
proxx - PROXX — a game of proximity

USAGE:
   proxx [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --board_size value, -s value  Build a board within size (SxS) (default: 10)
   --bombs value, -b value       Number of bombs (default: 5)
   --help, -h                    show help (default: false)
```

### Basic usage

```
make build
./bin/proxx -s 10 -b 10
```

### Docker usage

```
IMAGE_NAME=proxx make docker_build
docker run -ti proxx:latest /app/proxx -s 10 -b 10
```

### SubCommands
```
NAME:
   open - Open the square within the x and y coordinates

USAGE:
   open [command options] [arguments...]

OPTIONS:
   -x value  x axis (default: 0)
   -y value  y axis (default: 0)
```

### Game UX/UI

- `*` - Closed square  
- `#` - Open square, without any adjacent bombs  
- `[1,8]` - Open square with the count of the adjacent bombs
- `B` - Bomb
