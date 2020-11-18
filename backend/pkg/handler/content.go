package handler

import (
	"fmt"
	"linknote/backend/pkg/infra"
	"linknote/backend/pkg/model"
	"linknote/backend/pkg/usecase"

	"github.com/gin-gonic/gin"
)

type API struct {
	FireBaseClient *infra.FireBase
}

func Init_API() *API {
	return &API{
		FireBaseClient: infra.Init_firebase(),
	}
}

func (api API) Post(c *gin.Context) {
	var data model.Content
	c.BindJSON(&data)
	data.Content_id = usecase.Uuid4()
	data.Date = usecase.GetDeteInTokyo()
	api.FireBaseClient.InsertData(data)

	c.JSON(200, data)
}
func (api API) GetContent(c *gin.Context) {

	dataa := api.FireBaseClient.GetData(c.GetHeader("Uid"))
	fmt.Println(dataa)

	c.JSON(200, dataa)
}
