Data Puddle Setup Steps
---

## Requirements
* Unix System environment
* Configured [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html) with administrator permissions
* [Golang 1.20 or later](https://golang.org)
* [Node.js LTS](https://nodejs.org/en/download/)
* [Yarn - Package Manager](https://yarnpkg.com/)

## CDK Bootstrap
To use CDK, the AWS accounts must be prepared. See https://docs.aws.amazon.com/cdk/v2/guide/bootstrapping.html

```shell
$ export CDK_NEW_BOOTSTRAP=1
$ npx cdk@latest bootstrap \
    --profile <your-profile-name> \
    --cloudformation-execution-policies arn:aws:iam::aws:policy/AdministratorAccess \
    aws://<your-account-number>/eu-central-1
```

## Create CDK Project
[Projen](https://github.com/projen/projen) is used for the AWS CDK project generation.

### CDK Project Generation
The CDK project is created in the deploy/cdk directory using the projen tool. To generate the project,
the progen configuration ([.projenrc.js](.projenrc.js)) must first be copied to the /deploy/cdk directory.

Then the following shell commands should be executed.

```shell
$ npx projen@latest --no-post
$ yarn (sometimes this has to be done several times)
$ yarn build 
```
