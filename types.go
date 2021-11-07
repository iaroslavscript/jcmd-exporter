package main

import (
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

type ConvertFunction func(string) (float64, error)

type MetricDescAttr struct {
	ReGroup string `json:"regex_group"`
	Name    string `json:"name"`
	Help    string `json:"help"`
	Convert string `json:"convert"`
}

type Metric struct {
	Gauge     *prometheus.Gauge
	ConvertFn ConvertFunction
	// TODO labels set
}

type metricsMap map[string]Metric

type signalHandler func(os.Signal) (bool, int)
