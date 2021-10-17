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

func NewMetricLabelsMap() *metricLabelsMap {
	subsystem := "native_memory"

	m := make(metricLabelsMap)

	// Total Section (to_)
	m[subsystem+"total_reserved_bytes"] = metricAttr{
		"to_resv_kb",
		"total_reserved_bytes",
		"jcmd VM.native_memory metric Reserved Bytes",
	}
	m[subsystem+"total_committed_bytes"] = metricAttr{
		"to_comm_kb",
		"total_committed_bytes",
		"jcmd VM.native_memory metric Committed Bytes",
	}

	// Java Heap section (hp_)
	m[subsystem+"java_heap_reserved_bytes"] = metricAttr{
		"hp_resv_kb",
		"java_heap_reserved_bytes",
		"jcmd VM.native_memory metric Java Heap Reserved Bytes",
	}
	m[subsystem+"java_heap_committed_bytes"] = metricAttr{
		"hp_comm_kb",
		"java_heap_committed_bytes",
		"jcmd VM.native_memory metric Java Heap Committed Bytes",
	}

	m[subsystem+"java_heap_mmap_reserved_bytes"] = metricAttr{
		"hp_mm_resv_kb",
		"java_heap_mmap_reserved_bytes",
		"jcmd VM.native_memory metric Java Heap Mmap Reserved Bytes",
	}
	m[subsystem+"java_heap_mmap_committed_bytes"] = metricAttr{
		"hp_mm_comm_kb",
		"java_heap_mmap_committed_bytes",
		"jcmd VM.native_memory metric Java Heap Mmap Committed Bytes",
	}

	// Class section (cl_)
	m[subsystem+"class_reserved_bytes"] = metricAttr{
		"cl_resv_kb",
		"class_reserved_bytes",
		"jcmd VM.native_memory metric Class Reserved Bytes",
	}
	m[subsystem+"class_committed_bytes"] = metricAttr{
		"cl_comm_kb",
		"class_committed_bytes",
		"jcmd VM.native_memory metric Class Committed Bytes",
	}
	m[subsystem+"class_classes_total"] = metricAttr{
		"cl_classes_total",
		"class_classes_total",
		"jcmd VM.native_memory metric Class Classes Total",
	}
	m[subsystem+"class_instance_classes_total"] = metricAttr{
		"cl_ic_total",
		"class_instance_classes_total",
		"jcmd VM.native_memory metric Class Instance Classes Total",
	}
	m[subsystem+"class_array_classes_total"] = metricAttr{
		"cl_ac_total",
		"class_array_classes_total",
		"jcmd VM.native_memory metric Class Array Classes Total",
	}
	m[subsystem+"class_malloc_bytes"] = metricAttr{
		"cl_malloc_kb",
		"class_malloc_bytes",
		"jcmd VM.native_memory metric Class Malloc Bytes",
	}
	m[subsystem+"class_malloc_total"] = metricAttr{
		"cl_malloc_total",
		"class_malloc_total",
		"jcmd VM.native_memory metric Class Malloc Total",
	}
	m[subsystem+"class_mmap_reserved_bytes"] = metricAttr{
		"cl_mm_resv_kb",
		"class_mmap_reserved_bytes",
		"jcmd VM.native_memory metric Class Mmap Reserved Bytes",
	}
	m[subsystem+"class_mmap_committed_bytes"] = metricAttr{
		"cl_mm_comm_kb",
		"class_mmap_committed_bytes",
		"jcmd VM.native_memory metric Class Mmap Committed Bytes",
	}

	// Class Metadata subsection (cl_me_)
	m[subsystem+"class_metadata_reserved_bytes"] = metricAttr{
		"cl_me_resv_kb",
		"class_metadata_reserved_bytes",
		"jcmd VM.native_memory metric Class Metadata Reserved Bytes",
	}
	m[subsystem+"class_metadata_committed_bytes"] = metricAttr{
		"cl_me_comm_kb",
		"class_metadata_committed_bytes",
		"jcmd VM.native_memory metric Class Metadata Committed Bytes",
	}
	m[subsystem+"class_metadata_used_bytes"] = metricAttr{
		"cl_me_used_kb",
		"class_metadata_used_bytes",
		"jcmd VM.native_memory metric Class Metadata Used Bytes",
	}
	m[subsystem+"class_metadata_free_bytes"] = metricAttr{
		"cl_me_free_kb",
		"class_metadata_free_bytes",
		"jcmd VM.native_memory metric Class Metadata Free Bytes",
	}
	m[subsystem+"class_metadata_waste_bytes"] = metricAttr{
		"cl_me_waste_kb",
		"class_metadata_waste_bytes",
		"jcmd VM.native_memory metric Class Metadata Waste Bytes",
	}

	// Class Class Space subsection (cl_cs_)
	m[subsystem+"class_class_space_reserved_bytes"] = metricAttr{
		"cl_cs_resv_kb",
		"class_class_space_reserved_bytes",
		"jcmd VM.native_memory metric Class Class Space Reserved Bytes",
	}
	m[subsystem+"class_class_space_committed_bytes"] = metricAttr{
		"cl_cs_comm_kb",
		"class_class_space_committed_bytes",
		"jcmd VM.native_memory metric Class Class Space Committed Bytes",
	}
	m[subsystem+"class_class_space_used_bytes"] = metricAttr{
		"cl_cs_used_kb",
		"class_class_space_used_bytes",
		"jcmd VM.native_memory metric Class Class Space Used Bytes",
	}
	m[subsystem+"class_class_space_free_bytes"] = metricAttr{
		"cl_cs_free_kb",
		"class_class_space_free_bytes",
		"jcmd VM.native_memory metric Class Class Space Free Bytes",
	}
	m[subsystem+"class_class_space_waste_bytes"] = metricAttr{
		"cl_cs_waste_kb",
		"class_class_space_waste_bytes",
		"jcmd VM.native_memory metric Class Class Space Waste Bytes",
	}

	// Thread section (th_)
	m[subsystem+"thread_reserved_bytes"] = metricAttr{
		"th_resv_kb",
		"thread_reserved_bytes",
		"jcmd VM.native_memory metric Thread Reserved Bytes",
	}
	m[subsystem+"thread_committed_bytes"] = metricAttr{
		"th_comm_kb",
		"thread_committed_bytes",
		"jcmd VM.native_memory metric Thread Committed Bytes",
	}
	m[subsystem+"thread_total"] = metricAttr{
		"th_total",
		"thread_total",
		"jcmd VM.native_memory metric Thread Total",
	}
	m[subsystem+"thread_stack_reserved_bytes"] = metricAttr{
		"th_s_resv_kb",
		"thread_stack_reserved_bytes",
		"jcmd VM.native_memory metric Thread Stack Reserved Bytes",
	}
	m[subsystem+"thread_stack_committed_bytes"] = metricAttr{
		"th_s_comm_kb",
		"thread_stack_committed_bytes",
		"jcmd VM.native_memory metric Thread Stack Committed Bytes",
	}
	m[subsystem+"thread_malloc_bytes"] = metricAttr{
		"th_malloc_kb",
		"thread_malloc_bytes",
		"jcmd VM.native_memory metric Thread Malloc Bytes",
	}
	m[subsystem+"thread_malloc_total"] = metricAttr{
		"th_malloc_total",
		"thread_malloc_total",
		"jcmd VM.native_memory metric Thread Malloc Total",
	}
	m[subsystem+"thread_arena_bytes"] = metricAttr{
		"th_arena_kb",
		"thread_arena_bytes",
		"jcmd VM.native_memory metric Thread Arena Bytes",
	}
	m[subsystem+"thread_arena_total"] = metricAttr{
		"th_arena_total",
		"thread_arena_total",
		"jcmd VM.native_memory metric Thread Arena Total",
	}

	// Code section (co_)
	m[subsystem+"code_reserved_bytes"] = metricAttr{
		"co_resv_kb",
		"code_reserved_bytes",
		"jcmd VM.native_memory metric Code Reserved Bytes",
	}
	m[subsystem+"code_committed_bytes"] = metricAttr{
		"co_comm_kb",
		"code_committed_bytes",
		"jcmd VM.native_memory metric Code Committed Bytes",
	}
	m[subsystem+"code_malloc_bytes"] = metricAttr{
		"co_malloc_kb",
		"code_malloc_bytes",
		"jcmd VM.native_memory metric Code Malloc Bytes",
	}
	m[subsystem+"code_malloc_total"] = metricAttr{
		"co_malloc_total",
		"code_malloc_total",
		"jcmd VM.native_memory metric Code Malloc Total",
	}
	m[subsystem+"code_mmap_reserved_bytes"] = metricAttr{
		"co_mm_resv_kb",
		"code_mmap_reserved_bytes",
		"jcmd VM.native_memory metric Code Mmap Reserved Bytes",
	}
	m[subsystem+"code_mmap_committed_bytes"] = metricAttr{
		"co_mm_comm_kb",
		"code_mmap_committed_bytes",
		"jcmd VM.native_memory metric Code Mmap Committed Bytes",
	}

	// GC section (gc_)
	m[subsystem+"gc_reserved_bytes"] = metricAttr{
		"gc_resv_kb",
		"gc_reserved_bytes",
		"jcmd VM.native_memory metric GC Reserved Bytes",
	}
	m[subsystem+"gc_committed_bytes"] = metricAttr{
		"gc_comm_kb",
		"gc_committed_bytes",
		"jcmd VM.native_memory metric GC Committed Bytes",
	}
	m[subsystem+"gc_malloc_bytes"] = metricAttr{
		"gc_malloc_kb",
		"gc_malloc_bytes",
		"jcmd VM.native_memory metric GC Malloc Bytes",
	}
	m[subsystem+"gc_malloc_total"] = metricAttr{
		"gc_malloc_total",
		"gc_malloc_total",
		"jcmd VM.native_memory metric GC Malloc Total",
	}
	m[subsystem+"gc_mmap_reserved_bytes"] = metricAttr{
		"gc_mm_resv_kb",
		"gc_mmap_reserved_bytes",
		"jcmd VM.native_memory metric GC Mmap Reserved Bytes",
	}
	m[subsystem+"gc_mmap_committed_bytes"] = metricAttr{
		"gc_mm_comm_kb",
		"gc_mmap_committed_bytes",
		"jcmd VM.native_memory metric GC Mmap Committed Bytes",
	}

	// Compiler section (cp_)
	m[subsystem+"compiler_reserved_bytes"] = metricAttr{
		"cp_resv_kb",
		"compiler_reserved_bytes",
		"jcmd VM.native_memory metric Compiler Reserved Bytes",
	}
	m[subsystem+"compiler_committed_bytes"] = metricAttr{
		"cp_comm_kb",
		"compiler_committed_bytes",
		"jcmd VM.native_memory metric Compiler Committed Bytes",
	}
	m[subsystem+"compiler_malloc_bytes"] = metricAttr{
		"cp_malloc_kb",
		"compiler_malloc_bytes",
		"jcmd VM.native_memory metric Compiler Malloc Bytes",
	}
	m[subsystem+"compiler_malloc_total"] = metricAttr{
		"cp_malloc_total",
		"compiler_malloc_total",
		"jcmd VM.native_memory metric Compiler Malloc Total",
	}
	m[subsystem+"compiler_arena_bytes"] = metricAttr{
		"cp_arena_kb",
		"compiler_arena_bytes",
		"jcmd VM.native_memory metric Compiler Arena Bytes",
	}
	m[subsystem+"compiler_arena_total"] = metricAttr{
		"cp_arena_total",
		"compiler_arena_total",
		"jcmd VM.native_memory metric Compiler Arena Total",
	}

	// Internal section (in_)
	m[subsystem+"internal_reserved_bytes"] = metricAttr{
		"in_resv_kb",
		"internal_reserved_bytes",
		"jcmd VM.native_memory metric Internal Reserved Bytes",
	}
	m[subsystem+"internal_committed_bytes"] = metricAttr{
		"in_comm_kb",
		"internal_committed_bytes",
		"jcmd VM.native_memory metric Internal Committed Bytes",
	}
	m[subsystem+"internal_malloc_bytes"] = metricAttr{
		"in_malloc_kb",
		"internal_malloc_bytes",
		"jcmd VM.native_memory metric Internal Malloc Bytes",
	}
	m[subsystem+"internal_malloc_total"] = metricAttr{
		"in_malloc_total",
		"internal_malloc_total",
		"jcmd VM.native_memory metric Internal Malloc Total",
	}
	m[subsystem+"internal_mmap_reserved_bytes"] = metricAttr{
		"in_mm_resv_kb",
		"internal_mmap_reserved_bytes",
		"jcmd VM.native_memory metric Internal Mmap Reserved Bytes",
	}
	m[subsystem+"internal_mmap_committed_bytes"] = metricAttr{
		"in_mm_comm_kb",
		"internal_mmap_committed_bytes",
		"jcmd VM.native_memory metric Internal Mmap Committed Bytes",
	}

	// Symbol section (sy_)
	m[subsystem+"symbol_reserved_bytes"] = metricAttr{
		"sy_resv_kb",
		"symbol_reserved_bytes",
		"jcmd VM.native_memory metric Symbol Reserved Bytes",
	}
	m[subsystem+"symbol_committed_bytes"] = metricAttr{
		"sy_comm_kb",
		"symbol_committed_bytes",
		"jcmd VM.native_memory metric Symbol Committed Bytes",
	}
	m[subsystem+"symbol_malloc_bytes"] = metricAttr{
		"sy_malloc_kb",
		"symbol_malloc_bytes",
		"jcmd VM.native_memory metric Symbol Malloc Bytes",
	}
	m[subsystem+"symbol_malloc_total"] = metricAttr{
		"sy_malloc_total",
		"symbol_malloc_total",
		"jcmd VM.native_memory metric Symbol Malloc Total",
	}
	m[subsystem+"symbol_arena_bytes"] = metricAttr{
		"sy_arena_kb",
		"symbol_arena_bytes",
		"jcmd VM.native_memory metric Symbol Arena Bytes",
	}
	m[subsystem+"symbol_arena_total"] = metricAttr{
		"sy_arena_total",
		"symbol_arena_total",
		"jcmd VM.native_memory metric Symbol Arena Total",
	}

	// Native Memory Tracking section (nm_)
	m[subsystem+"native_memory_tracking_reserved_bytes"] = metricAttr{
		"nm_resv_kb",
		"native_memory_tracking_reserved_bytes",
		"jcmd VM.native_memory metric Native Memory Tracking Reserved Bytes",
	}
	m[subsystem+"native_memory_tracking_committed_bytes"] = metricAttr{
		"nm_comm_kb",
		"native_memory_tracking_committed_bytes",
		"jcmd VM.native_memory metric Native Memory Tracking Committed Bytes",
	}
	m[subsystem+"native_memory_tracking_malloc_bytes"] = metricAttr{
		"nm_malloc_kb",
		"native_memory_tracking_malloc_bytes",
		"jcmd VM.native_memory metric Native Memory Tracking Malloc Bytes",
	}
	m[subsystem+"native_memory_tracking_malloc_total"] = metricAttr{
		"nm_malloc_total",
		"native_memory_tracking_malloc_total",
		"jcmd VM.native_memory metric Native Memory Tracking Malloc Total",
	}
	m[subsystem+"native_memory_tracking_overhead_bytes"] = metricAttr{
		"nm_overhead_kb",
		"native_memory_tracking_overhead_bytes",
		"jcmd VM.native_memory metric Native Memory Tracking Overhead Bytes",
	}

	// Shared Class Space section (sc_)
	m[subsystem+"shared_class_space_reserved_bytes"] = metricAttr{
		"sc_resv_kb",
		"shared_class_space_reserved_bytes",
		"jcmd VM.native_memory metric Shared Class Space Reserved Bytes",
	}
	m[subsystem+"shared_class_space_committed_bytes"] = metricAttr{
		"sc_comm_kb",
		"shared_class_space_committed_bytes",
		"jcmd VM.native_memory metric Shared Class Space Committed Bytes",
	}
	m[subsystem+"shared_class_space_mmap_reserved_bytes"] = metricAttr{
		"sc_mm_resv_kb",
		"shared_class_space_mmap_reserved_bytes",
		"jcmd VM.native_memory metric Shared Class Space Mmap Reserved Bytes",
	}
	m[subsystem+"shared_class_space_mmap_committed_bytes"] = metricAttr{
		"sc_mm_comm_kb",
		"shared_class_space_mmap_committed_bytes",
		"jcmd VM.native_memory metric Shared Class Space Mmap Committed Bytes",
	}

	// Arena Chunk section (ac_)
	m[subsystem+"arena_chunk_reserved_bytes"] = metricAttr{
		"ac_resv_kb",
		"arena_chunk_reserved_bytes",
		"jcmd VM.native_memory metric Arena Chunk Reserved Bytes",
	}
	m[subsystem+"arena_chunk_committed_bytes"] = metricAttr{
		"ac_comm_kb",
		"arena_chunk_committed_bytes",
		"jcmd VM.native_memory metric Arena Chunk Committed Bytes",
	}
	m[subsystem+"arena_chunk_malloc_bytes"] = metricAttr{
		"ac_malloc_kb",
		"arena_chunk_malloc_bytes",
		"jcmd VM.native_memory metric Arena Chunk Malloc Bytes",
	}

	// Logging section (lo_)
	m[subsystem+"logging_reserved_bytes"] = metricAttr{
		"lo_resv_kb",
		"logging_reserved_bytes",
		"jcmd VM.native_memory metric Logging Reserved Bytes",
	}
	m[subsystem+"logging_committed_bytes"] = metricAttr{
		"lo_comm_kb",
		"logging_committed_bytes",
		"jcmd VM.native_memory metric Logging Committed Bytes",
	}
	m[subsystem+"logging_malloc_bytes"] = metricAttr{
		"lo_malloc_kb",
		"logging_malloc_bytes",
		"jcmd VM.native_memory metric Logging Malloc Bytes",
	}
	m[subsystem+"logging_malloc_total"] = metricAttr{
		"lo_malloc_total",
		"logging_malloc_total",
		"jcmd VM.native_memory metric Logging Malloc Total",
	}

	// Arguments section (ar_)
	m[subsystem+"arguments_reserved_bytes"] = metricAttr{
		"ar_resv_kb",
		"arguments_reserved_bytes",
		"jcmd VM.native_memory metric Arguments Reserved Bytes",
	}
	m[subsystem+"arguments_committed_bytes"] = metricAttr{
		"ar_comm_kb",
		"arguments_committed_bytes",
		"jcmd VM.native_memory metric Arguments Committed Bytes",
	}
	m[subsystem+"arguments_malloc_bytes"] = metricAttr{
		"ar_malloc_kb",
		"arguments_malloc_bytes",
		"jcmd VM.native_memory metric Arguments Malloc Bytes",
	}
	m[subsystem+"arguments_malloc_total"] = metricAttr{
		"ar_malloc_total",
		"arguments_malloc_total",
		"jcmd VM.native_memory metric Arguments Malloc Total",
	}

	// Module section (mo_)
	m[subsystem+"module_reserved_bytes"] = metricAttr{
		"mo_resv_kb",
		"module_reserved_bytes",
		"jcmd VM.native_memory metric Module Reserved Bytes",
	}
	m[subsystem+"module_committed_bytes"] = metricAttr{
		"mo_comm_kb",
		"module_committed_bytes",
		"jcmd VM.native_memory metric Module Committed Bytes",
	}
	m[subsystem+"module_malloc_bytes"] = metricAttr{
		"mo_malloc_kb",
		"module_malloc_bytes",
		"jcmd VM.native_memory metric Module Malloc Bytes",
	}
	m[subsystem+"module_malloc_total"] = metricAttr{
		"mo_malloc_total",
		"module_malloc_total",
		"jcmd VM.native_memory metric Module Malloc Total",
	}

	// Safepoint section (sp_)
	m[subsystem+"safepoint_reserved_bytes"] = metricAttr{
		"sp_resv_kb",
		"safepoint_reserved_bytes",
		"jcmd VM.native_memory metric Safepoint Reserved Bytes",
	}
	m[subsystem+"safepoint_committed_bytes"] = metricAttr{
		"sp_comm_kb",
		"safepoint_committed_bytes",
		"jcmd VM.native_memory metric Safepoint Committed Bytes",
	}
	m[subsystem+"safepoint_mmap_reserved_bytes"] = metricAttr{
		"sp_mm_resv_kb",
		"safepoint_mmap_reserved_bytes",
		"jcmd VM.native_memory metric Safepoint Mmap Reserved Bytes",
	}
	m[subsystem+"safepoint_mmap_committed_bytes"] = metricAttr{
		"sp_mm_comm_kb",
		"safepoint_mmap_committed_bytes",
		"jcmd VM.native_memory metric Safepoint Mmap Committed Bytes",
	}

	// Synchronization section (sn_)
	m[subsystem+"synchronization_reserved_bytes"] = metricAttr{
		"sn_resv_kb",
		"synchronization_reserved_bytes",
		"jcmd VM.native_memory metric Synchronization Reserved Bytes",
	}
	m[subsystem+"synchronization_committed_bytes"] = metricAttr{
		"sn_comm_kb",
		"synchronization_committed_bytes",
		"jcmd VM.native_memory metric Synchronization Committed Bytes",
	}
	m[subsystem+"synchronization_malloc_bytes"] = metricAttr{
		"sn_malloc_kb",
		"synchronization_malloc_bytes",
		"jcmd VM.native_memory metric Synchronization Malloc Bytes",
	}
	m[subsystem+"synchronization_malloc_total"] = metricAttr{
		"sn_malloc_total",
		"synchronization_malloc_total",
		"jcmd VM.native_memory metric Synchronization Malloc Total",
	}

	return &m
}

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
