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


# Upstart Configuration

Ubuntu come pre-packaged with a service called Upstart. A daemon for automatically starting services on system start-up and monitoring them to ensure they are restarted if they fail.

`$ cat /etc/init/go-pusher.conf`

    description     "go-pusher"
    author          "Dmitry Demenchuk"
    
    start on (net-device-up
      and local-filesystems
      and runlevel [2345])
    
    stop on runlevel [016]
    
    respawn
    
    script
      chdir /root/goPusher
      exec ./bin/goPusher
    end script
    
`$ sudo service go-pusher start`
