// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package toy contains support for managing toy inventory.
package toy

// Toy type
type Toy struct {
	Name   string
	Weight int

	onHand int
	sold   int
}

// New 初始化函数
func New(name string, weight int) *Toy {
	return &Toy{
		Name:   name,
		Weight: weight,
	}
}

func (t *Toy) OnHand() int {
	return t.onHand
}

func (t *Toy) UpdateOnHand(num int) int {
	t.onHand += num

	return t.onHand
}

func (t *Toy) Sold() int {
	return t.sold
}

func (t *Toy) UpdateSold(num int) int {
	t.sold += num

	return t.sold
}
