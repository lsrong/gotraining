package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

/**
15-file
!---listen16
|    |----employee
|    |    |----employee.exe
|    |    |----main.go
|    |----empty_interface
*/
func ListDir(dirPath string, deep int) error {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("Failed to open [%s]:%w", dirPath, err)
	}
	if 1 == deep {
		fmt.Printf("!----%s\n", filepath.Base(dirPath))
	}

	// 分隔符
	pathSeparator := string(os.PathSeparator)
	for _, v := range dir {
		for i := 0; i < deep; i++ {
			fmt.Print("|    ")
		}
		fmt.Printf("|----%s\n", v.Name())

		if v.IsDir() {
			_ = ListDir(dirPath+pathSeparator+v.Name(), deep+1)
			continue
		}
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "tree"

	app.Usage = "List dir file"
	app.Action = func(c *cli.Context) error {
		dirPath := "."
		if c.NArg() > 0 {
			dirPath = c.Args()[0]
		}
		err := ListDir(dirPath, 1)
		if err != nil {
			fmt.Printf("Failed %s", err)
			return err
		}

		return nil
	}

	_ = app.Run(os.Args)

}
