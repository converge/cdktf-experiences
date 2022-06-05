package main

import (
	"cdk.tf/go/stack/generated/hashicorp/aws"
	"cdk.tf/go/stack/generated/hashicorp/aws/ec2"
	"cdk.tf/go/stack/generated/hashicorp/aws/s3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// NewMyStack bundles Terraform plan
func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	// instantiate jsii
	stack := cdktf.NewTerraformStack(scope, &id)

	// create a provider
	aws.NewAwsProvider(stack, jsii.String("AWS"), &aws.AwsProviderConfig{
		Region: jsii.String("eu-central-1"),
	})

	// create an instance
	ec2Instance := ec2.NewInstance(stack, jsii.String("compute"), &ec2.InstanceConfig{
		// debian ami
		Ami:          jsii.String("ami-0245697ee3e07e755"),
		InstanceType: jsii.String("t2.micro"),
	})

	// EC2 outputs
	cdktf.NewTerraformOutput(stack, jsii.String("public_ip"), &cdktf.TerraformOutputConfig{
		Value: ec2Instance.PublicIp(),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("cpu counts"), &cdktf.TerraformOutputConfig{
		Value: ec2Instance.CpuCoreCount(),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("ami name"), &cdktf.TerraformOutputConfig{
		Value: ec2Instance.Ami(),
	})

	// creates S3 bucket
	bucket := s3.NewS3Bucket(stack, jsii.String("S3 Bucket 1"), &s3.S3BucketConfig{
		Bucket: jsii.String("unique-bucket-name-br-123"),
		Acl:    jsii.String("private"),
	})

	// S3 outputs
	cdktf.NewTerraformOutput(stack, jsii.String("Bucket region"), &cdktf.TerraformOutputConfig{
		Value: bucket.Region(),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("Bucket ARN"), &cdktf.TerraformOutputConfig{
		Value: bucket.Arn(),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdktf-exps")
	//cdktf.NewRemoteBackend(app, &cdktf.RemoteBackendProps{
	//	Hostname: jsii.String("domain.cloud"),
	//	Organization: jsii.String("JP"),
	//	Workspaces: cdktf.NewNamedRemoteWorkspace(jsii.String("cdktf-exps")),
	//})

	app.Synth()
}
