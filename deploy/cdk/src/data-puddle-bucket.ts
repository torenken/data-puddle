import { RemovalPolicy } from 'aws-cdk-lib';
import { BlockPublicAccess, Bucket, BucketEncryption } from 'aws-cdk-lib/aws-s3';
import { Construct } from 'constructs';

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