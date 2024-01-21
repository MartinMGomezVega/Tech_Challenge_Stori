package main

import (
	"log"

	"github.com/MartinMGomezVega/Tech_Challenge_Stori/lambda/go/UploadTransactionFile/pkg/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.Println("Start main")

	lambda.Start(handler.Handlers)

	log.Println("end main")
}
