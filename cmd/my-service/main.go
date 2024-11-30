package main

import (
	"dbproject/internal/db"
	"fmt"
)

func main() {
	dataSourceName := "KP="
	db.InitDB(dataSourceName)

	fmt.Println("yeAH!")
}
