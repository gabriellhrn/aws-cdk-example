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

	// Test if the number of resources is right.
	template.ResourceCountIs(jsii.String("AWS::SNS::Topic"), jsii.Number(1))
	template.ResourceCountIs(jsii.String("AWS::SNS::Subscription"), jsii.Number(1))
	template.ResourceCountIs(jsii.String("AWS::SQS::Queue"), jsii.Number(2))
	template.ResourceCountIs(jsii.String("AWS::S3::Bucket"), jsii.Number(2))
	template.ResourceCountIs(jsii.String("AWS::S3::BucketPolicy"), jsii.Number(1))

	// Test if the SNS Topic has the right values and has an SQS subscription.
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

	// Test if the FIFO Queue has the right values.
	template.HasResourceProperties(jsii.String("AWS::SQS::Queue"), map[string]interface{}{
		"FifoQueue": true,
	})

	// Test if the Standard Queue has the right values.
	template.HasResourceProperties(jsii.String("AWS::SQS::Queue"), map[string]interface{}{
		"FifoQueue":         assertions.Match_Absent(),
		"VisibilityTimeout": 300,
	})

	// Test if there is one bucket NOT configured for static hosting.
	template.HasResourceProperties(jsii.String("AWS::S3::Bucket"), map[string]interface{}{
		"WebsiteConfiguration": assertions.Match_Absent(),
	})

	// Test if there is one bucket configured for static hosting.
	bucketPolicyCapture := assertions.NewCapture(assertions.Match_ObjectLike(
		&map[string]interface{}{
			"Ref": "WebsiteBucket75C24D94",
		},
	))

	template.HasResourceProperties(jsii.String("AWS::S3::Bucket"), map[string]interface{}{
		"WebsiteConfiguration": map[string]string{
			"IndexDocument": "index.html",
		},
	})

	template.HasResourceProperties(jsii.String("AWS::S3::BucketPolicy"), map[string]interface{}{
		"Bucket": bucketPolicyCapture,
	})
}
