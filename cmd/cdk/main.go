package main

import (
	"context"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}

	client := sts.NewFromConfig(cfg)

	callerIdentity, err := client.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		panic(err)
	}

	app := awscdk.NewApp(&awscdk.AppProps{})

	stack := awscdk.NewStack(app, jsii.String("CdkDiffActionSample"), &awscdk.StackProps{
		Env: &awscdk.Environment{
			Account: callerIdentity.Account,
			Region:  &cfg.Region,
		},
	})

	awssqs.NewQueue(stack, jsii.String("Queue"), &awssqs.QueueProps{
		QueueName: jsii.String("Queue"),
	})

	app.Synth(nil)
}
