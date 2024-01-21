package handler

import (
	"context"
	"log"

	"github.com/MartinMGomezVega/Tech_Challenge_Stori/lambda/go/CreateUserService/internal/processor"
	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
	processor *processor.Processor
}

// New: creates a new instance of the Handler.
func New(p *processor.Processor) *Handler {
	return &Handler{processor: p}
}

func (h *Handler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	log.Printf("<start> <handler> <HandleRequest> Creating the user...")

	res := h.processor.Process(ctx, request)

	return res, nil
}
