/*
Copyright © 2023 todo-list <dorlib318@gmail.com>
*/

package main

import (
	"todo/cmd"
	"todo/data"
)

func main() {
	data.OpenDataBase()
	cmd.Execute()
}
