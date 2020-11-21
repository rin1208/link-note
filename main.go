package main

import (
	"link-note/backend/pkg"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

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
			"cache-control",
			"user_id",
		},
		AllowOrigins: []string{
			os.Getenv("FRONT_URL"),
		},
		MaxAge: 24 * time.Hour,
	}))

	pkg.Serve(r, ":"+os.Getenv("PORT"))
}
