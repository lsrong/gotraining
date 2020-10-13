package _6_interface

import (
	"fmt"
	"testing"
)

func TestCalcSalary(t *testing.T) {
	var employers []Employer
	program01 := NewProgram("test", 10000, 0)
	program02 := NewProgram("test", 8000, 0)
	program03 := NewProgram("test", 20000, 0)
	employers = append(employers, program01)
	employers = append(employers, program02)
	employers = append(employers, program03)

	sale01 := NewSale("test", 3000, 5000)
	sale02 := NewSale("test", 3000, 5000)
	sale03 := NewSale("test", 3000, 5000)
	employers = append(employers, sale01)
	employers = append(employers, sale02)
	employers = append(employers, sale03)
	fmt.Printf("Employee's salary is %d \n", int(CalcSalary(employers)))

	Just(program01)
}
