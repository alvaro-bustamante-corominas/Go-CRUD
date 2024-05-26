package main

import (
	"database/sql"
	"fmt"
	"go-crud-tareas/handlers"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	server := gin.Default()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(dsn)

	database, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Error conecting to the database: %s", err)
	}
	defer database.Close()

	if err := database.Ping(); err != nil {
		log.Fatalf("It was not possible to connect to the database: %s", err)
	}

	//Endpoints

	server.POST("/task", handlers.CreateTaskHandler(database))

	server.DELETE("/task/:id", handlers.DeleteTaskHandler(database))

	server.GET("/task", handlers.GetTaskHandler(database))

	server.PUT("/task/:id", handlers.UpdateTaskHandler(database))

	server.PATCH("/task/:id", handlers.UpdateStatusTaskHandler(database))

	server.Run(":8080")

}
