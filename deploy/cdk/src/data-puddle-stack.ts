import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { DataPuddleBucket } from './data-puddle-bucket';
import { DataPuddleSecret } from './data-puddle-secret';
import { TechnicalNotification } from './technical-notification';

export interface DataPuddleStackProps extends StackProps {
  readonly emailAddresses: string[];
}

export class DataPuddleStack extends Stack {
  constructor(scope: Construct, id: string, props: DataPuddleStackProps) {
    super(scope, id, props);

    //secret & notification
    new DataPuddleSecret(this, 'DataPuddleSecret');

    new TechnicalNotification(this, 'TechnicalNotification', {
      emailAddresses: props.emailAddresses,
    });


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