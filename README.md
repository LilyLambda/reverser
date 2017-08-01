# reverser

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