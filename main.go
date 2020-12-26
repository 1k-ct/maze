package main

import (
	"fmt"

	"github.com/1k-ct/maze/board"
	"github.com/1k-ct/maze/dig"
)

func main() {
	// s := board.Size{Width: 31, Height: 31}
	size := 37
	wall := 1
	aisle := 0
	s := board.New(size, wall, aisle)
	m := s.MakeField()
	dig.DiggingFind(m, []int{2, 2})
	res := dig.ShowField(m)
	for _, v := range res {
		fmt.Println(v)
	}
}
