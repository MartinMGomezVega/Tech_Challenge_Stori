module github.com/MartinMGomezVega/Tech_Challenge_Stori/lambda/go/UploadTransactionService

go 1.19

replace github.com/MartinMGomezVega/Tech_Challenge_Stori/bd => ../../../../../bd

replace github.com/MartinMGomezVega/Tech_Challenge_Stori/models => ../../../../../models

replace github.com/MartinMGomezVega/Tech_Challenge_Stori/awsgo => ../../../../../awsgo

require (
	github.com/MartinMGomezVega/Tech_Challenge_Stori v0.0.0-20240121152215-dbeb4740aad8
	github.com/aws/aws-sdk-go-v2/config v1.26.5 // indirect
)

require github.com/aws/aws-sdk-go v1.50.0

require github.com/jmespath/go-jmespath v0.4.0 // indirect

require (
	github.com/aws/aws-lambda-go v1.45.0
	github.com/aws/aws-sdk-go-v2 v1.24.1 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.16.16 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.11 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.7.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.26.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.18.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.21.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.26.7 // indirect
	github.com/aws/smithy-go v1.19.0 // indirect
	go.mongodb.org/mongo-driver v1.13.1 // indirect
)
