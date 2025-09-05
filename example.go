package main

import (
	"database/sql"
	"example/domain"
	"example/repository"
	"example/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "username123"
	password = "password123"
	dbname   = "testdb"
)

func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	personRepository := repository.NewPersonRepository(db)
	personService := service.NewPersonService(personRepository)

	router := gin.Default()

	router.POST("/persons/create", func(c *gin.Context) {
		var p domain.Person

		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		personService.AddPerson(p)

		c.JSON(http.StatusOK, gin.H{"status": "person created"})

	})
	router.Run("localhost:8080")
}
