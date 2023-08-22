/*
Copyright © 2023 todo-list <dorlib318@gmail.com>
*/

package main

import (
	cmd "todo/client"
	data "todo/data"
)

func main() {
	data.OpenDataBase()
	cmd.Execute()
}
