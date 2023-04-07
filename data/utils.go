package data

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func printTask(task Task) {
	timeLeft := task.Deadline.Sub(task.CreatedAt)
	fmt.Printf("# %d: %v %v %v %v %v %v", task.ID, task.Title, task.Description, task.Priority, task.CreatedAt, task.Deadline, timeLeft)
}

func printAllTasks(tasks []Task) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Title", "Description", "Priority", "Created At", "Deadline", "Time Left"})
	t.AppendRows([]table.Row{})
	t.AppendFooter(table.Row{"Total", 10000})
	t.Render()
}
