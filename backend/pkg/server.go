package pkg

import (
	"fmt"
	"linknote/backend/pkg/handler"
	"os"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine, port string) {

	CreateFireStoreJson()
	api := handler.Init_API()

	r.POST("/post", api.Post)
	r.GET("/getcontent", api.GetContent)
	r.Run(port)
}

func CreateFireStoreJson() {
	fp, err := os.Create("./firestore.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	file := fmt.Sprintf(` {
	"type": "%s",
	"project_id": "%s",
	"private_key_id": "%s",
	"private_key": "%s",
	"client_email": "%s",
	"client_id": "%s",
	"auth_uri": "%s",
	"token_uri": "%s",
	"auth_provider_x509_cert_url": "%s",
	"client_x509_cert_url": "%s"
}`,
		os.Getenv("FS_TYPE"),
		os.Getenv("FS_PROJECT_ID"),
		os.Getenv("FS_PRIVATE_KEY_ID"),
		os.Getenv("FS_PRIVATE_KEY"),
		os.Getenv("FS_CLIENT_EMAIL"),
		os.Getenv("FS_CLIENT_ID"),
		os.Getenv("FS_AUTH_URI"),
		os.Getenv("FS_TOKEN_URI"),
		os.Getenv("FS_AUTH_PROVIDER_X509_CERT_URL"),
		os.Getenv("FS_AUTH_PROVIDER_X509_CERT_URL"))

	_, err = fp.Write(([]byte)(file))
	if err != nil {
		fmt.Println(err)
	}
}
