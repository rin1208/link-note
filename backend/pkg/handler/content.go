package handler

import (
	"link-note/backend/pkg/infra"
	"link-note/backend/pkg/model"
	"link-note/backend/pkg/usecase"
	"regexp"

	"github.com/gin-gonic/gin"
)

type API struct {
	FireBaseClient usecase.UserInteractor
}

func Init_API(firebase infra.FireBaseHandler) *API {
	return &API{
		FireBaseClient: usecase.UserInteractor{
			UserRepository: &infra.FireBase{
				FireBaseHandler: firebase,
			},
		},
	}

}

func (api API) Post(c *gin.Context) {
	var data model.Content
	c.BindJSON(&data)
	match, _ := regexp.MatchString("http", data.Url)
	if !match {
		c.JSON(400, "Forbidden")
	} else {
		data.Content_id = usecase.Uuid4()
		data.Date = usecase.GetDeteInTokyo()
		data.Uid = c.GetHeader("Uid")
		api.FireBaseClient.InsertData(data)
		c.JSON(200, "ok")
	}

}
func (api API) GetContent(c *gin.Context) {

	data := api.FireBaseClient.GetData(c.GetHeader("Uid"))

	c.JSON(200, data)
}

func (api API) DeleteContent(c *gin.Context) {

	var data model.Content
	c.BindJSON(&data)

	err := api.FireBaseClient.DeleteData(c.GetHeader("Uid"), data.Content_id)

	c.JSON(200, err)
}
