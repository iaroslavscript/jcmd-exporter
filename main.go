package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var optBindAddr = flag.String("listen-address", ":2112", "The address to listen on for HTTP requests.")
var optMainClass = flag.String("main-class", "SingleThread", "The main class of Java application.")

func parse_response(s string, p *regexp.Regexp, m *metricsMap) {

	if matches := p.FindStringSubmatch(s); matches != nil {

		var v string
		var fv float64
		var err error

		for group_name, metric := range *m {

			v = matches[p.SubexpIndex(group_name)]

			if fv, err = metric.ConvertFn(v); err != nil {
				log.Printf("ERROR can not convert value '%s' of group '%s' - %v\n", v, group_name, err)
				continue
			}

			(*metric.Gauge).Set(fv)
		}

	} else {
		log.Println("\tERROR Regex not matched")
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

func main() {

	flag.Parse()

	app := NewApplication(context.Background())

	regestrySignalHandler(map[os.Signal]signalHandler{
		syscall.SIGINT:  app.terminate,
		syscall.SIGTERM: app.cleanup,
		syscall.SIGHUP:  app.reloadConfig,
	})

	_ = NewMetricsMap(ParseMetricDescJson([]byte(DEFAULT_METRICS_JSON)))
	pattern := regexp.MustCompile(DEFAULT_REGEX_PATTERN)

	tasks := make([]*JcmdTask, 0)
	RunTasks(app.ctx, pattern, tasks)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*optBindAddr, nil))
}
