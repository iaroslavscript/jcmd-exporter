.PHONY: all build clean run

all: build run

build:
	go build

clean:
	rm jcmd-exporter

run: build
	./jcmd-exporter
