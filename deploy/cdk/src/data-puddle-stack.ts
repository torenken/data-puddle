import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';

export interface DataPuddleStackProps extends StackProps {
  readonly emailAddresses: string[];
}

export class DataPuddleStack extends Stack {
  constructor(scope: Construct, id: string, props: DataPuddleStackProps) {
    super(scope, id, props);

    //secret & notification

    //data-buckets

    //functions
  }
}