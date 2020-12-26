package dig

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/1k-ct/maze/board"
)

const wall = 1
const aisle = 5 // 通路
const none = 0

func DiggingFind(m *board.Map, sp []int) {
	m.Maze[sp[0]][sp[1]] = aisle
	var up, right, down, left int = 0, 1, 2, 3

	// rand.Seed(rand.Int63())
	rand.Seed(time.Now().UnixNano())
	stop := true
	for stop {
		for _, v := range rand.Perm(4) {
			switch v {
			case up:
				if m.Maze[sp[0]-1][sp[1]] != 1 && m.Maze[sp[0]-2][sp[1]] == 0 {
					m.Maze[sp[0]-1][sp[1]], m.Maze[sp[0]-2][sp[1]] = aisle, aisle
					sp[0] = sp[0] - 2
				}
			case right:
				if m.Maze[sp[0]][sp[1]+1] != 1 && m.Maze[sp[0]][sp[1]+2] == 0 {
					m.Maze[sp[0]][sp[1]+1], m.Maze[sp[0]][sp[1]+2] = aisle, aisle
					sp[1] = sp[1] + 2
				}
			case down:
				if m.Maze[sp[0]+1][sp[1]] != 1 && m.Maze[sp[0]+2][sp[1]] == 0 {
					m.Maze[sp[0]+1][sp[1]], m.Maze[sp[0]+2][sp[1]] = aisle, aisle
					sp[0] = sp[0] + 2
				}
			case left:
				if m.Maze[sp[0]][sp[1]-1] != 1 && m.Maze[sp[0]][sp[1]-2] == 0 {
					m.Maze[sp[0]][sp[1]-1], m.Maze[sp[0]][sp[1]-2] = aisle, aisle
					sp[1] = sp[1] - 2
				}
			default:
				fmt.Println("error")
			}
		}
		b, err := noPassage(m, sp)
		if err != nil {
			fmt.Printf("%v", err)
		}
		if b {
			n, l := roadRand(m)
			if !l {
				stop = false
			}
			sp[0], sp[1] = n[0], n[1]
		}
	}
}

// 行き止まりtrue 通路ありfalse
func noPassage(m *board.Map, sp []int) (bool, error) {
	for i, v := range [][]int{[]int{-2, 0}, []int{0, 2}, []int{2, 0}, []int{0, -2}} {
		if m.Maze[sp[0]+v[0]][sp[1]+v[1]] == 0 {
			return false, nil
		} else if i == 3 {
			return true, nil
		}
		if sp[0] >= len(m.Maze) || sp[1] >= len(m.Maze) {
			return false, nil
		}
	}
	return false, errors.New("error: unknown")
}
func roadRand(m *board.Map) ([]int, bool) {
	l := len(m.Maze)
	l2 := len(m.Maze[0])
	res := [][]int{}
	for i := 2; i < l-2; i = i + 2 {
		for j := 2; j < l2-2; j = j + 2 {
			if m.Maze[i][j] == 5 {
				b, _ := noPassage(m, []int{i, j})
				if !b {
					res = append(res, []int{i, j})
				}
			}
		}
	}
	rand.Seed(time.Now().UnixNano())
	if len(res) == 0 {
		return []int{2, 2}, false
	}
	n := rand.Intn(len(res))
	return res[n], true
}

func ShowField(m *board.Map) [][]string {
	res := make([][]string, len(m.Maze))
	for i := 0; i < len(m.Maze); i++ {
		res[i] = make([]string, len(m.Maze[0]))
	}
	w := "■"
	a := "□"
	for i := 0; i < len(m.Maze); i++ {
		for ii := 0; ii < len(m.Maze[0]); ii++ {
			res[i][0], res[i][1] = w, w
			res[i][len(m.Maze[0])-1], res[i][len(m.Maze[0])-2] = w, w
			res[0][ii], res[1][ii] = w, w
			res[len(m.Maze)-1][ii], res[len(m.Maze)-2][ii] = w, w
			if m.Maze[i][ii] == 5 {
				res[i][ii] = a
			} else if m.Maze[i][ii] == 0 {
				res[i][ii] = w
			}
		}
	}
	return res
}
