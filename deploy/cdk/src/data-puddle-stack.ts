import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { DataPuddleBucket } from './data-puddle-bucket';

export interface DataPuddleStackProps extends StackProps {
  readonly emailAddresses: string[];
}

export class DataPuddleStack extends Stack {
  constructor(scope: Construct, id: string, props: DataPuddleStackProps) {
    super(scope, id, props);

    //secret & notification

    //data-buckets
    new DataPuddleBucket(this, 'CrmRawBucket', {
      bucketName: 'torenken-808-data-puddle-crm-raw-bucket',
    });

    new DataPuddleBucket(this, 'TicketOutputBucket', {
      bucketName: 'torenken-808-data-puddle-ticket-bucket',
    });

    //functions
  }
}