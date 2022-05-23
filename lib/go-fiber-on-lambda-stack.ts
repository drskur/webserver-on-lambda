import {aws_lambda, Stack, StackProps} from 'aws-cdk-lib';
import { Construct } from 'constructs';
import {Architecture, Code, Runtime} from "aws-cdk-lib/aws-lambda";

export class GoFiberOnLambdaStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    new aws_lambda.Function(this, "go-fiber-function", {
      runtime: Runtime.PROVIDED_AL2,
      handler: 'bootstrap',
      architecture: Architecture.ARM_64,
      code: Code.fromAsset('./go-fiber/dist'),
    });
  }
}
