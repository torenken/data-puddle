Cheat sheet for coding session-2
---

### data-puddle-api
src/endpoint/data-puddle-api.ts

```ts
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

    const serverAlarm = this.metricServerError({period: Duration.minutes(1)})
      .createAlarm(this, 'ApiMetrics5xAlarm', {
        alarmName: 'DataPuddleApiMetrics5xAlarm',
        threshold: 1,
        evaluationPeriods: 2
      });
    serverAlarm.addAlarmAction(new SnsAction(props.alarmNotification));

    const clientAlarm = this.metricClientError({period: Duration.minutes(5)})
      .createAlarm(this, 'ApiMetrics4xAlarm', {
        alarmName: 'DataPuddleApiMetrics4xAlarm',
        threshold: 3,
        evaluationPeriods: 1
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
```

### data-puddle-idp
src/endpoint/data-puddle-idp.ts

```ts
const ticketExportScope = new ResourceServerScope({
  scopeName: 'ticket-export-url',
  scopeDescription: 'provide url for ticket export',
});

export class DataPuddleUserPool extends UserPool {
  constructor(scope: Construct, id: string) {
    super(scope, id, {
      removalPolicy: RemovalPolicy.DESTROY,
    });

    this.addDomain('DataPuddleUserPoolDomain', {
      cognitoDomain: {
        domainPrefix: 'datapuddle',
      },
    });

    const userPoolResourceServer = this.addResourceServer('DataPuddleResourceServer', {
      identifier: 'datapuddle',
      scopes: [ticketExportScope],
    });

    this.addClient('TicketExportUserPoolClient', {
      generateSecret: true,
      oAuth: {
        flows: {
          clientCredentials: true,
        },
        scopes: [
          OAuthScope.resourceServer(userPoolResourceServer, ticketExportScope),
        ],
      },
    });
  }
}
```

### data-puddle-endpoint
src/endpoint/data-puddle-endpoint.ts

```ts
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
```

### provide-ticket-url
app/services/provide-ticket-url/main.go

```go
package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("it works! fine: %v", req)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "ok",
	}, nil
}

func main() {
	lambda.Start(handle)
}
```