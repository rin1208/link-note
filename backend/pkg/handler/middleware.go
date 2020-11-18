package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (api API) AuthJWT(c *gin.Context) {
	ua := c.GetHeader("Authorization")

	fmt.Println(ua, c.GetHeader("Uid"))
	c.Next()
}
