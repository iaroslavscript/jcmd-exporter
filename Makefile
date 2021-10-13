.PHONY: all build clean run

all: build run

build:
	go build

clean:
	rm m

run: build
	./m
