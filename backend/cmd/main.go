package main

import (
	"link-note/backend/pkg"
	"time"

	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			fmt.Println(err)
		}
		port = os.Getenv("LOCAL_PORT")
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Content-Type",
			"Content-Length",
			"Authorization",
			"Uid",
		},
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://link-no-te.web.app",
			"https://link-no-te.firebaseapp.com",
		},
		MaxAge: 24 * time.Hour,
	}))

	pkg.Serve(r, ":"+port)

}
