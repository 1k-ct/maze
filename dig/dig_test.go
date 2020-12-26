package dig

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/1k-ct/maze/board"
)

// func TestDigging(t *testing.T) {
// 	type args struct {
// 		m  *board.Map
// 		sp []int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			Digging(tt.args.m, tt.args.sp)
// 		})
// 	}
// }

func TestDiggingFind(t *testing.T) {
	s := board.Size{Width: 31, Height: 31}
	m := s.MakeField()
	DiggingFind(m, []int{2, 2})
	res := ShowField(m)
	for _, v := range res {
		fmt.Println(v)
	}
}
func TestRoadRand(t *testing.T) {
	// s := board.Size{Width: 20, Height: 15}
	// m := s.MakeField()
	// n, l := roadRand(m)
}
func TestR(t *testing.T) {
	for i := 0; i < 20; i++ {
		n := rand.Intn(5)
		fmt.Println(n)
	}
}
func TestShowField(t *testing.T) {
	s := board.Size{Width: 20, Height: 15}
	m := s.MakeField()
	res := ShowField(m)
	for _, v := range res {
		fmt.Println(v)
	}
}
