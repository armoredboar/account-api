[![Build Status](https://travis-ci.org/armoredboar/account-api.svg?branch=master)](https://travis-ci.org/armoredboar/account-api)
[![GitHub issues](https://img.shields.io/github/issues/armoredboar/account-api.svg)](https://github.com/armoredboar/account-api/issues)

# account-api

## Build with [Go](https://golang.org/) 1.11.1

### Install package dependencies

Navigate to the project folder and execute the following command to install all the external dependencies:

```
go get ./...
```

### Setting the environment variables

Before running the project some environment variables must be set.

- Configure the **noreply** email variables.

```
AB_EMAIL_NOREPLY = 'your@email.com'
AB_EMAIL_NOREPLY_NAME = 'display name of your email'
AB_EMAIL_NOREPLY_PASSWORD = 'password of the email account'
AB_EMAIL_SMTP = 'smtp server of your email account'
AB_EMAIL_PORT = 'port of your smtp server'
```

### Running the project

To initialize the service use the following command.

```
go run ./cmd/account-api/main.go
```

By default the service will be listening and serving HTTP on port 3500.

### Building the executable

Please build the executable to the bin folder to avoid commits of the binary files and run them from that folder.

```
go build -o bin/account-api cmd/account-api
./bin/account-api
```