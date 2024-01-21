package main

import (
	"log"

	"github.com/MartinMGomezVega/Tech_Challenge_Stori/lambda/go/CreateUserService/pkg/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.Println("Start main")

	h := handler.New()
	lambda.Start(h)

	log.Println("end main")
}
