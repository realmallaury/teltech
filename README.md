# teltech

Teltech coding challange.

## Installation & Run

Make file commands:

mod - Get dependecy modules

build - Build the binary file

test - Run unit tests

race - Run data race detector

build-docker-image - Build docker image

up - Start docker container

down - Stop docker container

help - Display this help screen</code>

To run locally: <code>go run ./cmd</code>, to run as docker container first <code>make build-docker-image</code> and then <code>make up</code> to start container and <code>make down</code> to stop & cleanup.

## Technical limitaitons

The solution can be run through docker, but cache is implemented as in memory, so miltiple instance will have their own local cache instances, this can be further improved by adding second implementation that uses some distributed cache solution.

Serivce accepts values of max math.MaxFloat64 size, and for larger values it returns "value out of range" and for larger result "+Inf" and "-Inf", for this reason anwser field is formated as string. To solve this limitation, service could be improved by supporting arbitraty precision arithmetic.

Tests cover majority of basic cases, but detailed for cache and integration tests with random generated test tabels are needed. 