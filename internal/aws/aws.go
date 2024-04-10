package aws

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type AwsContext struct {
  config aws.Config
}

func NewAwsContext(profile string) AwsContext {
  config, err := config.LoadDefaultConfig(
    context.TODO(),
    config.WithSharedConfigProfile(profile),
  )
  if err != nil {
    fmt.Fprintln(os.Stderr,"Issue getting AWS Shared Config Profile:", profile)
    os.Exit(1)
  }

  return AwsContext{config: config}
}
