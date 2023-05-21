package data

func Seeder() {
	CreateTask("make unit tests", "make unit tests for the authorizer", "High", Date{
		DeadlineYear:  "2023",
		DeadlineMonth: "08",
		DeadlineDay:   "23",
	}, User{})

	CreateTask("add dependabot migration", "design and implement dependabot migration", "Medium", Date{
		DeadlineYear:  "2023",
		DeadlineMonth: "09",
		DeadlineDay:   "23",
	}, User{})

	CreateTask("make docker image", "make dockerfile and docker compose", "High", Date{
		DeadlineYear:  "2023",
		DeadlineMonth: "08",
		DeadlineDay:   "20",
	}, User{})

	CreateTask("add tracing", "design and implement tracing including 3rd parties", "High", Date{
		DeadlineYear:  "2023",
		DeadlineMonth: "08",
		DeadlineDay:   "27",
	}, User{})
}
