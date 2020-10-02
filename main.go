package main

import (
	"context"
	"fmt"

	"github.com/davetweetlive/offersapp/routes"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func main() {

	conn, err := connectDB()
	if err != nil {
		return
	}

	router := gin.Default()

	router.Use(dbMiddleware(*conn))

	usersGroup := router.Group("users")
	{
		usersGroup.POST("register", routes.UserRegister)
	}

	router.Run(":8000")

}

func connectDB() (c *pgx.Conn, err error) {
	conn, err := pgx.Connect(context.Background(), `postgres://buxolvcz:3Tw8-FCDL4qHPO0Vicshej9fFGcH1v6c@lallah.db.elephantsql.com:5432/buxolvcz`)
	if err != nil {
		fmt.Println("Error connecting to DB")
		fmt.Println(err.Error())
	}
	if err := conn.Ping(context.Background()); err != nil {
		fmt.Println("Internal server error!")
	}
	return conn, err
}

func dbMiddleware(conn pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}
