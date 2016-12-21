package main

import (
	"fmt"
	"os"

	"github.com/ycheng/rpdb/dirs"
)

func useage(arg0 string) {
	fmt.Println("Useage: " + arg0 + " list\n")
}

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// arg := os.Args[3]

	// dbPath := os.Getenv("RPDB_HOME")
	// fmt.Println("db_path: " + dbPath)
	fmt.Println(dirs.GlobalDataDir)
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	// fmt.Println(arg)
}
