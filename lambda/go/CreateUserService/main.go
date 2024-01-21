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
	p := processor.New(r)
	h := handler.New(p)
	lambda.Start(h)

	fmt.Println("Start main")

}
