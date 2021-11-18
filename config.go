package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
)

type SocketConfig struct {
	Bind string
}

type MuxConfig struct {
	Url string
}

type GlobalConfig struct {
	SocketConfig
	MuxConfig
	IncludeDir []string
}

type TaskConfig struct {
	SocketConfig
	MuxConfig
}

type ApplicationConfig struct {
	GlobalConfig
	Tasks []TaskConfig
}

func (a *ApplicationConfig) Merge(rhs *ApplicationConfig) {
}

// See https://www.terraform.io/docs/language/values/variables.html#variable-definition-precedence
//
//Default values
//Environment variables
//Configuration file
//CLI arguments

//Search config file at default location /etc/..
//Search config file at location from ENV
//Search config file at location from CLI
//Read config file from stdin if CLI == "-"

func NewApplicationConfigFromJson(data []byte) (*ApplicationConfig, error) {

	var err error = nil

	a := &ApplicationConfig{}

	if !json.Valid([]byte(data)) {

		err = errors.New("invalid json data")
	} else {

		err = json.Unmarshal(data, a)
	}

	return a, err
}

func ReadFileFromPath(path string) ([]byte, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	return readBytes(file)
}

func ReadFileFromStdin() ([]byte, error) {

	return readBytes(os.Stdin)
}

func readBytes(file *os.File) ([]byte, error) {
	const MAX_FILE_SIZE int64 = 1024 * 1024 * 1024
	const SECKTOR_SIZE int = 2 * 1024

	r := bufio.NewReaderSize(io.LimitReader(file, MAX_FILE_SIZE), SECKTOR_SIZE)

	data, err := r.ReadBytes(0) // Read until EOF

	if err == io.EOF {
		err = nil
	}

	return data, err
}
