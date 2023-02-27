service: {{.name}}
provider:
  name: aws
  runtime: go1.x
  region: {{.awsRegion}}
package:
  patterns:
    - services/**/bin/*
  individually: true
functions:
  hello:
    handler: services/hello/hello_handler
    events:
      - httpApi:
          method: GET
          path: /hello