# goPusher

Server-Sent Events as a standalone service

## Motivation

In some languages such as PHP it's not easy to keep connection between a client and a server open.
This standalone service is made to solve that problem.

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
