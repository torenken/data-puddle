import * as path from 'path';
import { GoFunction } from '@aws-cdk/aws-lambda-go-alpha';
import { Architecture } from 'aws-cdk-lib/aws-lambda';
import { RetentionDays } from 'aws-cdk-lib/aws-logs';
import { Construct } from 'constructs';

export interface DataPuddleHandlerProps{
  readonly serviceName: string;
  readonly environment: Record<string, string>;
}

export class DataPuddleHandler extends GoFunction {
  constructor(scope: Construct, id: string, props: DataPuddleHandlerProps) {
    super(scope, id, {
      entry: path.join(__dirname, `../../../app/services/${props.serviceName}`),
      functionName: `data-puddle-${props.serviceName}`,

      memorySize: 1024,
      logRetention: RetentionDays.THREE_MONTHS,
      architecture: Architecture.ARM_64,

      bundling: {
        goBuildFlags: ['-ldflags "-s -w"'],
        cgoEnabled: false,
      },
      environment: {
        ...props.environment,
      },
    });
  }
}