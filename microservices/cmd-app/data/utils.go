package data

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/jedib0t/go-pretty/v6/table"
	"log"
	"os"
	"strconv"
	"time"
)

func printTask(task Task) {
	year, err := strconv.Atoi(task.DeadlineDate.DeadlineYear)
	if err != nil {
		fmt.Println(err)
	}

	month, err := strconv.Atoi(task.DeadlineDate.DeadlineMonth)
	if err != nil {
		fmt.Println(err)
	}

	day, err := strconv.Atoi(task.DeadlineDate.DeadlineDay)
	if err != nil {
		fmt.Println(err)
	}

	endingTime := task.CreatedAt.AddDate(year, month, day)
	timeLeft := endingTime.Sub(task.CreatedAt)

	fmt.Printf("# %d: %v %v %v %v %v %v", task.ID, task.Title, task.Description, task.Priority, task.CreatedAt, task.Deadline, timeLeft)
}

func printTasks(tasks []taskSummery) {
	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "User", "Title", "Description", "Priority", "Created At", "Deadline", "Status", "Days Left"})

	for i := 0; i < len(tasks); i++ {
		diff, err := getDiffOfDates(tasks[i].Deadline)
		if err != nil {
			log.Println(err)
		}

		tasks[i].TimeLeft = diff
		t.AppendRows([]table.Row{structs.Values(tasks[i])})
	}

	t.AppendFooter(table.Row{"Total", len(tasks)})
	t.Render()
}

func TaskExistByID(taskID uint) bool {
	var task Task

	DB.First(&task, taskID)

	return task.ID != 0
}

func TaskExistByName(taskName string) bool {
	var task Task

	DB.Where("name = ?", taskName).First(&task)

	return task.ID != 0
}

func CheckLegalPriority(priority string) bool {
	if priority == Critical || priority == VeryHigh || priority == High || priority == Medium || priority == Low {
		return true
	}

	return false
}

func getDiffOfDates(deadline string) (int, error) {
	layout := "02/01/2006" // dd/mm/yyyy format

	date1, err := time.Parse(layout, deadline)
	if err != nil {
		err = fmt.Errorf("error while parsing date: %v", deadline)

		return 0, err
	}

	date2 := time.Now()

	diff := int(date1.Sub(date2).Hours() / 24)

	return diff, nil
}
