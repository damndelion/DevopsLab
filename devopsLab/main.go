package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=password dbname=app sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	router := gin.New()
	router.POST("/users", createUser)

	router.Run(":8080")
}

func createUser(c *gin.Context) {
	user := User{}
	c.BindJSON(&user)

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&user)

	c.JSON(201, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}
