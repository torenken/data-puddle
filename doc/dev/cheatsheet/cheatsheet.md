Cheat sheet for coding session
---
### projenrc.js
.projenrc.js

```js
context: {
'technicalStakeholders': ['test@example.com'], //todo change email
'@aws-cdk/core:newStyleStackSynthesis': true,
...        
}
```

### data-puddle-stack
src/data-puddle-stack.ts

```ts
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
```

### data-puddle-app
src/data-puddle-app.ts

```ts
const app = new App();

Tags.of(app).add('domain', 'demo');
Tags.of(app).add('owner', 'owner');

const technicalStakeholders = app.node.tryGetContext('technicalStakeholders');

new DataPuddleStack(app, 'DataPuddleStack', {
  emailAddresses: technicalStakeholders,
});

app.synth();
```

### cdk-test-helper
test/cdk-test-helper.ts

```ts
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
```

### data-puddle-stack.test.ts
test/data-puddle-stack.test.ts

```ts
test('DataPuddleStackSnapshotTest', () => {
  const app = new TestApp();
  const stack = new TestDataPuddleStack(app, 'DataPuddleStack');

  const template = Template.fromStack(stack);

  //template.hasResource('AWS::S3::Bucket', {});
  //template.hasResource('AWS::SNS::Topic', {});

  //template.hasResource('AWS::ApiGateway::RestApi', {});

  //template.hasResource('AWS::Lambda::Function', {});

  expect(template.toJSON()).toMatchSnapshot();
});
```

### Topic
src/technical-notification.ts

```ts
export interface TechnicalNotificationProps {
  readonly emailAddresses: string[];
}

export class TechnicalNotification extends Topic {
  constructor(scope: Construct, id: string, props: TechnicalNotificationProps) {
    super(scope, id);
    props.emailAddresses.forEach((email ) => {
      this.addSubscription(new EmailSubscription(email));
    });
  }
}
```

### Bucket
src/data-puddle-bucket.ts

```ts
export interface DataPuddleBucketProps {
  readonly bucketName: string;
}

export class DataPuddleBucket extends Bucket {
  constructor(scope: Construct, id: string, props: DataPuddleBucketProps) {
    super(scope, id, {
      bucketName: props.bucketName,
      encryption: BucketEncryption.KMS_MANAGED,
      publicReadAccess: false,
      blockPublicAccess: BlockPublicAccess.BLOCK_ALL,
      removalPolicy: RemovalPolicy.DESTROY,
      autoDeleteObjects: true,
    });
  }
}
```

### Handler
src/data-puddle-handler.ts

```ts
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
```

### Secret
src/data-puddle-secret.ts

```ts
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
```

### provide-ticket-data
app/services/provide-ticket-data/main.go

```go
package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handle(ctx context.Context, s3e events.S3Event) error {
	log.Printf("it works! fine: %v", s3e)
	return nil
}

func main() {
	lambda.Start(handle)
}
```