package repository

import (
	"context"
	"encoding/json"
	"log"

	"github.com/MartinMGomezVega/Tech_Challenge_Stori/bd"
	"github.com/MartinMGomezVega/Tech_Challenge_Stori/models"
	"github.com/aws/aws-lambda-go/events"
)

func CreateUser(ctx context.Context, request events.APIGatewayProxyRequest) models.ResposeAPI {
	var u models.User
	var r models.ResposeAPI
	r.Status = 400

	log.Println("Creating the user...")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &u)
	if err != nil {
		r.Message = err.Error()
		log.Println(r.Message)
		return r
	}

	if len(u.Email) == 0 {
		r.Message = "User email is required."
		log.Println(r.Message)
		return r
	}
	if len(u.Cuil) == 0 {
		r.Message = "The cuil is required."
		log.Println(r.Message)
		return r
	}

	_, encontrado, _ := bd.CheckUserExists(u.Email)
	if encontrado {
		r.Message = "The user already exists in the DB."
		log.Println(r.Message)
		return r
	}

	_, status, err := bd.InsertUser(u)
	if err != nil {
		r.Message = "An error occurred while trying to perform user registration. Error: " + err.Error()
		log.Println(r.Message)
		return r
	}

	if !status {
		r.Message = "Failed to insert user record."
		log.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "The user was successfully created."
	log.Println(r.Message)
	return r
}
