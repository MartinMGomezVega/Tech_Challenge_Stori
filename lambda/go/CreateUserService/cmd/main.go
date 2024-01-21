package main

import (
	"github.com/MartinMGomezVega/Tech_Challenge_Stori/lambda/go/CreateUserService/internal/processor"
	"github.com/MartinMGomezVega/Tech_Challenge_Stori/lambda/go/CreateUserService/internal/repository"
	"github.com/MartinMGomezVega/Tech_Challenge_Stori/lambda/go/CreateUserService/pkg/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	r := repository.NewCreateUser()
	p := processor.New(r)
	handler := handler.New(p)
	lambda.Start(handler)
}
