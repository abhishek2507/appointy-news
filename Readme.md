# Appointy-news

This project is made in Go-lang and provides the user to create and fetch the news feed through api calls.

Its made using Go-lang and has it's Database storage in MongoDB (Atlas).

## The API Endpoints

#### It currently supports 3 Enpoints:

##### 1. Create an article (POST Request)

##### 2. Fetch all the articles (GET Request)

##### 3. Fetch a single article by ID (GET Request with and ID in URL encoding)

## Setup Development environment

Clone the repo

    git clone https://github.com/abhishek2507/appointy-news.git

Run main.go (Make sure you have the correct version of Go Running)

    go run main.go

Build main.go (Make sure you have all the dependencies)

    go build main.go

For official documentation for Go-Lang refer the link below:

### [hello](hello/) ([godoc](//godoc.org/github.com/golang/example/hello)) and [stringutil](stringutil/) ([godoc](//godoc.org/github.com/golang/example/stringutil))

    go get github.com/golang/example/hello

A trivial "Hello, world" program that uses a stringutil package.

Command [hello](hello/) covers:

- The basic form of an executable command
- Importing packages (from the standard library and the local repository)
- Printing strings ([fmt](//golang.org/pkg/fmt/))

Library [stringutil](stringutil/) covers:

- The basic form of a library
- Conversion between string and []rune
- Table-driven unit tests ([testing](//golang.org/pkg/testing/))
