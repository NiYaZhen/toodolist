package main

import (
	db "todolist/todolist/app/infra/persistence/mongo"
	t "todolist/todolist/cmd"
)

func init() {
	db.LoadTheEnv()
	db.CreateDBInstance()

}

func main() {
	t.RunToDoListServer()

}
