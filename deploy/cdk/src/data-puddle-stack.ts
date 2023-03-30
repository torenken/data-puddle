import { Stack, StackProps } from 'aws-cdk-lib';
import { S3EventSource } from 'aws-cdk-lib/aws-lambda-event-sources';
import { EventType } from 'aws-cdk-lib/aws-s3';
import { Construct } from 'constructs';
import { DataPuddleBucket } from './data-puddle-bucket';
import { DataPuddleHandler } from './data-puddle-handler';
import { DataPuddleSecret } from './data-puddle-secret';
import { TechnicalNotification } from './technical-notification';

export interface DataPuddleStackProps extends StackProps {
  readonly emailAddresses: string[];
}

export class DataPuddleStack extends Stack {
  constructor(scope: Construct, id: string, props: DataPuddleStackProps) {
    super(scope, id, props);

    //secret & notification
    const secret = new DataPuddleSecret(this, 'DataPuddleSecret');

    new TechnicalNotification(this, 'TechnicalNotification', {
      emailAddresses: props.emailAddresses,
    });

    //data-buckets
    const crmRawBucket = new DataPuddleBucket(this, 'CrmRawBucket', {
      bucketName: 'torenken-808-data-puddle-crm-raw-bucket',
    });

    const ticketOutputBucket = new DataPuddleBucket(this, 'TicketOutputBucket', {
      bucketName: 'torenken-808-data-puddle-ticket-bucket',
    });

    //functions
    const provideTicketDataFunc = new DataPuddleHandler(this, 'ProvideTicketData', {
      serviceName: 'provide-ticket-data',
      environment: {
        CRM_RAW_BUCKET: crmRawBucket.bucketName,
        TICKET_SYS_OUT_BUCKET: ticketOutputBucket.bucketName,
        SECRET_STORE_ARM: secret.secretArn,
      },
    });
    provideTicketDataFunc.addEventSource(new S3EventSource(crmRawBucket, {
      events: [EventType.OBJECT_CREATED],
    }));

    crmRawBucket.grantRead(provideTicketDataFunc);
    ticketOutputBucket.grantWrite(provideTicketDataFunc);
  }
}