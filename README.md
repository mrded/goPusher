# goPusher
(SSE) Server-Sent Event Service

## Dependencies

- [GO](https://golang.org) programming language;
- [gb](http://getgb.io), the project based build tool for GO;

## Usage

- `make vendor` to download all required libraries;
- `make build` to build the project;
- `make serve` to run the server;
- `make kill` to kill the server;
- `make restart` to restart the server;

## Conversion API

    curl -HContent-Type:application/json -d '{"id":"1234","event":"comment","data":"hello world"}' http://localhost:9090/events
