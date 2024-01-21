package processor

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/MartinMGomezVega/Tech_Challenge_Stori/awsgo"
	"github.com/MartinMGomezVega/Tech_Challenge_Stori/bd"
	"github.com/MartinMGomezVega/Tech_Challenge_Stori/lambda/go/CreateUserService/internal/repository"
	"github.com/MartinMGomezVega/Tech_Challenge_Stori/models"
	"github.com/MartinMGomezVega/Tech_Challenge_Stori/secretmanager"

	"github.com/aws/aws-lambda-go/events"
)

type Processor struct {
	r repository.ICreateUser
}

type IProcessor interface {
	Process(ctx context.Context, request events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse
}

func (p Processor) Process(ctx context.Context, request events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse {
	var res *events.APIGatewayProxyResponse
	awsgo.InitialiseAWS()

	if !ValidateParameter() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error in environment variables. must include 'SecretName', 'BucketName",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error in reading Secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res
	}

	path := strings.Replace(request.PathParameters["techchallengego"], os.Getenv("UrlPrefix"), "", -1)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	if (awsgo.Ctx).Value(models.Key("path")).(string) != "uploadTransactionFile" {
		log.Println("Conectando a la DB...")
		// Chequeo Conexión a la BD o Conecto la BD
		err = bd.ConectBD(awsgo.Ctx)
		if err != nil {
			res = &events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Error connecting to DB: " + err.Error(),
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			}
			return res
		}
	}

	respAPI := p.r.CreateUser(awsgo.Ctx, request)

	if respAPI.CustomResp == nil {
		headersResp := map[string]string{
			"Content-Type": "application/json",
		}
		res = &events.APIGatewayProxyResponse{
			StatusCode: respAPI.Status,
			Body:       string(respAPI.Message),
			Headers:    headersResp,
		}

		return res
	} else {
		return respAPI.CustomResp
	}
}

func ValidateParameter() bool {
	_, parameter := os.LookupEnv("SecretName")
	if !parameter {
		return parameter
	}

	_, parameter = os.LookupEnv("BucketName")
	if !parameter {
		return parameter
	}

	_, parameter = os.LookupEnv("UrlPrefix")
	if !parameter {
		return parameter
	}

	return parameter
}
