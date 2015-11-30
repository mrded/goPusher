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

    curl -X POST -H "X-Token:secret" -d 'hello world' http://localhost:9090/event/:event/:id

# How To Set Up goPusher on Ubuntu 14.04

Create a non-root user named **go-pusher**: `sudo useradd -m -G www-data -s /bin/bash go-pusher`

Add the following lines to the end of the `/home/go-pusher/.bashrc` file:

    export GOPATH=/home/go-pusher/go
    export GOROOT=/usr/local/src/go
    export PATH=${PATH}:$GOROOT/bin:$GOPATH/bin
    
Then download latest complied version of Go from its [website](http://golang.org/dl/). And unarchive it into `$GOROOT` directory.

`sudo curl -H 'Accept-encoding:gzip' https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz | gzip -dc | tar -xf - -C /usr/local/src`

Install [GB](http://getgb.io/): `go get github.com/constabulary/gb/...`

Create a directory for the log files to live in: `sudo mkdir -p /usr/local/var/log/goPusher`

Clone source code of goPusher project into ~/goPusher folder:

`git clone https://github.com/mrded/goPusher ~/goPusher; cd ~/goPusher`

Install all dependencies and build the project: `make {vendor,build}`

# Install and Start goPusher as a Service

## Upstart Configuration

Ubuntu come pre-packaged with a service called Upstart. A daemon for automatically starting services on system start-up and monitoring them to ensure they are restarted if they fail.

`cat /etc/init/go-pusher.conf`

    description     "go-pusher"
    author          "Dmitry Demenchuk"
    
    start on (net-device-up
      and local-filesystems
      and runlevel [2345])
    
    stop on runlevel [016]
    
    respawn
    
    script
      chdir /home/go-pusher/goPusher
      exec ./bin/goPusher
    end script
    
- Run the service: `sudo service go-pusher start`
- Add the service to autoload: `sudo update-rc.d go-pusher defaults`
- Remove the service from autoload: `sudo update-rc.d -f go-pusher remove`

## Supervisor Configuration

You can also use [supervisor](http://supervisord.org/) to manage the service.

Add a configuration file for Supervisor: `cat /etc/supervisor/conf.d/go-pusher.conf`

    [program:go-pusher]
    directory=/home/go-pusher/goPusher/
    command=./bin/goPusher
    autostart=true
    autorestart=true
    startsecs=10
    stdout_logfile=/usr/local/var/log/goPusher/stdout.log
    stdout_logfile_maxbytes=1MB
    stdout_logfile_backups=10
    stdout_capture_maxbytes=1MB
    stderr_logfile=/usr/local/var/log/goPusher/stderr.log
    stderr_logfile_maxbytes=1MB
    stderr_logfile_backups=10
    stderr_capture_maxbytes=1MB
    environment = HOME="/home/go-pusher", USER="go-pusher"
    user = go-pusher

This section defines the command we want to execute to start goPusher, automatically starts it with Supervisor, and specifies the locations of log files and corresponding environment variables.

Now restart Supervisor: `sudo service supervisor restart`
