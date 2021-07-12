# on-cancel-connection-go-vs-node

Testing golang and node.js http request context-cancellation behavior

# How to run:

## Go:

1. `go run main.go`
2. `curl localhost:8080/cancel`

## Node.js:

1. `cd js`
2. `npm install`
3. `npm run start` / `node server.js`
4. `curl localhost:3000/cancel`