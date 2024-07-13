package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	var showHelpMenuFlag = flag.Bool("h", false, "Show menu with all commands and usage.")
	var addNewTaskFlag = flag.Bool("a", false, "Add a new task.")
	var updateTaskFlag = flag.Bool("u", false, "Update a task.")
	var deleteTaskFlag = flag.Bool("d", false, "Delete a task.")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	if *showHelpMenuFlag {
		fmt.Printf("\t-a\tAdd a new task\n\r\t-u\tUpdate a task(Completed or not).\n\r\t-e\tEdit a task\n\r\t-d\tDelete a task\n\r")
		return
	}

	if *addNewTaskFlag {
		fmt.Printf("Add new task: ")
		task, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading new task.")
			return
		}
		fmt.Printf("New task %s\n", task)
		return

	}

	if *updateTaskFlag {
		fmt.Println("Enter task ID.")
		taskId,err:=reader.ReadString('\n')
		if err!= nil{
			fmt.Println("Error reading task ID.")
			return
		}

		fmt.Println("Marked task as completed/not completed.")
		return
	}

	if *deleteTaskFlag {
		fmt.Println("Enter task ID.")
		taskId, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading task ID")
			return
		}
		fmt.Println("Deleted task ", taskId)
		return
	}

	fmt.Println("All tasks")

}
