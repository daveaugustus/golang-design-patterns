package routes

import (
	"fmt"
	"net/http"

	"github.com/davetweetlive/offersapp/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

// UsersLogin vw
func UsersLogin(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	err := user.IsAuthenticated(&conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := user.GetAuthToken()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": "There was an error authenticating.",
	})
}

// UserRegister asd
func UserRegister(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)
	if err := user.Register(&conn); err != nil {
		fmt.Println("Error in user.Register()")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := user.GetAuthToken()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	}
	// c.JSON(http.StatusCreated, gin.H{"user_id": user.ID})
	// c.JSON(http.StatusCreated, gin.H{"error": "There was an error authenticating."})
}
