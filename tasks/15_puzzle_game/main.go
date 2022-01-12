package main

import (
	"fmt"
	"strings"
)

type board [16]cell
type cell uint8
type move uint8

func (b *board) String() string {
	var s strings.Builder
	for i, c := range b {
		if c == 0 {
			s.WriteString("  #")
		} else {
			_, _ = fmt.Fprintf(&s, "%3d", c)
		}
		if i%4 == 3 {
			s.WriteString("\n")
		}
	}
	return s.String()
}

var solvedBoard = board{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0}

const (
	up move = iota
	down
	left
	right
)

const (
	easy   = 10
	normal = 50
	hard   = 100
)

// puzzle 处理逻辑, 移动相当于交换空格和执行数字的下标,交换的下标等同于移动的位置, 向上 -4, 向下 +4, 向左 -1, 向右 +1(处理边界)
// 结束时候和标准数组匹配是否成功.
type puzzle struct {
	board board
	empty int
	moves int
	quit  bool
	level int
}

func newPuzzle(level uint) *puzzle {
	p := puzzle{
		board: solvedBoard,
		empty: 15,
	}

	// 初始化
	p.shuffle(level)
	p.moves -= int(level)

	return &p
}

func (p *puzzle) shuffle(moves uint) {

}

func (p *puzzle) move(m move) bool {
	emptyIndex := p.empty
	newIndex, ok := p.isValidMove(m)
	if ok {
		p.board[emptyIndex], p.board[newIndex] = p.board[newIndex], p.board[emptyIndex]
		p.empty = newIndex
		p.moves++
	}

	return ok
}

// isValidMove 计算下一次移动位置,处理位置边界.
func (p *puzzle) isValidMove(m move) (int, bool) {
	switch m {
	case up:
		// 向上 -4
		return p.empty - 4, p.empty/4 > 0
	case down:
		// 向下 + 4
		return p.empty + 4, p.empty/4 < 3
	case left:
		// 向左 -1, 左边边界为 0 , 4, 8, 16
		return p.empty - 1, p.empty%4 > 0
	case right:
		// 向右 +1, 右边边界为 3, 7, 11, 15
		return p.empty + 1, p.empty%4 < 3
	default:
		panic("invalid move!")
	}
}

func (p *puzzle) play() {

}

func main() {
	fmt.Print(solvedBoard.String(), easy, normal, hard)
}
