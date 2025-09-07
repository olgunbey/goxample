package main

import (
	"example/internal/person/dtos"
	"example/internal/person/repository"
	"example/internal/person/usecase"
	"example/pkg/db/postgres"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	dizi := make(map[string]string)
	dizi["test"] = "data"

	for key, value := range dizi {
		fmt.Println(key, value)
	}

	db := postgres.NewDbConn()
	personRepository := repository.NewPersonRepository(db)
	personService := usecase.NewPersonService(personRepository)
	router := gin.Default()
	router.POST("/persons/create", func(c *gin.Context) {
		pService := personService
		var p dtos.AddPersonRequestDto

		if c.ContentType() != "application/json" {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Content-Type must be application/json"})
			return
		}
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		pService.AddPerson(&p)

		c.JSON(http.StatusOK, gin.H{"status": "person created"})
	})
	router.Run("localhost:8080")
}
