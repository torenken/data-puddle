import { Duration, RemovalPolicy } from 'aws-cdk-lib';
import { AccessLogFormat, ApiKey, EndpointType, LogGroupLogDestination, RestApi } from 'aws-cdk-lib/aws-apigateway';
import { SnsAction } from 'aws-cdk-lib/aws-cloudwatch-actions';
import { LogGroup, RetentionDays } from 'aws-cdk-lib/aws-logs';
import { ITopic } from 'aws-cdk-lib/aws-sns';
import { Construct } from 'constructs';

export interface DataPuddleApiProps {
  readonly alarmNotification: ITopic;
}

export class DataPuddleApi extends RestApi {
  constructor(scope: Construct, id: string, props: DataPuddleApiProps) {
    super(scope, id, {
      endpointConfiguration: {
        types: [EndpointType.REGIONAL],
      },
      deployOptions: {
        stageName: 'dev',
        accessLogDestination: new LogGroupLogDestination(new LogGroup(scope, 'AccessLog', {
          retention: RetentionDays.THREE_MONTHS,
          removalPolicy: RemovalPolicy.DESTROY,
        })),
        accessLogFormat: AccessLogFormat.jsonWithStandardFields(),
        throttlingRateLimit: 10,
        throttlingBurstLimit: 5,
      },
    });

    const serverAlarm = this.metricServerError({ period: Duration.minutes(1) })
      .createAlarm(this, 'ApiMetrics5xAlarm', {
        alarmName: 'DataPuddleApiMetrics5xAlarm',
        threshold: 1,
        evaluationPeriods: 2,
      });
    serverAlarm.addAlarmAction(new SnsAction(props.alarmNotification));

    const clientAlarm = this.metricClientError({ period: Duration.minutes(5) })
      .createAlarm(this, 'ApiMetrics4xAlarm', {
        alarmName: 'DataPuddleApiMetrics4xAlarm',
        threshold: 3,
        evaluationPeriods: 1,
      });
    clientAlarm.addAlarmAction(new SnsAction(props.alarmNotification));

    const usagePlan = this.addUsagePlan('DataPuddleUsagePlan', {
      throttle: {
        rateLimit: 10,
        burstLimit: 5,
      },
      apiStages: [{
        api: this,
        stage: this.deploymentStage,
      }],
    });

    usagePlan.addApiKey(new ApiKey(this, 'DataPuddleApiKey'));
  }
}

