import { App } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { DataPuddleStack } from '../src/data-puddle-stack';

export class TestApp extends App {
  constructor() {
    super({
      context: {
        'aws:cdk:bundling-stacks': ['NoStack'], //disable bundling lambda (asset), by using dummy stack-name (=> reduce the unit-test-time. jest-booster)
        '@aws-cdk/core:newStyleStackSynthesis': 'true',
      },
    });
  }
}

export class TestDataPuddleStack extends DataPuddleStack {
  constructor(scope: Construct, id: string) {
    super(scope, id, {
      emailAddresses: ['test@example.com'],
    });
  }
}