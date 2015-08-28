# goPusher
(SSE) Server-Sent Event Service

## Dependencies

- [go](https://golang.org), the programming language;
- [gb](http://getgb.io), the project based build tool for GO;

## Usage

- `make vendor` to download all required libraries;
- `make build` to build the project;
- `make serve` to run the server;
- `make kill` to kill the server;
- `make restart` to restart the server;

## Conversion API

    curl -X POST -d 'hello world' http://localhost:9090/event/:event/:id
