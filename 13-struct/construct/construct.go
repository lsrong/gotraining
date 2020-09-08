package construct

import (
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Name     string
	Username string
	Password string
}

func NewUser(name, username, password string) *User {
	user := new(User)
	user.Name = name
	user.Username = username
	m := md5.Sum([]byte(password))
	user.Password = hex.EncodeToString(m[:])

	return user
}
