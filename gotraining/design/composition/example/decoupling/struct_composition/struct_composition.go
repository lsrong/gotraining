package main

// Sample program demonstrating decoupling with interfaces.
// 演示利用接口解耦的示例程序。
import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Data 数据实体
type Data struct {
	Line string
}

// Xenia 拉取数据操作体
type Xenia struct {
	Host    string
	Timeout time.Duration
}

func (*Xenia) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF
	case 5:
		return errors.New("error reading data from Xenia")
	default:
		d.Line = "data"
		fmt.Println("In: ", d.Line)
		return nil
	}
}

// Pillar 保存数据操作体。
type Pillar struct {
	Host    string
	Timeout time.Duration
}

func (*Pillar) Store(d *Data) error {
	fmt.Println("Out: ", d.Line)
	return nil
}

// System 组成系统结构体。
type System struct {
	Xenia
	Pillar
}

func pull(x *Xenia, data []Data) (int, error) {
	for i := range data {
		if err := x.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

func store(p *Pillar, data []Data) (int, error) {
	for i := range data {
		if err := p.Store(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

// Copy 拷贝操作，先拉取数据，然后保存数据
func Copy(s *System, batch int) error {
	data := make([]Data, batch)
	for {
		i, err := pull(&s.Xenia, data)
		if i > 0 {
			if _, err := store(&s.Pillar, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

func main() {
	// System 组合Xenia，Pillar分别实现行为接口Puller, Storer.
	sys := System{
		Xenia:  Xenia{Host: "localhost:3000"},
		Pillar: Pillar{Host: "localhost:4000"},
	}
	batch := 3
	if err := Copy(&sys, batch); err != io.EOF {
		fmt.Println(err)
	}
}
