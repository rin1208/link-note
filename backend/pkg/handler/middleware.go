package handler

import (
	"github.com/gin-gonic/gin"
)

func (api API) AuthJWT(c *gin.Context) {
	jwt := c.GetHeader("Authorization")

	err := api.FireBaseClient.AuthJWT(jwt)
	if err != nil {
		c.JSON(404, "Forbidden")
		c.Abort()
	}
	c.Next()
}
