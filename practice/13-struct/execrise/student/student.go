package student

type Student struct {
	Name  string
	Sex   string
	Scope int
	Grade string
}

func NewStudent(name, sex string, scope int, grade string) *Student {
	return &Student{
		Name:  name,
		Sex:   sex,
		Scope: scope,
		Grade: grade,
	}
}
