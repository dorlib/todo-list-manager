package data

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Seeder() error {
	// create users.
	pass, _ := bcrypt.GenerateFromPassword([]byte("12345678"), 14)

	user1, err := CreateUser("Dor", "Admin", string(pass))
	if err != nil {
		return err
	}

	user2, err := CreateUser("Noam", "Admin", string(pass))
	if err != nil {
		return err
	}

	user3, err := CreateUser("Lychee", "User", string(pass))
	if err != nil {
		return err
	}

	users := []User{user1, user2, user3}

	// create groups.
	group, err := CreateGroup("Project Mayhem", "Project Mayhem Is a data science project for learning purposes", users)
	if err != nil {
		return err
	}

	fmt.Printf("group created: %v", group)

	task1, err := CreateTask("make unit tests", "make unit tests for the authorizer", "High", Date{
		DeadlineYear:  "2023",
		DeadlineMonth: "08",
		DeadlineDay:   "23",
	}, user1)
	if err != nil {
		return err
	}

	fmt.Printf("task created: %v", task1)

	task2, err := CreateTask("add dependabot migration", "design and implement dependabot migration", "Medium", Date{
		DeadlineYear:  "2023",
		DeadlineMonth: "09",
		DeadlineDay:   "23",
	}, user2)
	if err != nil {
		return err
	}

	fmt.Printf("task created: %v", task2)

	task3, err := CreateTask("make docker image", "make dockerfile and docker compose", "High", Date{
		DeadlineYear:  "2023",
		DeadlineMonth: "08",
		DeadlineDay:   "20",
	}, user3)
	if err != nil {
		return err
	}

	fmt.Printf("task created: %v", task3)

	task4, err := CreateTask("add tracing", "design and implement tracing including 3rd parties", "High", Date{
		DeadlineYear:  "2023",
		DeadlineMonth: "08",
		DeadlineDay:   "27",
	}, user2)
	if err != nil {
		return err
	}

	fmt.Printf("task created: %v", task4)
	return nil
}
