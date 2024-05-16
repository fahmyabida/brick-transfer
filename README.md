# paper-transfer
Brick Transfer Service

## Run
### run rest API
`go run cmd/main.go rest`
## run worker
`go run cmd/main.go deduct-balance-transfer-worker`
`go run cmd/main.go proceed-transfer-worker`
`go run cmd/main.go reversal-transfer-worker`
## run All with docker-compose
`docker-compose up`


## Build
`make build` or `go build -o paper-transfer ./cmd`

## Folder structure
```
/cmd
  /app                # Main application code (eg. rest api & worker)
  /config             # Configuration files
/docs                 # Docs related to this service (API, etc)
/internal
  /app                # business logic
    /domain           # Domain entities
    /repository       # Interfaces defining repository contracts
    /usecase          # Use cases or interactors
/localdevscripts      # Contain localstack script to init SNS topic & SQS queue
/pkg
  /database           # Database related code (repositories implementation, migrations, etc.)
    /migrations       # SQL script for migrations
  /errors             # Errors library & handling
  /external           # External service clients
    /client           # Code for calling other services with HTTP client
      /bank           # Client for a specific external service
  /http               # HTTP delivery mechanism
    /handler          # HTTP request handlers
    /middleware       # Middleware functions
```

## Usecase Diagram
![Activity Diagram](https://github.com/fahmyabida/paper-transfer/assets/32190889/be98775c-a4c4-44e2-a27b-c925b7055ce7)

## Activity Diagram
### Create Transfer
![Create Transfer](https://github.com/fahmyabida/paper-transfer/assets/32190889/5349371d-3412-470c-8d33-0c51db3afabf)
### Deduct Balance Worker
![Deduct Balance Worker](https://github.com/fahmyabida/paper-transfer/assets/32190889/e9237c3b-d37b-4ef1-9e05-99bcd5358d0d)

### Transfer Status
- ACCEPTED - transfer received on our end
- REJECTED - if user balance insufficient, after request the transfer
- DEDUCTED - if balance sufficient, the user balance will be deducted
- PROCEED - proceed the transfer after we sent to the bank
- SUCCEEDED - after callback, we received SUCCESS status from bank
- REVERSAL - after callback, we received FAILED status from bank 

## Room for improvement
- add dead letter queue when error publish message to queue & add error notification (to slack maybe)
- add process / worker to listen after deducted money to disburse money to the destionation bank
- add auto migration / command for up and down migration script
- improve transaction code inside repository
