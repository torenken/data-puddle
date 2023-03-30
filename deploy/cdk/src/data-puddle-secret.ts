import * as crypto from 'crypto';
import { RemovalPolicy, SecretValue } from 'aws-cdk-lib';
import { Secret } from 'aws-cdk-lib/aws-secretsmanager';
import { Construct } from 'constructs';

export class DataPuddleSecret extends Secret {
  constructor(scope: Construct, id: string) {
    super(scope, id, {
      removalPolicy: RemovalPolicy.DESTROY,
      secretObjectValue: {
        encryption_key: SecretValue.unsafePlainText(crypto.randomBytes(32).toString('ascii')),
      },
    });
  }
}