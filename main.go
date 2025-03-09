package main

import (
	"fmt"
	"net/http"
)

// array: they are of fixed size
var shortTask = "Learn Go"
var longTask = "Watch Nana's Golang full course"
var reward = "Reward myself with the power of Go" //alternate way of declaration

var taskItems = []string{shortTask, longTask, reward} //slice: they are of dynamic size

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":3010", nil)
}

func showTasks(w http.ResponseWriter, r *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(w, task)
	}
}

func helloUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello User")
}

func printTasks(taskItems []string) {
	for index, task := range taskItems { //_ Blank identifier
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTask = append(taskItems, newTask)
	return updatedTask
}
