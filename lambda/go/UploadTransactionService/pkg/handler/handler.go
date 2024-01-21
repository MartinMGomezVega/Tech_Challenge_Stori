package handler

import (
	"context"
	"fmt"

	"github.com/MartinMGomezVega/Tech_Challenge_Stori/lambda/go/UploadTransactionService/internal/processor"
	"github.com/aws/aws-lambda-go/events"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	fmt.Print("<start> <handler> <HandleRequest> Creating the user...")

	res := processor.Process(ctx, request)

	fmt.Print("<end> <handler> <HandleRequest> End Handlers.")

	return res, nil
}
