package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/charmbracelet/log"
)

type AwsContext struct {
	config    aws.Config
	CfnClient *cloudformation.Client
}

func NewAwsContext(profile string) AwsContext {
	config, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile(profile),
	)
	if err != nil {
		log.Fatal("Issue getting AWS Shared Config Profile:", profile)
	}

	cfnClient := NewCfnClient(config)

	return AwsContext{config: config, CfnClient: cfnClient}
}
