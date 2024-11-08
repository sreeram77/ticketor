# Ticketor

Ticketor is a ticketing system for trains from London to France. The train ticket will cost $20.

API where you can submit a purchase for a ticket.  

Details included in the receipt are:
- From, To, User, price paid.
- User should include first and last name, email address

The user is allocated a seat in the train.  Assume the train has only 2 sections, section A and section B.
- API that shows the details of the receipt for the user
- An API that lets you view the users and seat they are allocated by the requested section
- An API to remove a user from the train 
- An API to modify a user's seat


## Test
```bash
go test -v ./...
```

## Run Server
```bash
go run cmd/server/main.go

# or use Docker
docker build -t ticketor . && docker run -p 8000:8000 -d ticketor

```

## Run Client
```bash
go run cmd/client/main.go
```

## Proto
The proto files are available [here](./proto) and generated files are [here](./protogen)
