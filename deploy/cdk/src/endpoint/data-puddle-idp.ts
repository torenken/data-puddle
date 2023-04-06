import { RemovalPolicy } from 'aws-cdk-lib';
import { OAuthScope, ResourceServerScope, UserPool } from 'aws-cdk-lib/aws-cognito';
import { Construct } from 'constructs';

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