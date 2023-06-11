## TODO List Manager - For Managing And Tracking Your Daily Tasks

Hello everyone 👋 </br>
This is a CLI for task managment using Go, Cobra cli & GORM!

The main purposes of this project are to gain more experience with golang and known & popular frameworks and technologies.
Also practise clean code, SOLID principles of golang and microservices architacture in golang.
In addition, gain experience of taste of what its like working with other team members on the same 
codebase, working with git and with the technologies github has to offer.

The CLI is written with the cobra framework for CLI creation, and GORM as ORM library and database connection, which is MySQL.
In the future, we will continue to extend the app and create more features, and even a UI is on our mind!

<p align="center">
    <img alt="Open Source? Yes!" src="https://badgen.net/badge/Open%20Source%20%3F/Yes%21/blue?icon=github"/>
    <img alt="Issues" src="https://img.shields.io/github/issues-raw/dorlib/todo-list-manager"/>
    <img alt="pull request" src="https://img.shields.io/github/issues-pr-closed/dorlib/todo-list-manager"/>
    <img alt="stars" src="https://img.shields.io/github/stars/dorlib/todo-list-manager?style=social">
    <img alt="updated" src="https://img.shields.io/github/last-commit/dorlib/todo-list-manager">
    <img alt="size" src="https://img.shields.io/github/repo-size/dorlib/todo-list-manager" >
</p>

!Picrute will be here


## Built With

- [Cobra](https://github.com/spf13/cobra) as CLI framework.
- [viper](https://github.com/spf13/viper) complete configuration solution for Go applications.
- [GORM](https://gorm.io/) as ORM library for Golang.
- [go-pretty](https://github.com/jedib0t/go-pretty/tree/main/table) for printing tables.
- [golangci-lint](https://golangci-lint.run/) as golang linter in order to preserve our code's quality.

## What the Application include?

The main entites in the app are : Tasks, Groups & Users.
As you can see in the welcome page of the app, the main "things" you can do are : managing and tracking, so Let's cover the main features!

1. Create Group - you can create a group in order to organize the task's of everyone in one place!
2. Create user - you can register in order to make yourself a user and login in order to access the team's task management system.
3. Manage users - give different premissions to different users.
4. Manage tasks - Create, update and task's content, update progress, and delete tasks (admin only).
5. Track tasks - print in tasks different variations and options to make efficient tracking.
6. Alert system - which will update the user who's the task is belong to by mail when deadline is coming near!
7. Define authorization - you can give admin access for some of the group's members inorder to have extra capabilities!

#### pleae note the README file in the cmd directory of more info about the available commands & flags.

## Microservices Architacture

![services](https://github.com/dorlib/todo-list-manager/assets/90474428/3d466019-81bd-4ed0-a68e-e4036550c744)

## Entity Relationship Diagram

![Entity Relationship Diagram](https://github.com/dorlib/todo-list-manager/assets/90474428/d466c886-2be9-47ee-9c06-2053ba340eaa)

## Running the Application Yourself

Here's what needs to be done in order to get the app running locally on your machine (NOT RELAVENT for now, need to be aligned with the new architacture).
1. Clone the project to your machine with `git clone https://github.com/dorlib/todo-list-manager.git && cd todo-list-manager`.
2. download mysql and create new database (pay attention to give the connection string the right arguments).
3. run `go install`
4. run `todo`
5. help

### Run the docker container:

In order to run the app using docker-compose  (NOT RELAVENT for now, need to be aligned with the new architacture).:
1. Clone the project to your machine with `git clone https://github.com/dorlib/todo-list-manager.git && cd todo-list-manager`.
2. Run `docker-compose buid && docker-compose up`.
3. Enjoy managing your tasks.

### Or follow those steps:

1. Clone the project to your machine with `git clone https://github.com/dorlib/todo-list-manager.git`
2. Install Go on your machine from [here](https://go.dev/doc/install).
3. Install MySQL on your machine from [here](https://www.mysql.com/downloads/).
4. On the root of the project, run the app with `go run main.go`
5. Run `go install`
6. Run in the command line `todo --help` and start manging your tasks!
7. Open [http://localhost:3000](http://localhost:3000) to manage your tasks on the web (IN PROGRESS).

## Seed data to your database

In order to seed pre-made data, for quick testing the app, you can type the command `todo seed`

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. Please make sure to update tests as appropriate.

### _notise the "mission.txt" file in "public" directory to see futre features which planned to be added!_.

### How To Contribute

1. Fork the repository to your own Github account.
2. Clone the project to your machine.
3. Create a branch locally with a succinct but descriptive name.
4. Commit changes to the branch.
5. Following any formatting and testing guidelines specific to this repo.
6. Push changes to your fork.
7. Open a Pull Request in my repository.

## Creator / Maintainer

Dor Liberman ([dorlib](https://github.com/anniedotexe))

If you have any questions or feedback, I would be glad if you will contact me via mail.

<p align="left">
  <a href="dorlibrm@gmail.com"> 
    <img alt="Connect via Email" src="https://img.shields.io/badge/Gmail-c14438?style=flat&logo=Gmail&logoColor=white" />
  </a>
</p>

This project was created for educational purposes, for personal and open-source use.

If you like my content or find my code useful, give it a :star: 

