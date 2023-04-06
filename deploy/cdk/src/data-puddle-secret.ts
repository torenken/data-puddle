import { RemovalPolicy, SecretValue } from 'aws-cdk-lib';
import { Secret } from 'aws-cdk-lib/aws-secretsmanager';
import { Construct } from 'constructs';

export class DataPuddleSecret extends Secret {
  constructor(scope: Construct, id: string) {
    super(scope, id, {
      removalPolicy: RemovalPolicy.DESTROY,
      secretObjectValue: {
        secretStringValue: SecretValue.unsafePlainText('Add a 32bit encryption key as base64 encoding later'),
      },
    });
  }
}