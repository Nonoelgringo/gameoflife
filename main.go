package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Cell represents a cell in the game of life
// aliveNext is used as a buffer to store the next cell's status
type Cell struct {
	alive          bool
	aliveNext      bool
	aliveNeighbors int
}

func main() {
	// number of milliseconds between prints
	var millisecondsNb int = 150
	var filename string
	args := len(os.Args)
	switch {
	case args <= 1 || args > 3:
		fmt.Printf("1 or 2 arguments expected.Exiting.")
		os.Exit(1)
	case args == 2:
		filename = os.Args[1]
		fallthrough
	case args <= 3:
		filename = os.Args[1]
		millisecondsNb, _ = strconv.Atoi(os.Args[2])
	}

	clear, err := exec.Command("/usr/bin/clear").Output()
	if err != nil {
		log.Fatal(err)
	}

	seed, err := seedFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	cellsGrid := initializeCells(seed)
	_, x := sizecellsGrid(cellsGrid)

	for {
		if aliveCells := aliveCells(cellsGrid); aliveCells < 5 {
			break
		}
		fmt.Println(strings.Repeat("- ", x))
		printCells(cellsGrid)
		fmt.Println(strings.Repeat("- ", x))
		computeStates(cellsGrid)
		time.Sleep(time.Duration(millisecondsNb) * time.Millisecond)
		os.Stdout.Write(clear)
	}
}

// seedFromFile returns the initial seed as a *[][]string from a file
func seedFromFile(filename string) (*[][]string, error) {

	seed := make([][]string, 0)
	for i := range seed {
		seed[i] = make([]string, 0)
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error during the opening of file : '%v'", filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seed = append(seed, strings.Split(scanner.Text(), " "))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error encountered by the Scanner")
	}

	return &seed, nil
}

func printCells(cellArray *[][]Cell) {
	for col := range *cellArray {
		for row := range (*cellArray)[col] {
			if (*cellArray)[col][row].alive == true {
				fmt.Printf("0 ")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Print("\n")
	}
}

func initializeCells(seed *[][]string) *[][]Cell {
	// retrieve size from seed
	y, x := sizeStringGrid(seed)

	// initialize cell grid using seed size
	cellsGrid := make([][]Cell, y)
	for i := range cellsGrid {
		cellsGrid[i] = make([]Cell, x)
	}

	// Temp variables
	var alive bool
	var aliveNeighbors int
	// loop through seed
	for col := range *seed {
		for row := range (*seed)[col] {
			// if not '-' in seed -> Cell is alive
			if (*seed)[col][row] != "-" {
				alive = true
			} else {
				alive = false
			}
			// counts the number of alive neighbors around the Cell
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					// if not out of bounds
					if (col+i >= 0 && row+j >= 0) && (col+i < y && row+j < x) {
						// if its not dead and its not the Cell itself -> its an alive neighbor
						if (*seed)[col+i][row+j] != "-" && !(i == 0 && j == 0) {
							aliveNeighbors++
						}
					}
				}
			}
			// Add Cell to the grid
			cellsGrid[col][row] = Cell{alive, false, aliveNeighbors}
			// reset aliveNeighbors
			aliveNeighbors = 0
		}
	}
	return &cellsGrid
}

// computeStates computes Cells Grid next state
func computeStates(grid *[][]Cell) {
	y, x := sizecellsGrid(grid)
	var aliveNeighbors int
	// compute Cells states
	for col := range *grid {
		for row := range (*grid)[col] {
			cell := &(*grid)[col][row]
			// if Cell is alive and has neighbors < 1 => dead
			// else if Cell is dead and has neighbors > 4 => born(alive)
			// else it keeps its actual state
			if (*cell).isAlive() && ((*cell).aliveNeighbors <= 1 || (*cell).aliveNeighbors >= 4) {
				(*cell).aliveNext = false
			} else if !(*cell).isAlive() && ((*cell).aliveNeighbors == 3) {
				(*cell).aliveNext = true
			} else {
				(*cell).aliveNext = (*cell).alive
			}
		}
	}
	// compute alive neighbors
	for col := range *grid {
		for row := range (*grid)[col] {
			// counts the number of alive neighbors around the Cell
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					// Boundaries
					if (col+i >= 0 && row+j >= 0) && (col+i < y && row+j < x) {
						// if its not dead and its not the Cell itself -> its an alive neighbor
						if (*grid)[col+i][row+j].isAliveNext() && !(i == 0 && j == 0) {
							aliveNeighbors++
						}
					}
				}
			}
			(*grid)[col][row].aliveNeighbors = aliveNeighbors
			// put aliveNext into alive for printing as the final computes Cell state
			(*grid)[col][row].alive = (*grid)[col][row].aliveNext
			aliveNeighbors = 0
		}
	}
}

// Tool functions
func sizecellsGrid(p *[][]Cell) (y int, x int) {
	for _, h := range *p {
		x = len(h)
		break
	}
	y = len(*p)

	return y, x
}

func sizeStringGrid(p *[][]string) (y int, x int) {
	for _, h := range *p {
		x = len(h)
		break
	}
	y = len(*p)

	return y, x
}

func (c Cell) isAlive() bool {
	return c.alive
}

func (c Cell) isAliveNext() bool {
	return c.aliveNext
}

func aliveCells(p *[][]Cell) int {
	aliveCells := 0
	for col := range *p {
		for row := range (*p)[col] {
			if (*p)[col][row].isAlive() {
				aliveCells++
			}
		}
	}
	return aliveCells
}
