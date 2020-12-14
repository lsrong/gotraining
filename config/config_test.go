package config

import (
	"testing"
)

type Config struct {
	Redis Redis `ini:"redis"`
	Mysql Mysql `ini:"mysql"`
}

/**
host=localhost
port=6379
password=test
*/
type Redis struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
}

/**
host=127.0.0.1
port=3306
username=root
password=root
database=golang
charset=utf-8
*/
type Mysql struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database string `ini:"database"`
	Charset  string `ini:"charset"`
}

func TestUnMarshalFile(t *testing.T) {
	path := "./app.ini"
	var config Config
	_ = UnMarshalFile(path, &config)
	t.Logf("UnMarshalFile success, config:%#v", config)
}
