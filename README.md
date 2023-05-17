# A TCP server and client written in Go

## Server
- Binding to port and accepting connections
- Reading client messages
- Responding to clients
- On <kbd>Ctrl</kbd> + <kbd>C</kbd> closes all connections and stops reading any messages

## Client
- Connecting to server
- Sending messages
- Receiving responces
- In case of an error closes the connection (considers an error to always be the server closing the connection)

