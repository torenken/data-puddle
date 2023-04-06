import { CfnOutput } from 'aws-cdk-lib';
import { AuthorizationType, CognitoUserPoolsAuthorizer, LambdaIntegration } from 'aws-cdk-lib/aws-apigateway';
import { IBucket } from 'aws-cdk-lib/aws-s3';
import { ITopic } from 'aws-cdk-lib/aws-sns';
import { Construct } from 'constructs';
import { DataPuddleApi } from './data-puddle-api';
import { DataPuddleUserPool } from './data-puddle-idp';
import { DataPuddleHandler } from '../data-puddle-handler';

export interface DataPuddleEndpointProps {
  readonly alarmNotification: ITopic;
  readonly ticketDataBucket: IBucket;
}

export class DataPuddleEndpoint extends Construct {
  public readonly urlOutput: CfnOutput;

  constructor(scope: Construct, id: string, props: DataPuddleEndpointProps) {
    super(scope, id);

    //idp
    const userPool = new DataPuddleUserPool(this, 'DataPuddleUserPool');

    //api
    const api = new DataPuddleApi(this, 'DataPuddleApi', {
      alarmNotification: props.alarmNotification,
    });

    //authorizer
    const authorizer = new CognitoUserPoolsAuthorizer(this, 'DataPuddleUserPoolAuthorizer', {
      cognitoUserPools: [userPool],
    });

    //functions
    const provideTicketUrlFunc = new DataPuddleHandler(this, 'ProvideTicketUrlFunc', {
      serviceName: 'provide-ticket-url',
      environment: {
        TICKET_BUCKET_NAME: props.ticketDataBucket.bucketName,
      },
    });
    props.ticketDataBucket.grantRead(provideTicketUrlFunc);

    //routing
    const dataLakeResource = api.root.addResource('datalake');
    const ticketExportResource = dataLakeResource.addResource('ticket').addResource('export');
    ticketExportResource.addMethod('GET', new LambdaIntegration(provideTicketUrlFunc), {
      authorizer: authorizer,
      authorizationType: AuthorizationType.COGNITO,
      authorizationScopes: ['datapuddle/ticket-export-url'],
      apiKeyRequired: true,
    });

    this.urlOutput = new CfnOutput(this, 'Url', {
      value: api.url,
    });
  }
}