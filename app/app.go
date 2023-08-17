package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type AppStackProps struct {
	awscdk.StackProps
}

func NewAppStack(scope constructs.Construct, id string, props *AppStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	createBucket(stack, "AppBucket", nil)
	createWebsiteBucket(stack, "WebsiteBucket", nil)

	queue := createQueue(stack, "AppQueue", &awssqs.QueueProps{
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	})

	topic := awssns.NewTopic(stack, jsii.String("AppTopic"), &awssns.TopicProps{})
	topic.AddSubscription(awssnssubscriptions.NewSqsSubscription(
		queue,
		&awssnssubscriptions.SqsSubscriptionProps{},
	))

	createFifoQueue(stack, "FifoQueue", nil)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewAppStack(app, "AppStack", &AppStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String("736270483889"),
		Region:  jsii.String("eu-west-1"),
	}
}

func createBucket(stack awscdk.Stack, id string, props *awss3.BucketProps) awss3.Bucket {
	return awss3.NewBucket(stack, jsii.String(id), props)
}

func createWebsiteBucket(stack awscdk.Stack, id string, props *awss3.BucketProps) awss3.Bucket {
	if props == nil {
		props = &awss3.BucketProps{
			WebsiteIndexDocument: jsii.String("index.html"),
			PublicReadAccess:     jsii.Bool(true),
		}
	}

	return createBucket(stack, id, props)
}

func createQueue(stack awscdk.Stack, id string, props *awssqs.QueueProps) awssqs.Queue {
	return awssqs.NewQueue(stack, jsii.String(id), props)
}

func createFifoQueue(stack awscdk.Stack, id string, props *awssqs.QueueProps) awssqs.Queue {
	if props == nil {
		props = &awssqs.QueueProps{}
	}

	props.Fifo = jsii.Bool(true)

	return awssqs.NewQueue(stack, jsii.String(id), props)
}
