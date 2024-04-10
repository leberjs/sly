package cmd

import (
	"fmt"
	"os"

	"github.com/leberjs/sly/internal/aws"
	"github.com/spf13/cobra"
)

var (
	pn string
	sn string

	rc = &cobra.Command{
		Use:   "sly",
		Short: "An AWS CloudFormation Stack Analyzer",
	}
)

func Execute() {
	if err := rc.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rc.Flags().StringVar(&pn, "profile", "", "aws profile name")
	rc.Flags().StringVar(&sn, "stack-name", "", "aws cfn stack name (required)")
	rc.MarkFlagRequired("stack-name")

	if pn == "" {
		val, ok := os.LookupEnv("AWS_PROFILE")
		if !ok || val == "" {
			fmt.Fprintln(os.Stderr, "Please provide profile name using --profile arg or AWS_PROFILE env var")
			os.Exit(1)
		} else {
			pn = val
		}
	}

	rc.Run = func(_ *cobra.Command, _ []string) {
		_ = aws.NewAwsContext(pn)

		fmt.Printf("PN - %s -- SN - %s", pn, sn)
	}
}
