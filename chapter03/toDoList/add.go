package main

import (
	"fmt"
	"strconv"
)

var todos = make([]map[string]string, 0)

const (
	id     = "id"
	name   = "name"
	sTime  = "start_time"
	eTime  = "end_time"
	status = "status"
	user   = "user"
)

func genId() int {
	var rt int
	for _, todo := range todos {
		todoId, _ := strconv.Atoi(todo["id"])
		if rt < todoId {
			rt = todoId
		}
	}
	return rt + 1
}

func newTask() map[string]string {
	task := make(map[string]string)
	task[id] = strconv.Itoa(genId())
	task[name] = ""
	task[sTime] = ""
	task[eTime] = ""
	task[status] = ""
	task[user] = ""
	return task
}

func main() {

	var text string

	task := newTask()

	fmt.Println("Please Input: ")
	fmt.Print("Mission Name: ")
	fmt.Scan(&text)
	task[name] = text
	fmt.Print("Start Time: ")
	fmt.Scan(&text)
	task[sTime] = text
	fmt.Print("On duty: ")
	fmt.Scan(&text)
	task[user] = text

	todos = append(todos, task)
	fmt.Println(todos)
}
