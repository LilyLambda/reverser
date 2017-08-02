# reverser - [![CircleCI](https://circleci.com/gh/LilyLambda/reverser.svg?style=svg&circle-token=eed3aa531adffca69a3011ceaf9f9b5cba953178)](https://circleci.com/gh/LilyLambda/reverser) [![Go Report Card](https://goreportcard.com/badge/github.com/lilylambda/reverser)](https://goreportcard.com/report/github.com/lilylambda/reverser) [![GoDoc](https://godoc.org/github.com/lilylambda/reverser?status.svg)](http://godoc.org/github.com/lilylambda/reverser) [![codecov](https://codecov.io/gh/lilylambda/reverser/branch/master/graph/badge.svg)](https://codecov.io/gh/lilylambda/reverser) [![GitHub version](https://badge.fury.io/gh/LilyLambda%2Freverser.svg)](https://badge.fury.io/gh/LilyLambda%2Freverser)

Revolutionary String Reversing.

Source code for a talk "Docker for Painless & Reliable Continuous Integration" given on August 2, 2017.

## Build & run with Docker

- Install Docker
- Build container:

        $ cd reverser/
        $ docker build -t reverser .

- Run container:

        $ docker run reverser 世界你好

## Develop on Mac

- Install Go, making sure to set `$GOROOT` and `$GOPATH` correctly.

- Build & run:

        $ cd reverser/
        $ go install . ./cmd/reverser
        $ reverser 世界你好

- Run tests:

        $ cd reverser/
        $ go test -v -tags="fast pure" . ./cmd/reverser