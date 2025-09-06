package main

import (
	"database/sql"
	"example/internal/models"
	"example/internal/person/repository"
	"example/internal/person/usecase"
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
	personService := usecase.NewPersonService(personRepository)

	router := gin.Default()

	router.POST("/persons/create", func(c *gin.Context) {
		pService := personService
		var p models.Person

		if c.ContentType() != "application/json" {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Content-Type must be application/json"})
			return
		}

		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		pService.AddPerson(p)

		c.JSON(http.StatusOK, gin.H{"status": "person created"})
	})
	router.Run("localhost:8080")
}
