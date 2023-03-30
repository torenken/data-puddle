import { Topic } from 'aws-cdk-lib/aws-sns';
import { EmailSubscription } from 'aws-cdk-lib/aws-sns-subscriptions';
import { Construct } from 'constructs';

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