package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type task struct {
	id         int
	title      string
	completed  int
	created_at string
	updated_at string
}

func handleFlag(db *sql.DB, showHelpMenuFlag, addNewTaskFlag, updateTaskFlag, deleteTaskFlag bool, allTasks []task) {
	reader := bufio.NewReader(os.Stdin)

	if showHelpMenuFlag {
		fmt.Printf("\t-a\tAdd a new task\n\r\t-u\tUpdate a task(Completed or not).\n\r\t-e\tEdit a task\n\r\t-d\tDelete a task\n\r")
		return
	} else if addNewTaskFlag {
		addNewTask(reader, db)
	} else if updateTaskFlag {
		updateTask(reader, db, allTasks)
	} else if deleteTaskFlag {
		deleteTask(reader, db, allTasks)
	}
	printAllTasks(getAllTasks(db))
}

func addNewTask(reader *bufio.Reader, db *sql.DB) {
	fmt.Printf("Enter new task: ")
	task, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading new task.")
		return
	}
	task = strings.TrimSuffix((strings.TrimSpace(task)), "\n")

	stmt, err := db.Prepare("INSERT INTO tasks(title,completed) VALUES(?,0)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	stmt.Exec(task)
}

func updateTask(reader *bufio.Reader, db *sql.DB, allTasks []task) {
	fmt.Printf("Enter task ID: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	taskId, err := strconv.Atoi(strings.TrimSuffix((strings.TrimSpace(input)), "\n"))
	if err != nil {
		log.Fatal("Enter valid task ID")
	}

	stmt, err := db.Prepare("UPDATE tasks SET completed = CASE WHEN completed = 0 THEN 1 ELSE 0 END WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	fmt.Println("taskid", taskId)
	stmt.Exec(allTasks[taskId].id)
}

func deleteTask(reader *bufio.Reader, db *sql.DB, allTasks []task) {
	fmt.Printf("Enter task ID: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	taskId, err := strconv.Atoi(strings.TrimSuffix((strings.TrimSpace(input)), "\n"))
	if err != nil {
		log.Fatal("Enter valid task ID")
	}

	stmt, err := db.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(allTasks[taskId].id)
}

func getAllTasks(db *sql.DB) []task {
	var allTasks []task
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var t task
		rows.Scan(&t.id, &t.title, &t.completed, &t.created_at, &t.updated_at)

		allTasks = append(allTasks, t)
	}

	return allTasks

}

func printAllTasks(allTasks []task) {
	if len(allTasks) == 0 {
		fmt.Println("No tasks.")
	} else {
		for i, t := range allTasks {
			fmt.Printf("%d.\t%s |\t", i, t.title)
			if t.completed == 1 {
				fmt.Printf("%s\n", "Completed")
			} else {
				fmt.Printf("%s\n", "Not Completed")
			}
		}
		fmt.Printf("\nTotal tasks: %d\n", len(allTasks))
	}
}

func main() {

	var showHelpMenuFlag = flag.Bool("h", false, "Show menu with all commands and usage.")
	var addNewTaskFlag = flag.Bool("a", false, "Add a new task.")
	var updateTaskFlag = flag.Bool("u", false, "Update a task.")
	var deleteTaskFlag = flag.Bool("d", false, "Delete a task.")
	flag.Parse()

	db, err := sql.Open("sqlite3", os.Getenv("HOME")+"/task.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS tasks('id' INTEGER PRIMARY KEY AUTOINCREMENT, 'title' TEXT, 'completed' INTEGER CHECK(completed IN (0,1)), 'created_at' TEXT, 'updated_at' TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	allTasks := getAllTasks(db)

	handleFlag(db, *showHelpMenuFlag, *addNewTaskFlag, *updateTaskFlag, *deleteTaskFlag, allTasks)
}
