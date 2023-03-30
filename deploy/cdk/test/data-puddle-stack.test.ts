import { Template } from 'aws-cdk-lib/assertions';
import { TestApp, TestDataPuddleStack } from './cdk-test-helper';

test('DataPuddleStackSnapshotTest', () => {
  const app = new TestApp();
  const stack = new TestDataPuddleStack(app, 'DataPuddleStack');

  const template = Template.fromStack(stack);

  template.hasResource('AWS::S3::Bucket', {});
  template.hasResource('AWS::SNS::Topic', {});

  //template.hasResource('AWS::ApiGateway::RestApi', {});

  template.hasResource('AWS::Lambda::Function', {});

  expect(template.toJSON()).toMatchSnapshot();
});