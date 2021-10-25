package main

type Student struct {
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Scope int    `json:"scope"`
	Grade string `json:"grade"`
}

func NewStudent(name string, sex string, scope int, grade string) *Student {
	return &Student{
		Name:  name,
		Sex:   sex,
		Scope: scope,
		Grade: grade,
	}
}
