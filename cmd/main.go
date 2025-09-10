package main

import (
	"example/internal/person/dtos"
	"example/internal/person/repository"
	"example/internal/person/usecase"
	"example/pkg/db/postgres"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	db := postgres.NewDbConn()
	personRepository := repository.NewPersonRepository(db)
	personService := usecase.NewPersonService(personRepository)
	router := gin.Default()
	router.POST("/person/create", func(c *gin.Context) {
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

		response := pService.AddPerson(&p)
		c.JSON(response.StatusCode, response)
	})
	router.GET("/person/RemovePersonGetById", func(c *gin.Context) {
		pService := personService
		id := c.Query("id")

		val, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Error converting id to integer:", err)
		}
		responseDto := pService.RemovePersonGetById(val)
		c.JSON(responseDto.StatusCode, responseDto)

	})

	router.Run("localhost:8080")
}
