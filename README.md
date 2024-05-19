# brick-transfer
Brick Transfer Service

## Run Service & Worker
### run web service
`go run cmd/main.go rest`
### run worker
`go run cmd/main.go deduct-balance-transfer-worker`

`go run cmd/main.go proceed-transfer-worker`

`go run cmd/main.go reversal-transfer-worker`

## run all service & worker with docker-compose
`docker-compose up`

## Build Application
`make build` or `go build -o paper-transfer ./cmd`

## Folder structure
```
/cmd
  /app                      # Main application code (eg. rest api & worker)
  /config                   # Configuration files
/docs                       # Docs related to this service (API, etc)
/internal
  /app                      # business logic
    /domain                 # Domain entities
    /repository             # Interfaces defining repository contracts
    /usecase                # Use cases or interactors
/localdevscripts            # Contain localstack script to init SNS topic & SQS queue
/pkg
  /database                 # Database related code (repositories implementation, migrations, etc.)
    /migrations             # SQL script for migrations
  /errors                   # Errors library & handling
  /external                 # External service clients
    /client                 # Code for calling other services with HTTP client
      //bankserviceclient   # Client for a specific external service
  /http                     # HTTP delivery mechanism
    /handler                # HTTP request handlers
    /middleware             # Middleware functions
```

## Usecase Diagram
![usecase](https://github.com/fahmyabida/brick-transfer/assets/32190889/2c9ee91c-3cff-470a-9099-d050bef83cc0)

## Activity Diagram
### 1. Create Transfer
![1 create transfer](https://github.com/fahmyabida/brick-transfer/assets/32190889/ebd5f30d-6fdf-4833-9ce3-0e314ab5f395)
### 2. Deduct Balance Worker
![2 deduct balance](https://github.com/fahmyabida/brick-transfer/assets/32190889/2985f453-5145-4842-97ad-046488225277)
### 3. Proceed Transfer Worker
![3 proceed transfer](https://github.com/fahmyabida/brick-transfer/assets/32190889/0819067c-256f-43ec-b32c-d00793324602)
### 4. Receive Callback
![4 callback](https://github.com/fahmyabida/brick-transfer/assets/32190889/d9d21f3c-f6d9-4925-83d2-0214c5954e14)
### 5. Reversal Balance
![5 reversal](https://github.com/fahmyabida/brick-transfer/assets/32190889/f0ecbcdd-3f25-4b5b-80fa-052b266486a8)
### Note
Combining database transactions with message queues like AWS SQS (Simple Queue Service) can provide a powerful mechanism for handling concurrency and preventing race conditions in distributed systems

### Transfer Status
```
- ACCEPTED  : transfer received on our end
- REJECTED  : if user balance insufficient, after request the transfer
- DEDUCTED  : if balance sufficient, the user balance will be deducted
- PROCEED   : proceed the transfer after we sent to the bank
- SUCCEEDED : after callback, we received SUCCESS status from bank
- REVERSAL  : after callback, we received FAILED status from bank 
```

## Demo
### Validate Bank Account
https://github.com/fahmyabida/brick-transfer/assets/32190889/d9bd87fd-b26f-4c7f-b59a-b8e423d062c2
### Transfer & Callback - succeeded & failed
https://github.com/fahmyabida/brick-transfer/assets/32190889/a345e0c3-e111-498a-bf87-980789838d0f

## Room for improvement
- add dead letter queue when error publish message to queue & add error notification (to slack maybe)
- add auto migration / command for up and down migration script
- improve transaction code inside repository
