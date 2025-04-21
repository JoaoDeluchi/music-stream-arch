package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	// TODO: implement each layer and after that implement the build functions 
	db := buildDatabase()
	repo := buildRepositories(db)
	services := buildServices(repo)
	controllers := buildApp(services)

	_, err := RunApp()
}

func buildDatabase() db {
	// TODO: implement Enviroment configurations for database and provide to datasource
	usersDB := datasource.NewDatabase("Users")
}