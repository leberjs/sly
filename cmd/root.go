package cmd

import (
	slog "log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/leberjs/sly/internal/aws"
	"github.com/spf13/cobra"
)

var (
	pn string
	sn string
	db bool

	rc = &cobra.Command{
		Use:   "sly",
		Short: "An AWS CloudFormation Stack Analyzer",
	}
)

func Execute() {
	if err := rc.Execute(); err != nil {
		log.Fatal("failed to execute program", err)
	}
}

func init() {
	rc.Flags().StringVar(&pn, "profile", "", "aws profile name")
	rc.Flags().StringVar(&sn, "stack-name", "", "aws cfn stack name (required)")
	rc.Flags().BoolVar(&db, "debug", false, "pass this flag to write logs to debug.log")
	rc.MarkFlagRequired("stack-name")

	rc.Run = func(_ *cobra.Command, _ []string) {
		var logFile *os.File

		if pn == "" {
			val, ok := os.LookupEnv("AWS_PROFILE")
			if !ok || val == "" {
				log.Fatal("Please provide profile name using --profile arg or AWS_PROFILE env var")
			} else {
				pn = val
			}
		}

		if db {
			var fileErr error
			newLogFile, fileErr := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if fileErr == nil {
				log.SetOutput(newLogFile)
				log.SetTimeFormat(time.Kitchen)
				log.SetReportCaller(true)
				log.SetLevel(log.DebugLevel)
				log.Debug("Starting debug.log")
			} else {
				logFile, _ = tea.LogToFile("debug.log", "debug")
				slog.Print("Logging setup failed", fileErr)
			}
		}

		if logFile != nil {
			defer logFile.Close()
		}

		ctx := aws.NewAwsContext(pn)

		ctx.GetStackResources(&sn)
	}
}
