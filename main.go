package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metricAttr struct {
	reGroup string
	name    string
	help    string
}
type metricLabelsMap map[string]metricAttr

var addr = flag.String("listen-address", ":2112", "The address to listen on for HTTP requests.")

func NewMetricLabelsMap() *metricLabelsMap {

	subsystem := "native_memory"

	m := make(metricLabelsMap)

	// Total Section (to_)
	m[subsystem+"_total_reserved_bytes"] = metricAttr{
		"to_resv_kb",
		"_total_reserved_bytes",
		"jcmd VM.native_memory metric Reserved Bytes",
	}
	m[subsystem+"_total_committed_bytes"] = metricAttr{
		"to_comm_kb",
		"_total_committed_bytes",
		"jcmd VM.native_memory metric Committed Bytes",
	}

	// Java Heap section (hp_)
	m[subsystem+"_java_heap_reserved_bytes"] = metricAttr{
		"hp_resv_kb",
		"_java_heap_reserved_bytes",
		"jcmd VM.native_memory metric Java Heap Reserved Bytes",
	}
	m[subsystem+"_java_heap_committed_bytes"] = metricAttr{
		"hp_comm_kb",
		"_java_heap_committed_bytes",
		"jcmd VM.native_memory metric Java Heap Committed Bytes",
	}

	m[subsystem+"_java_heap_mmap_reserved_bytes"] = metricAttr{
		"hp_mm_resv_kb",
		"_java_heap_mmap_reserved_bytes",
		"jcmd VM.native_memory metric Java Heap Mmap Reserved Bytes",
	}
	m[subsystem+"_java_heap_mmap_committed_bytes"] = metricAttr{
		"hp_mm_comm_kb",
		"_java_heap_mmap_committed_bytes",
		"jcmd VM.native_memory metric Java Heap Mmap Committed Bytes",
	}

	// Class section (cl_)
	m[subsystem+"_class_reserved_bytes"] = metricAttr{
		"cl_resv_kb",
		"_class_reserved_bytes",
		"jcmd VM.native_memory metric Class Reserved Bytes",
	}
	m[subsystem+"_class_committed_bytes"] = metricAttr{
		"cl_comm_kb",
		"_class_committed_bytes",
		"jcmd VM.native_memory metric Class Committed Bytes",
	}
	m[subsystem+"_class_classes_total"] = metricAttr{
		"cl_classes_total",
		"_class_classes_total",
		"jcmd VM.native_memory metric Class Classes Total",
	}
	m[subsystem+"_class_instance_classes_total"] = metricAttr{
		"cl_ic_total",
		"_class_instance_classes_total",
		"jcmd VM.native_memory metric Class Instance Classes Total",
	}
	m[subsystem+"_class_array_classes_total"] = metricAttr{
		"cl_ac_total",
		"_class_array_classes_total",
		"jcmd VM.native_memory metric Class Array Classes Total",
	}
	m[subsystem+"_class_malloc_bytes"] = metricAttr{
		"cl_malloc_kb",
		"_class_malloc_bytes",
		"jcmd VM.native_memory metric Class Malloc Bytes",
	}
	m[subsystem+"_class_malloc_total"] = metricAttr{
		"cl_malloc_total",
		"_class_malloc_total",
		"jcmd VM.native_memory metric Class Malloc Total",
	}
	m[subsystem+"_class_mmap_reserved_bytes"] = metricAttr{
		"cl_mm_resv_kb",
		"_class_mmap_reserved_bytes",
		"jcmd VM.native_memory metric Class Mmap Reserved Bytes",
	}
	m[subsystem+"_class_mmap_committed_bytes"] = metricAttr{
		"cl_mm_comm_kb",
		"_class_mmap_committed_bytes",
		"jcmd VM.native_memory metric Class Mmap Committed Bytes",
	}

	// Class Metadata subsection (cl_me_)
	m[subsystem+"_class_metadata_reserved_bytes"] = metricAttr{
		"cl_me_resv_kb",
		"_class_metadata_reserved_bytes",
		"jcmd VM.native_memory metric Class Metadata Reserved Bytes",
	}
	m[subsystem+"_class_metadata_committed_bytes"] = metricAttr{
		"cl_me_comm_kb",
		"_class_metadata_committed_bytes",
		"jcmd VM.native_memory metric Class Metadata Committed Bytes",
	}
	m[subsystem+"_class_metadata_used_bytes"] = metricAttr{
		"cl_me_used_kb",
		"_class_metadata_used_bytes",
		"jcmd VM.native_memory metric Class Metadata Used Bytes",
	}
	m[subsystem+"_class_metadata_free_bytes"] = metricAttr{
		"cl_me_free_kb",
		"_class_metadata_free_bytes",
		"jcmd VM.native_memory metric Class Metadata Free Bytes",
	}
	m[subsystem+"_class_metadata_waste_bytes"] = metricAttr{
		"cl_me_waste_kb",
		"_class_metadata_waste_bytes",
		"jcmd VM.native_memory metric Class Metadata Waste Bytes",
	}

	// Class Class Space subsection (cl_cs_)
	m[subsystem+"_class_class_space_reserved_bytes"] = metricAttr{
		"cl_cs_resv_kb",
		"_class_class_space_reserved_bytes",
		"jcmd VM.native_memory metric Class Class Space Reserved Bytes",
	}
	m[subsystem+"_class_class_space_committed_bytes"] = metricAttr{
		"cl_cs_comm_kb",
		"_class_class_space_committed_bytes",
		"jcmd VM.native_memory metric Class Class Space Committed Bytes",
	}
	m[subsystem+"_class_class_space_used_bytes"] = metricAttr{
		"cl_cs_used_kb",
		"_class_class_space_used_bytes",
		"jcmd VM.native_memory metric Class Class Space Used Bytes",
	}
	m[subsystem+"_class_class_space_free_bytes"] = metricAttr{
		"cl_cs_free_kb",
		"_class_class_space_free_bytes",
		"jcmd VM.native_memory metric Class Class Space Free Bytes",
	}
	m[subsystem+"_class_class_space_waste_bytes"] = metricAttr{
		"cl_cs_waste_kb",
		"_class_class_space_waste_bytes",
		"jcmd VM.native_memory metric Class Class Space Waste Bytes",
	}

	// Thread section (th_)
	m[subsystem+"_thread_reserved_bytes"] = metricAttr{
		"th_resv_kb",
		"_thread_reserved_bytes",
		"jcmd VM.native_memory metric Thread Reserved Bytes",
	}
	m[subsystem+"_thread_committed_bytes"] = metricAttr{
		"th_comm_kb",
		"_thread_committed_bytes",
		"jcmd VM.native_memory metric Thread Committed Bytes",
	}
	m[subsystem+"_thread_total"] = metricAttr{
		"th_total",
		"_thread_total",
		"jcmd VM.native_memory metric Thread Total",
	}
	m[subsystem+"_thread_stack_reserved_bytes"] = metricAttr{
		"th_s_resv_kb",
		"_thread_stack_reserved_bytes",
		"jcmd VM.native_memory metric Thread Stack Reserved Bytes",
	}
	m[subsystem+"_thread_stack_committed_bytes"] = metricAttr{
		"th_s_comm_kb",
		"_thread_stack_committed_bytes",
		"jcmd VM.native_memory metric Thread Stack Committed Bytes",
	}
	m[subsystem+"_thread_malloc_bytes"] = metricAttr{
		"th_malloc_kb",
		"_thread_malloc_bytes",
		"jcmd VM.native_memory metric Thread Malloc Bytes",
	}
	m[subsystem+"_thread_malloc_total"] = metricAttr{
		"th_malloc_total",
		"_thread_malloc_total",
		"jcmd VM.native_memory metric Thread Malloc Total",
	}
	m[subsystem+"_thread_arena_bytes"] = metricAttr{
		"th_arena_kb",
		"_thread_arena_bytes",
		"jcmd VM.native_memory metric Thread Arena Bytes",
	}
	m[subsystem+"_thread_arena_total"] = metricAttr{
		"th_arena_total",
		"_thread_arena_total",
		"jcmd VM.native_memory metric Thread Arena Total",
	}

	// Code section (co_)
	m[subsystem+"_code_reserved_bytes"] = metricAttr{
		"co_resv_kb",
		"_code_reserved_bytes",
		"jcmd VM.native_memory metric Code Reserved Bytes",
	}
	m[subsystem+"_code_committed_bytes"] = metricAttr{
		"co_comm_kb",
		"_code_committed_bytes",
		"jcmd VM.native_memory metric Code Committed Bytes",
	}
	m[subsystem+"_code_malloc_bytes"] = metricAttr{
		"co_malloc_kb",
		"_code_malloc_bytes",
		"jcmd VM.native_memory metric Code Malloc Bytes",
	}
	m[subsystem+"_code_malloc_total"] = metricAttr{
		"co_malloc_total",
		"_code_malloc_total",
		"jcmd VM.native_memory metric Code Malloc Total",
	}
	m[subsystem+"_code_mmap_reserved_bytes"] = metricAttr{
		"co_mm_resv_kb",
		"_code_mmap_reserved_bytes",
		"jcmd VM.native_memory metric Code Mmap Reserved Bytes",
	}
	m[subsystem+"_code_mmap_committed_bytes"] = metricAttr{
		"co_mm_comm_kb",
		"_code_mmap_committed_bytes",
		"jcmd VM.native_memory metric Code Mmap Committed Bytes",
	}

	// GC section (gc_)
	m[subsystem+"_gc_reserved_bytes"] = metricAttr{
		"gc_resv_kb",
		"_gc_reserved_bytes",
		"jcmd VM.native_memory metric GC Reserved Bytes",
	}
	m[subsystem+"_gc_committed_bytes"] = metricAttr{
		"gc_comm_kb",
		"_gc_committed_bytes",
		"jcmd VM.native_memory metric GC Committed Bytes",
	}
	m[subsystem+"_gc_malloc_bytes"] = metricAttr{
		"gc_malloc_kb",
		"_gc_malloc_bytes",
		"jcmd VM.native_memory metric GC Malloc Bytes",
	}
	m[subsystem+"_gc_malloc_total"] = metricAttr{
		"gc_malloc_total",
		"_gc_malloc_total",
		"jcmd VM.native_memory metric GC Malloc Total",
	}
	m[subsystem+"_gc_mmap_reserved_bytes"] = metricAttr{
		"gc_mm_resv_kb",
		"_gc_mmap_reserved_bytes",
		"jcmd VM.native_memory metric GC Mmap Reserved Bytes",
	}
	m[subsystem+"_gc_mmap_committed_bytes"] = metricAttr{
		"gc_mm_comm_kb",
		"_gc_mmap_committed_bytes",
		"jcmd VM.native_memory metric GC Mmap Committed Bytes",
	}

	// Compiler section (cp_)
	m[subsystem+"_compiler_reserved_bytes"] = metricAttr{
		"cp_resv_kb",
		"_compiler_reserved_bytes",
		"jcmd VM.native_memory metric Compiler Reserved Bytes",
	}
	m[subsystem+"_compiler_committed_bytes"] = metricAttr{
		"cp_comm_kb",
		"_compiler_committed_bytes",
		"jcmd VM.native_memory metric Compiler Committed Bytes",
	}
	m[subsystem+"_compiler_malloc_bytes"] = metricAttr{
		"cp_malloc_kb",
		"_compiler_malloc_bytes",
		"jcmd VM.native_memory metric Compiler Malloc Bytes",
	}
	m[subsystem+"_compiler_malloc_total"] = metricAttr{
		"cp_malloc_total",
		"_compiler_malloc_total",
		"jcmd VM.native_memory metric Compiler Malloc Total",
	}
	m[subsystem+"_compiler_arena_bytes"] = metricAttr{
		"cp_arena_kb",
		"_compiler_arena_bytes",
		"jcmd VM.native_memory metric Compiler Arena Bytes",
	}
	m[subsystem+"_compiler_arena_total"] = metricAttr{
		"cp_arena_total",
		"_compiler_arena_total",
		"jcmd VM.native_memory metric Compiler Arena Total",
	}

	// Internal section (in_)
	m[subsystem+"_internal_reserved_bytes"] = metricAttr{
		"in_resv_kb",
		"_internal_reserved_bytes",
		"jcmd VM.native_memory metric Internal Reserved Bytes",
	}
	m[subsystem+"_internal_committed_bytes"] = metricAttr{
		"in_comm_kb",
		"_internal_committed_bytes",
		"jcmd VM.native_memory metric Internal Committed Bytes",
	}
	m[subsystem+"_internal_malloc_bytes"] = metricAttr{
		"in_malloc_kb",
		"_internal_malloc_bytes",
		"jcmd VM.native_memory metric Internal Malloc Bytes",
	}
	m[subsystem+"_internal_malloc_total"] = metricAttr{
		"in_malloc_total",
		"_internal_malloc_total",
		"jcmd VM.native_memory metric Internal Malloc Total",
	}
	m[subsystem+"_internal_mmap_reserved_bytes"] = metricAttr{
		"in_mm_resv_kb",
		"_internal_mmap_reserved_bytes",
		"jcmd VM.native_memory metric Internal Mmap Reserved Bytes",
	}
	m[subsystem+"_internal_mmap_committed_bytes"] = metricAttr{
		"in_mm_comm_kb",
		"_internal_mmap_committed_bytes",
		"jcmd VM.native_memory metric Internal Mmap Committed Bytes",
	}

	// Symbol section (sy_)
	m[subsystem+"_symbol_reserved_bytes"] = metricAttr{
		"sy_resv_kb",
		"_symbol_reserved_bytes",
		"jcmd VM.native_memory metric Symbol Reserved Bytes",
	}
	m[subsystem+"_symbol_committed_bytes"] = metricAttr{
		"sy_comm_kb",
		"_symbol_committed_bytes",
		"jcmd VM.native_memory metric Symbol Committed Bytes",
	}
	m[subsystem+"_symbol_malloc_bytes"] = metricAttr{
		"sy_malloc_kb",
		"_symbol_malloc_bytes",
		"jcmd VM.native_memory metric Symbol Malloc Bytes",
	}
	m[subsystem+"_symbol_malloc_total"] = metricAttr{
		"sy_malloc_total",
		"_symbol_malloc_total",
		"jcmd VM.native_memory metric Symbol Malloc Total",
	}
	m[subsystem+"_symbol_arena_bytes"] = metricAttr{
		"sy_arena_kb",
		"_symbol_arena_bytes",
		"jcmd VM.native_memory metric Symbol Arena Bytes",
	}
	m[subsystem+"_symbol_arena_total"] = metricAttr{
		"sy_arena_total",
		"_symbol_arena_total",
		"jcmd VM.native_memory metric Symbol Arena Total",
	}

	// Native Memory Tracking section (nm_)
	m[subsystem+"_native_memory_tracking_reserved_bytes"] = metricAttr{
		"nm_resv_kb",
		"_native_memory_tracking_reserved_bytes",
		"jcmd VM.native_memory metric Native Memory Tracking Reserved Bytes",
	}
	m[subsystem+"_native_memory_tracking_committed_bytes"] = metricAttr{
		"nm_comm_kb",
		"_native_memory_tracking_committed_bytes",
		"jcmd VM.native_memory metric Native Memory Tracking Committed Bytes",
	}
	m[subsystem+"_native_memory_tracking_malloc_bytes"] = metricAttr{
		"nm_malloc_kb",
		"_native_memory_tracking_malloc_bytes",
		"jcmd VM.native_memory metric Native Memory Tracking Malloc Bytes",
	}
	m[subsystem+"_native_memory_tracking_malloc_total"] = metricAttr{
		"nm_malloc_total",
		"_native_memory_tracking_malloc_total",
		"jcmd VM.native_memory metric Native Memory Tracking Malloc Total",
	}
	m[subsystem+"_native_memory_tracking_overhead_bytes"] = metricAttr{
		"nm_overhead_kb",
		"_native_memory_tracking_overhead_bytes",
		"jcmd VM.native_memory metric Native Memory Tracking Overhead Bytes",
	}

	// Shared Class Space section (sc_)
	m[subsystem+"_shared_class_space_reserved_bytes"] = metricAttr{
		"sc_resv_kb",
		"_shared_class_space_reserved_bytes",
		"jcmd VM.native_memory metric Shared Class Space Reserved Bytes",
	}
	m[subsystem+"_shared_class_space_committed_bytes"] = metricAttr{
		"sc_comm_kb",
		"_shared_class_space_committed_bytes",
		"jcmd VM.native_memory metric Shared Class Space Committed Bytes",
	}
	m[subsystem+"_shared_class_space_mmap_reserved_bytes"] = metricAttr{
		"sc_mm_resv_kb",
		"_shared_class_space_mmap_reserved_bytes",
		"jcmd VM.native_memory metric Shared Class Space Mmap Reserved Bytes",
	}
	m[subsystem+"_shared_class_space_mmap_committed_bytes"] = metricAttr{
		"sc_mm_comm_kb",
		"_shared_class_space_mmap_committed_bytes",
		"jcmd VM.native_memory metric Shared Class Space Mmap Committed Bytes",
	}

	// Arena Chunk section (ac_)
	m[subsystem+"_arena_chunk_reserved_bytes"] = metricAttr{
		"ac_resv_kb",
		"_arena_chunk_reserved_bytes",
		"jcmd VM.native_memory metric Arena Chunk Reserved Bytes",
	}
	m[subsystem+"_arena_chunk_committed_bytes"] = metricAttr{
		"ac_comm_kb",
		"_arena_chunk_committed_bytes",
		"jcmd VM.native_memory metric Arena Chunk Committed Bytes",
	}
	m[subsystem+"_arena_chunk_malloc_bytes"] = metricAttr{
		"ac_malloc_kb",
		"_arena_chunk_malloc_bytes",
		"jcmd VM.native_memory metric Arena Chunk Malloc Bytes",
	}

	// Logging section (lo_)
	m[subsystem+"_logging_reserved_bytes"] = metricAttr{
		"lo_resv_kb",
		"_logging_reserved_bytes",
		"jcmd VM.native_memory metric Logging Reserved Bytes",
	}
	m[subsystem+"_logging_committed_bytes"] = metricAttr{
		"lo_comm_kb",
		"_logging_committed_bytes",
		"jcmd VM.native_memory metric Logging Committed Bytes",
	}
	m[subsystem+"_logging_malloc_bytes"] = metricAttr{
		"lo_malloc_kb",
		"_logging_malloc_bytes",
		"jcmd VM.native_memory metric Logging Malloc Bytes",
	}
	m[subsystem+"_logging_malloc_total"] = metricAttr{
		"lo_malloc_total",
		"_logging_malloc_total",
		"jcmd VM.native_memory metric Logging Malloc Total",
	}

	// Arguments section (ar_)
	m[subsystem+"_arguments_reserved_bytes"] = metricAttr{
		"ar_resv_kb",
		"_arguments_reserved_bytes",
		"jcmd VM.native_memory metric Arguments Reserved Bytes",
	}
	m[subsystem+"_arguments_committed_bytes"] = metricAttr{
		"ar_comm_kb",
		"_arguments_committed_bytes",
		"jcmd VM.native_memory metric Arguments Committed Bytes",
	}
	m[subsystem+"_arguments_malloc_bytes"] = metricAttr{
		"ar_malloc_kb",
		"_arguments_malloc_bytes",
		"jcmd VM.native_memory metric Arguments Malloc Bytes",
	}
	m[subsystem+"_arguments_malloc_total"] = metricAttr{
		"ar_malloc_total",
		"_arguments_malloc_total",
		"jcmd VM.native_memory metric Arguments Malloc Total",
	}

	// Module section (mo_)
	m[subsystem+"_module_reserved_bytes"] = metricAttr{
		"mo_resv_kb",
		"_module_reserved_bytes",
		"jcmd VM.native_memory metric Module Reserved Bytes",
	}
	m[subsystem+"_module_committed_bytes"] = metricAttr{
		"mo_comm_kb",
		"_module_committed_bytes",
		"jcmd VM.native_memory metric Module Committed Bytes",
	}
	m[subsystem+"_module_malloc_bytes"] = metricAttr{
		"mo_malloc_kb",
		"_module_malloc_bytes",
		"jcmd VM.native_memory metric Module Malloc Bytes",
	}
	m[subsystem+"_module_malloc_total"] = metricAttr{
		"mo_malloc_total",
		"_module_malloc_total",
		"jcmd VM.native_memory metric Module Malloc Total",
	}

	// Safepoint section (sp_)
	m[subsystem+"_safepoint_reserved_bytes"] = metricAttr{
		"sp_resv_kb",
		"_safepoint_reserved_bytes",
		"jcmd VM.native_memory metric Safepoint Reserved Bytes",
	}
	m[subsystem+"_safepoint_committed_bytes"] = metricAttr{
		"sp_comm_kb",
		"_safepoint_committed_bytes",
		"jcmd VM.native_memory metric Safepoint Committed Bytes",
	}
	m[subsystem+"_safepoint_mmap_reserved_bytes"] = metricAttr{
		"sp_mm_resv_kb",
		"_safepoint_mmap_reserved_bytes",
		"jcmd VM.native_memory metric Safepoint Mmap Reserved Bytes",
	}
	m[subsystem+"_safepoint_mmap_committed_bytes"] = metricAttr{
		"sp_mm_comm_kb",
		"_safepoint_mmap_committed_bytes",
		"jcmd VM.native_memory metric Safepoint Mmap Committed Bytes",
	}

	// Synchronization section (sn_)
	m[subsystem+"_synchronization_reserved_bytes"] = metricAttr{
		"sn_resv_kb",
		"_synchronization_reserved_bytes",
		"jcmd VM.native_memory metric Synchronization Reserved Bytes",
	}
	m[subsystem+"_synchronization_committed_bytes"] = metricAttr{
		"sn_comm_kb",
		"_synchronization_committed_bytes",
		"jcmd VM.native_memory metric Synchronization Committed Bytes",
	}
	m[subsystem+"_synchronization_malloc_bytes"] = metricAttr{
		"sn_malloc_kb",
		"_synchronization_malloc_bytes",
		"jcmd VM.native_memory metric Synchronization Malloc Bytes",
	}
	m[subsystem+"_synchronization_malloc_total"] = metricAttr{
		"sn_malloc_total",
		"_synchronization_malloc_total",
		"jcmd VM.native_memory metric Synchronization Malloc Total",
	}

	return &m
}

