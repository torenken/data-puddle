import { App, Tags } from 'aws-cdk-lib';
import { DataPuddleStack } from './data-puddle-stack';

const app = new App();

Tags.of(app).add('domain', 'data-puddle');
Tags.of(app).add('owner', 'torenken');

const technicalStakeholders = app.node.tryGetContext('technicalStakeholders');

new DataPuddleStack(app, 'DataPuddleStack', {
  emailAddresses: technicalStakeholders,
});

app.synth();