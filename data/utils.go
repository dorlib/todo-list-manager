package data

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"strconv"
)

func printTask(task Task) {
	year, err := strconv.Atoi(task.Deadline.Year)
	if err != nil {
		fmt.Println(err)
	}

	month, err := strconv.Atoi(task.Deadline.Month)
	if err != nil {
		fmt.Println(err)
	}

	day, err := strconv.Atoi(task.Deadline.Day)
	if err != nil {
		fmt.Println(err)
	}

	endingTime := task.CreatedAt.AddDate(year, month, day)
	timeLeft := endingTime.Sub(task.CreatedAt)

	fmt.Printf("# %d: %v %v %v %v %v %v", task.ID, task.Title, task.Description, task.Priority, task.CreatedAt, task.Deadline, timeLeft)
}

func printAllTasks(tasks []Task) {
	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Title", "Description", "Priority", "Created At", "Deadline", "Time Left"})

	for i := 0; i < len(tasks); i++ {
		var row = table.Row{tasks[i]}

		t.AppendRows([]table.Row{row})
	}

	t.AppendFooter(table.Row{"Total", 10000})
	t.Render()
}

func TaskExistByID(taskID uint) bool {
	var task Task

	tx := DB.First(&task, taskID)

	return tx != nil
}

func TaskExistByName(taskName string) bool {
	var task Task

	tx := DB.Where("name = ?", taskName).First(&task)

	return tx != nil
}

func CheckLegalPriority(priority string) bool {
	if priority == Critical || priority == VeryHigh || priority == High || priority == Medium || priority == Low {
		return true
	}

	return false
}
