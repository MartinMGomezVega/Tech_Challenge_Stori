package main

import (
	"fmt"

	"github.com/MartinMGomezVega/Tech_Challenge_Stori/lambda/go/CreateUserService/internal/processor"
	"github.com/MartinMGomezVega/Tech_Challenge_Stori/lambda/go/CreateUserService/internal/repository"
	"github.com/MartinMGomezVega/Tech_Challenge_Stori/lambda/go/CreateUserService/pkg/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	fmt.Println("Start main")
	r := repository.NewCreateUser()

	fmt.Println("new processor")
	p := processor.New(r)

	fmt.Println("new handler")
	handler := handler.New(p)

	fmt.Println("Start handler")
	lambda.Start(handler)

	fmt.Println("Start main")

}
