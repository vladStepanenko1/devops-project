package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vladStepanenko1/project/repository"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error with loading environment variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	})
	if err != nil {
		log.Fatalf("Failed to initialize database: %s", err.Error())
	}
	db.Ping()

	router := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/", func(c *gin.Context) {
				var tempUser repository.User
				var usersList []*repository.User
				selectUsersQuery := "SELECT name, email, phone_number as phoneNumber from users"
				rows, err := db.Queryx(selectUsersQuery)

				for rows.Next() {
					err = rows.StructScan(&tempUser)
					usersList = append(usersList, &tempUser)
				}
				if err != nil {
					log.Fatalf("Error with selecting all users: %s", err.Error())
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
					})
					return
				}

				c.JSON(http.StatusOK, &usersList)
			})
			users.POST("/", func(c *gin.Context) {
				var input repository.User
				if err := c.BindJSON(&input); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
					return
				}

				var id int
				createUserQuery := "INSERT INTO users (name, email, phone_number) values ($1, $2, $3) returning id"
				row := db.QueryRow(createUserQuery, input.Name, input.Email, input.PhoneNumber)
				if err = row.Scan(&id); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
					})
					return
				}

				c.JSON(http.StatusCreated, gin.H{
					"id": id,
				})
			})
		}
	}
	router.Run(":8000")
}
