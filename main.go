package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metricAttr struct {
	reGroup string
	name    string
	help    string
}
type metricLabelsMap map[string]metricAttr

type metricMap map[string]prometheus.Gauge

type signalHandler func(os.Signal) (bool, int)

var optBindAddr = flag.String("listen-address", ":2112", "The address to listen on for HTTP requests.")
var optMainClass = flag.String("main-class", "SingleThread", "The main class of Java application.")

func NewMetricMap(m *metricLabelsMap) *metricMap {

	mm := make(metricMap, len(*m))

	metricsNamespace := "jcmd"
	metricsSubsystem := "native_memory"

	for _, attr := range *m {
		gauge := promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      attr.name,
			Help:      attr.help,
		})

		gauge.Set(0.0)
		mm[attr.reGroup] = gauge
	}

	return &mm
}

func call_jcmd(mainClass string) string {
	app := "jcmd"
	arg1 := "VM.native_memory"

	cmd := exec.Command(app, mainClass, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(stdout)
}

func parse_response(s string, p *regexp.Regexp, m *metricMap) {

	if matches := p.FindStringSubmatch(s); matches != nil {

		var v string
		var f float64
		var err error

		for group_name, metric := range *m {

			v = matches[p.SubexpIndex(group_name)]

			if f, err = strconv.ParseFloat(v, 64); err == nil {

				if strings.HasSuffix(group_name, "kb") {
					metric.Set(f * 1024)
				} else {
					metric.Set(f)
				}
			} else {
				fmt.Printf("ERROR can not convert to float %s=%s\n", group_name, v)
			}

		}

	} else {
		fmt.Println("\tnot matched TOTAL")
	}
}

func regestrySignalHandler(handlers map[os.Signal]signalHandler) {
	c := make(chan os.Signal, 1)

	signals := make([]os.Signal, 0, len(handlers))

	for k := range handlers {
		signals = append(signals, k)
	}

	signal.Notify(c, signals...)

	go func() {
		for {
			s := <-c
			fmt.Printf("INFO cache signal %v\n", s)

			if handler, ok := handlers[s]; ok {

				exit_required, exit_code := handler(s)
				if exit_required {
					os.Exit(exit_code)
				}
			} else {
				err := fmt.Errorf("got unknown signal %v", s)
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	}()
}

func cleanup(s os.Signal) (bool, int) {
	return true, 0
}

func reloadConfig(s os.Signal) (bool, int) {
	return true, 0
}

func main() {

	regestrySignalHandler(map[os.Signal]signalHandler{
		syscall.SIGINT:  cleanup,
		syscall.SIGTERM: cleanup,
		syscall.SIGHUP:  reloadConfig,
	})

	flag.Parse()

	metrics := NewMetricMap(NewMetricLabelsMap())
	pattern := NewPattern()

	stdout := call_jcmd(*optMainClass)

	parse_response(stdout, pattern, metrics)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*optBindAddr, nil))
}
