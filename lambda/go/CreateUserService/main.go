package main

import (
	"log"

	"github.com/MartinMGomezVega/Tech_Challenge_Stori/lambda/go/CreateUserService/internal/processor"
	"github.com/MartinMGomezVega/Tech_Challenge_Stori/lambda/go/CreateUserService/internal/repository"
	"github.com/MartinMGomezVega/Tech_Challenge_Stori/lambda/go/CreateUserService/pkg/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.Println("Start main")
	r := repository.NewCreateUser()
	p := processor.New(r)
	h := handler.New(p)

	lambda.Start(h)

	log.Println("end main")
}
