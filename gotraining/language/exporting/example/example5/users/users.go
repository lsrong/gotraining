package users

type user struct {
	Name string
	ID   int
}

type Manager struct {
	Title string

	user
}
