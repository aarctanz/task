package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type task struct {
	id         int
	title      string
	completed  int
	created_at string
	updated_at string
}

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
		fmt.Printf("Enter task ID: ")
		taskId, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading task ID.")
			return
		}

		fmt.Println("Marked task as completed/not completed.", taskId)
		return
	}

	if *deleteTaskFlag {
		fmt.Printf("Enter task ID: ")
		taskId, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading task ID")
			return
		}
		fmt.Println("Deleted task ", taskId)
		return
	}

	db, err := sql.Open("sqlite3", "./task.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS tasks('id' INTEGER, 'title' TEXT, 'completed' BOOL, 'created_at' DATETIME, 'updated_at' DATETME, PRIMARY KEY('id' AUTOINCREMENT))")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM tasks;")
	if err != nil {
		log.Fatal(err)
	}
	var t task
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&t.id, &t.title, &t.completed, &t.created_at, &t.updated_at)

		fmt.Println(t.id, " | ", t.title, " | ", t.completed)
	}

}
