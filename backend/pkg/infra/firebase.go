package infra

import (
	"context"
	"link-note/backend/pkg/model"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/iterator"
)

type FirestoreAuth struct {
	Type                        string `json:"type"`
	Project_id                  string `json:"project_id"`
	Private_key_id              string `json:"private_key_id"`
	Private_key                 string `json:"private_key"`
	Client_email                string `json:"client_email"`
	Client_id                   string `json:"client_id"`
	Auth_uri                    string `json:"auth_uri"`
	Token_uri                   string `json:"token_uri"`
	Auth_provider_x509_cert_url string `json:"auth_provider_x509_cert_url"`
	Client_x509_cert_url        string `json:"client_x509_cert_url"`
}

type FireBaseClient struct {
	FireBase      *firebase.App
	FireStore     *firestore.Client
	Ctx           context.Context
	CollectionRef *firestore.CollectionRef
	DocumentRef   *firestore.DocumentRef
	Auth          *auth.Client
}

type FireBaseHandler interface {
	Collection(path string) *FireBaseClient
	Set(ctx context.Context, data interface{}) error
	Doc(id string) *FireBaseClient
	Documents(ctx context.Context) *firestore.DocumentIterator
	Delete(ctx context.Context) error
}

type FireBase struct {
	FireBaseHandler
}

func Init_firebase() FireBaseHandler {

	ctx := context.Background()
	sa := option.WithCredentialsFile("./firestore.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil
	}
	client, err := app.Firestore(ctx)

	if err != nil {
		return nil
	}
	auth, err := app.Auth(ctx)
	if err != nil {
		return nil

	}

	return &FireBaseClient{
		FireBase:  app,
		FireStore: client,
		Ctx:       ctx,
		Auth:      auth,
	}
}

func (fb *FireBase) InsertData(data model.Content) {

	updateError := fb.Collection(data.Uid).Doc(data.Content_id).Set(context.Background(), map[string]interface{}{
		"content_id": data.Content_id,
		"comment":    data.Comment,
		"url":        data.Url,
		"date":       data.Date,
	})
	if updateError != nil {
		log.Printf("An error has occurred: %s", updateError)
	}
}

func (fb *FireBase) DeleteData(uid, id string) error {

	err := fb.Collection(uid).Doc(id).Delete(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (fb *FireBase) GetData(uid string) []model.Content {

	var res_data []model.Content
	iter := fb.Collection(uid).Documents(context.Background())

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		data := doc.Data()
		var content model.Content
		content.Comment = data["comment"].(string)
		content.Url = data["url"].(string)
		content.Content_id = data["content_id"].(string)
		content.Date = int(data["date"].(int64))

		res_data = append(res_data, content)
	}

	return res_data

}

func (fb *FireBaseClient) AuthJWT(jwt string) error {

	auth, err := fb.FireBase.Auth(fb.Ctx)
	if err != nil {

		return err

	}
	idToken := strings.Replace(jwt, "Bearer ", "", 1)

	_, err = auth.VerifyIDToken(fb.Ctx, idToken)
	if err != nil {

		return err
	}
	return nil
}

func (fb *FireBaseClient) Collection(path string) *FireBaseClient {
	fb.CollectionRef = fb.FireStore.Collection(path)
	return fb
}

func (fb *FireBaseClient) Set(ctx context.Context, data interface{}) error {
	_, err := fb.DocumentRef.Set(ctx, data, firestore.MergeAll)
	return err
}

func (fb *FireBaseClient) Doc(id string) *FireBaseClient {
	fb.DocumentRef = fb.CollectionRef.Doc(id)
	return fb
}
func (fb *FireBaseClient) Documents(ctx context.Context) *firestore.DocumentIterator {
	res := fb.CollectionRef.Documents(ctx)
	return res
}

func (fb *FireBaseClient) Delete(ctx context.Context) error {
	_, err := fb.DocumentRef.Delete(ctx)
	return err
}
