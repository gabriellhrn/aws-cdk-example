# AWS CDK Example

This repository contains a sample CDK stack for my blog post about testing with
CDK. The provided Dockerfile can be used to run the tests or for use with the
CDK toolkit.

Deploying this stack is possible after configuring AWS and bootstrapping CDK, but
this is not covered in this README.
You can check the [CDK docs](https://docs.aws.amazon.com/cdk/v2/guide/getting_started.html)
for details.

## Testing locally

You don't need an AWS account to run these tests locally or even to print out the
resulting CloudFormation template. First, you need to build the docker image:

```bash
$ make docker-build
```

Then you can run the [unit tests](./app/app_test.go) with:

```bash
$ make test
```

To print the CloudFormation template for this sample stack, run:

```bash
$ make synth
```

## License

This code is licensed under the Apache-2.0 license. See [LICENSE](./LICENSE).
