#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import {GoFiberOnLambdaStack} from '../lib/go-fiber-on-lambda-stack';

const app = new cdk.App();
new GoFiberOnLambdaStack(app, 'WebserverOnLambdaStack');