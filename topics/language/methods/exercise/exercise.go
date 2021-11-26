package main

import "fmt"

type player struct {
	name string
	bats int
	hits int
	avg  float64
}

func (p *player) setAverage() {
	if p.bats == 0 {
		p.avg = 0.0
		return
	}

	p.avg = float64(p.hits) / float64(p.bats)
}

func main() {
	players := []player{
		{name: "Li", bats: 20, hits: 14},
		{name: "Ye", bats: 12, hits: 6},
		{name: "Ming", bats: 14, hits: 7},
	}

	// 不推荐以下的range值遍历, 不会作用到原切片元素里面 !!!!!, 副本遍历
	for _, p := range players {
		p.setAverage()
		fmt.Printf("Name: %s, Average: [.%.f] \n", p.name, p.avg*1000)
	}

	fmt.Printf("%v\n", players)

	fmt.Println("#############################")

	// 指针遍历方式,直接用切片的原值调用,可作用到原值, 指针遍历
	for i := range players {
		players[i].setAverage()
		fmt.Printf("Name: %s, Average: [.%.f] \n", players[i].name, players[i].avg*1000)
	}

	fmt.Printf("%v\n", players)

}
