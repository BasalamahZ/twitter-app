# twitter-app

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

#### Golang

You need to have [Go v1.19.4](https://golang.org/dl/) installed on your machine. Follow the [official installation guide](https://golang.org/doc/install) to install Go. Or, follow [managing installations guide](https://go.dev/doc/manage-install) to have multiple Go versions on your machine.

#### PostgreSQL

This service has dependency with PostgreSQL. For development environment, you need to have a PostgreSQL server running on your machine.

### Building

1. Once you have all the prerequisites, you can start by cloning this repository into your machine.

```sh
$ mkdir -p $GOPATH/src/github.com/alwisisva/
$ cd $GOPATH/src/github.com/alwisisva
$ git clone https://github.com/alwisisva/twitter-app.git
$ cd twitter-app
```

2. Build binaries using the `go build` command.

```sh
$ go build ./cmd/app
```

### Running

1. Execute the binary to start the service

```sh
$ ./app
```

2. Try hit one of the HTTP API in browser, open: http://127.0.0.1:8080/hello

## Directory Structure

This repository is organized with the following structure

```
twitter-app
|-- cmd      # Contains executables codes
|-- internal # Application service packages      
```
