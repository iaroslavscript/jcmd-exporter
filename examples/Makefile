
.PHONY: all clean build run show

all: build run

clean:
	rm SingleThread.class

build:
	javac SingleThread.java

run:
	java -XX:NativeMemoryTracking=summary SingleThread

show:
	jcmd SingleThread VM.native_memory


