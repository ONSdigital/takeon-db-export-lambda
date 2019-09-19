# takeon-db-export-lambda
This repository is used to do the following -
  1. access Business layer's GraphQL Endpoint that returns the whole database as Json String
  2. Saves the Json String of Step 1 into a file in S3 bucket

## Dependency
The business layer's Load Balancer should be up and running. Also it should allow access to NAT's public IP so that this Lambda can access the endpoint.

## Environment variables in Lamda
  GRAPHQL_ENDPOINT
  S3_BUCKET

## Lambda URL
  https://eu-west-2.console.aws.amazon.com/lambda/home?region=eu-west-2#/functions/takeon-validationdb-export-lambda-dev-main?tab=graph