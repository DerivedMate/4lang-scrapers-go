# Golang images scraper
This project is one of the scrapers built by me in 4 different languages.

## Pros :thumbsup:
### Speed
Golang is one of languages that compile to native binaries. This fact makes it extremely fast (compared to interpreted languages such as python).
### Easy Async
Golang thanks to it's [channels](https://tour.golang.org/concurrency/2) allows for easy concurrency and handling of async operation (here: fetching html data). 
### Documentation
The documentation of golang is full of [examples](https://tour.golang.org/) which allow you to understand the core concepts of the language in a few minutes.
### Core Lib
The core library of golang contains most of the packages one will ever use. As a proof of that, this project uses only one [external library](http://go-colly.org/).

## Cons :thumbsdown:
### Core Library
Although the core library is indeed very useful, it also makes for enormous file sizes. Even though the source code is only 100 lines long, the compiled binary weighs **13.8MB**.