package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type board [16]cell
type cell uint8
type move uint8

// String 打印4*4宫格
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

// newPuzzle 创建一个新迷宫游戏对象
func newPuzzle(level uint) *puzzle {
	p := puzzle{
		board: solvedBoard,
		empty: 15,
	}

	// 初始化: 10-easy, 50-normal,100-hard
	p.shuffle(level)
	p.moves -= int(level)

	return &p
}

// shuffle 打乱顺序,初始化游戏难度.
func (p *puzzle) shuffle(moves uint) {
	for i := 0; i < int(moves) || p.board == solvedBoard; {
		m := move(rand.Intn(4))
		if p.move(m) {
			i++
		}
	}
}

// move 移动一次步伐
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

// playOneMove 移动一次操作
func (p *puzzle) playOneMove() {
	for {
		fmt.Printf("Please Move #%d (U-up, D-down, L-left, R-right, Q-quit): ", p.moves+1)
		// 读取标准输入的一个字符
		var s string
		if n, err := fmt.Scanln(&s); err != nil || n != 1 {
			continue
		}

		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		var m move
		switch s[0] {
		case 'U', 'u':
			m = up
		case 'd', 'D':
			m = down
		case 'l', 'L':
			m = left
		case 'r', 'R':
			m = right
		default:
			fmt.Println(`
Please enter "U", "D", "L", or "R" to move the # cell
up, down, left, or right. You can also enter "Q" to quit.
Upper or lowercase is accepted and only the first non-blank
character is important (i.e. you may enter "up" if you like).`)
			continue

		}
		if !p.move(m) {
			fmt.Println("That is not a valid move at the moment.")
		}
		return
	}
}

// play 开始游戏
func (p *puzzle) play() {
	fmt.Println("Starting puzzle...")
	// 执行一次移动
	for p.board != solvedBoard && !p.quit {
		fmt.Printf("\n%+v\n", &p.board)
		p.playOneMove()
	}

	// 结束判断是否成功解决15-puzzle
	if p.board == solvedBoard {
		fmt.Printf("You solved the puzzle in %d moves.\n", p.moves)
	}
}

func main() {
	p := newPuzzle(normal)
	p.play()
}
