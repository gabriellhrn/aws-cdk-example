package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestAppStack(t *testing.T) {
	app := awscdk.NewApp(nil)
	id := "test"
	stack := NewAppStack(app, id, nil)

	template := assertions.Template_FromStack(stack, nil)
	template.ResourceCountIs(jsii.String("AWS::SNS::Topic"), jsii.Number(1))
	template.ResourceCountIs(jsii.String("AWS::SQS::Queue"), jsii.Number(1))
	template.ResourceCountIs(jsii.String("AWS::S3::Bucket"), jsii.Number(1))

	subscriptionCapture := assertions.NewCapture(assertions.Match_ObjectLike(
		&map[string]interface{}{
			"Fn::GetAtt": []string{
				"AppQueueFD3F4958",
				"Arn",
			},
		},
	))

	template.HasResourceProperties(jsii.String("AWS::SNS::Topic"), map[string]interface{}{})
	template.HasResourceProperties(jsii.String("AWS::SNS::Subscription"), map[string]interface{}{
		"Protocol": "sqs",
		"Endpoint": subscriptionCapture,
	})

	template.HasResourceProperties(jsii.String("AWS::SQS::Queue"), map[string]interface{}{
		"VisibilityTimeout": 300,
	})

	template.HasResourceProperties(jsii.String("AWS::S3::Bucket"), map[string]interface{}{})
}
