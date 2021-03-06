package user

import (
	"com.example.go-kafka-restful/db"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

func Router(r *gin.RouterGroup) {
	_ = db.Connection.AutoMigrate(&User{})
	r.POST("/", addUser())
	r.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		value, _ := strconv.Atoi(id)
		c.JSON(200, getUserFromId(value))
	})
}

func addUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		var userDTO UserDTO
		_ = json.Unmarshal(body, &userDTO)
		c.JSON(200, gin.H{
			"id": insertUser(userDTO),
		})
	}
}