/*
func NewMetricsNativeMemory() *MetricsNativeMemory {

	metricsNamespace := "jcmd"
	metricsSubsystem := "native_memory"

	type metricsMap struct {
		opsReservedBytes  prometheus.Gauge
	}
	m := MetricsNativeMemory{
		opsReservedBytes: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "reserved_bytes",
			Help:      "",
		}),
		opsCommittedBytes: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "commited_bytes",
			Help:      "",
		}),
	}

	m.opsReservedBytes.Set(0.0)
	m.opsCommittedBytes.Set(0.0)

	return &m
}
*/

func call_jcmd() string {
	app := "jcmd"

	arg0 := "SingleThread"
	arg1 := "VM.native_memory"

	cmd := exec.Command(app, arg0, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(stdout)
}

/*func parse_response(s string, m *MetricsNativeMemory) {

	pattern := ``

	re_total := regexp.MustCompile(pattern)

	fields := make(map[string]struct{string, string, string,})

	if matches := re_total.FindStringSubmatch(s); matches != nil {

		var v string
		var f float64
		var err error

		matches := re_total.FindStringSubmatch(s)

		for group_name, field := range fields {

			v = matches[re_total.SubexpIndex(group_name)]

			if f, err = strconv.ParseFloat(v, 64); err == nil {
				(*field).Set(f * 1024)
			}

		}

	} else {
		fmt.Println("\tnot matched TOTAL")
	}
}*/

func NewPattern() *regexp.Regexp {

	pattern := `(?ms)` +

		// Total section (to_)
		`^Total: reserved=(?P<to_resv_kb>\d+)KB, committed=(?P<to_comm_kb>\d+)KB.+` +

		// Java Heap section (hp_)
		`-\s+Java Heap \(reserved=(?P<hp_resv_kb>\d+)KB, committed=(?P<hp_comm_kb>\d+)KB\).+` +
		`\(mmap: reserved=(?P<hp_mm_resv_kb>\d+)KB, committed=(?P<hp_mm_comm_kb>\d+)KB\).+` +

		// Class section (cl_)
		`-\s+Class \(reserved=(?P<cl_resv_kb>\d+)KB, committed=(?P<cl_comm_kb>\d+)KB\).+` +
		`\(classes #(?P<cl_classes_total>\d+)\).+` +
		`\(\s+instance classes #(?P<cl_ic_total>\d+), array classes #(?P<cl_ac_total>\d+)\).+` +
		`\(malloc=(?P<cl_malloc_kb>\d+)KB #(?P<cl_malloc_total>\d+)\).+` +
		`\(mmap: reserved=(?P<cl_mm_resv_kb>\d+)KB, committed=(?P<cl_mm_comm_kb>\d+)KB\).+` +

		// Class Metadata subsection (cl_me_)
		`\(\s+reserved=(?P<cl_me_resv_kb>\d+)KB, committed=(?P<cl_me_comm_kb>\d+)KB\).+` +
		`\(\s+used=(?P<cl_me_used_kb>\d+)KB\).+` +
		`\(\s+free=(?P<cl_me_free_kb>\d+)KB\).+` +
		`\(\s+waste=(?P<cl_me_waste_kb>\d+)KB =.+%\).+` +

		// Class Class Space subsection (cl_cs_)
		`\(\s+reserved=(?P<cl_cs_resv_kb>\d+)KB, committed=(?P<cl_cs_comm_kb>\d+)KB\).+` +
		`\(\s+used=(?P<cl_cs_used_kb>\d+)KB\).+` +
		`\(\s+free=(?P<cl_cs_free_kb>\d+)KB\).+` +
		`\(\s+waste=(?P<cl_cs_waste_kb>\d+)KB =.+%\).+` +

		// Thread section (th_)
		`-\s+Thread\s+\(reserved=(?P<th_resv_kb>\d+)KB, committed=(?P<th_comm_kb>\d+)KB\).+` +
		`\(thread #(?P<th_total>\d+)\).+` +
		`\(stack: reserved=(?P<th_s_resv_kb>\d+)KB, committed=(?P<th_s_comm_kb>\d+)KB\).+` +
		`\(malloc=(?P<th_malloc_kb>\d+)KB #(?P<th_malloc_total>\d+)\).+` +
		`\(arena=(?P<th_arena_kb>\d+)KB #(?P<th_arena_total>\d+)\).+` +

		// Code section (co_)
		`-\s+Code \(reserved=(?P<co_resv_kb>\d+)KB, committed=(?P<co_comm_kb>\d+)KB\).+` +
		`\s+\(malloc=(?P<co_malloc_kb>\d+)KB #(?P<co_malloc_total>\d+)\).+` +
		`\s+\(mmap: reserved=(?P<co_mm_resv_kb>\d+)KB, committed=(?P<co_mm_comm_kb>\d+)KB\).+` +

		// GC section (gc_)
		`-\s+GC \(reserved=(?P<gc_resv_kb>\d+)KB, committed=(?P<gc_comm_kb>\d+)KB\).+` +
		`\s+\(malloc=(?P<gc_malloc_kb>\d+)KB #(?P<gc_malloc_total>\d+)\).+` +
		`\s+\(mmap: reserved=(?P<gc_mm_resv_kb>\d+)KB, committed=(?P<gc_mm_comm_kb>\d+)KB\).+` +

		// Compiler section (cp_)
		`-\s+Compiler \(reserved=(?P<cp_resv_kb>\d+)KB, committed=(?P<cp_comm_kb>\d+)KB\).+` +
		`\s+\(malloc=(?P<cp_malloc_kb>\d+)KB #(?P<cp_malloc_total>\d+)\).+` +
		`\s+\(arena=(?P<cp_arena_kb>\d+)KB #(?P<cp_arena_total>\d+)\).+` +

		// Internal section (in_)
		`-\s+Internal \(reserved=(?P<in_resv_kb>\d+)KB, committed=(?P<in_comm_kb>\d+)KB\).+` +
		`\s+\(malloc=(?P<in_malloc_kb>\d+)KB #(?P<in_malloc_total>\d+)\).+` +
		`\s+\(mmap: reserved=(?P<in_mm_resv_kb>\d+)KB, committed=(?P<in_mm_comm_kb>\d+)KB\).+` +

		// Symbol section (sy_)
		`-\s+Symbol \(reserved=(?P<sy_resv_kb>\d+)KB, committed=(?P<sy_comm_kb>\d+)KB\).+` +
		`\s+\(malloc=(?P<sy_malloc_kb>\d+)KB #(?P<sy_malloc_total>\d+)\).+` +
		`\s+\(arena=(?P<sy_arena_kb>\d+)KB #(?P<sy_arena_total>\d+)\).+` +

		// Native Memory Tracking section (nm_)
		`-\s+Native Memory Tracking \(reserved=(?P<nm_resv_kb>\d+)KB, committed=(?P<nm_comm_kb>\d+)KB\).+` +
		`\s+\(malloc=(?P<nm_malloc_kb>\d+)KB #(?P<nm_malloc_total>\d+)\).+` +
		`\s+\(tracking overhead=(?P<nm_overhead_kb>\d+)KB\).+` +

		// Shared Class Space section (sc_)
		`-\s+Shared class space \(reserved=(?P<sc_resv_kb>\d+)KB, committed=(?P<sc_comm_kb>\d+)KB\).+` +
		`\s+\(mmap: reserved=(?P<sc_mm_resv_kb>\d+)KB, committed=(?P<sc_mm_comm_kb>\d+)KB\).+` +

		// Arena Chunk section (ac_)
		`-\s+Arena Chunk \(reserved=(?P<ac_resv_kb>\d+)KB, committed=(?P<ac_comm_kb>\d+)KB\).+` +
		`\s+\(malloc=(?P<ac_malloc_kb>\d+)KB\).+` +

		// Logging section (lo_)
		`-\s+Logging \(reserved=(?P<lo_resv_kb>\d+)KB, committed=(?P<lo_comm_kb>\d+)KB\).+` +
		`\s+\(malloc=(?P<lo_malloc_kb>\d+)KB #(?P<lo_malloc_total>\d+)\).+` +

		// Arguments section (ar_)
		`-\s+Arguments \(reserved=(?P<ar_resv_kb>\d+)KB, committed=(?P<ar_comm_kb>\d+)KB\).+` +
		`\s+\(malloc=(?P<ar_malloc_kb>\d+)KB #(?P<ar_malloc_total>\d+)\).+` +

		// Module section (mo_)
		`-\s+Module \(reserved=(?P<mo_resv_kb>\d+)KB, committed=(?P<mo_comm_kb>\d+)KB\).+` +
		`\s+\(malloc=(?P<mo_malloc_kb>\d+)KB #(?P<mo_malloc_total>\d+)\).+` +

		// Safepoint section (sp_)
		`-\s+Safepoint \(reserved=(?P<sp_resv_kb>\d+)KB, committed=(?P<sp_comm_kb>\d+)KB\).+` +
		`\s+\(mmap: reserved=(?P<sp_mm_resv_kb>\d+)KB, committed=(?P<sp_mm_comm_kb>\d+)KB\).+` +

		// Synchronization section (sn_)
		`-\s+Synchronization \(reserved=(?P<sn_resv_kb>\d+)KB, committed=(?P<sn_comm_kb>\d+)KB\).+` +
		`\s+\(malloc=(?P<sn_malloc_kb>\d+)KB #(?P<sn_malloc_total>\d+)\)`

	return regexp.MustCompile(pattern)
}

func main1() {

	nativeMemory := NewMetricsNativeMemory()

	stdout := call_jcmd()

	parse_response(stdout, nativeMemory)

	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func main() {

	dat, err := os.ReadFile("./report.txt")
	if err != nil {
		panic(err)
	}
	s := string(dat)

	re_total := NewPattern()

	if matches := re_total.FindStringSubmatch(s); matches != nil {
		fmt.Println("Matched!")
	}

}
