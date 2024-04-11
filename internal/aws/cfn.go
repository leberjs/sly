package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/charmbracelet/log"
)

func NewCfnClient(cfg aws.Config) *cloudformation.Client {
	client := cloudformation.NewFromConfig(cfg)

	return client
}

func (ctx AwsContext) GetStackResources(stackName *string) {
	input := &cloudformation.ListStackResourcesInput{StackName: stackName}

	out, err := ctx.CfnClient.ListStackResources(context.TODO(), input)
	if err != nil {
		log.Fatal("Error getting resources for stack:", stackName, "-", err)
	}

	for _, rs := range out.StackResourceSummaries {
		fmt.Println("**", *rs.LogicalResourceId)
	}
}
