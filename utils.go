package main

import (
	"encoding/json"
	"regexp"
)

const DEFAULT_METRICS_JSON string = `{
	"native_memory": [
		{
			"regex_group": "to_resv_kb",
			"name": "total_reserved_bytes",
			"help": "jcmd VM.native_memory section Total metric Reserved Bytes"
		},
		{
			"regex_group": "to_comm_kb",
			"name": "total_committed_bytes",
			"help": "jcmd VM.native_memory section Total metric Committed Bytes"
		},

		{
			"regex_group": "hp_resv_kb",
			"name": "java_heap_reserved_bytes",
			"help": "jcmd VM.native_memory section Java Heap metric Reserved Bytes"
		},
		{
			"regex_group": "hp_comm_kb",
			"name": "java_heap_committed_bytes",
			"help": "jcmd VM.native_memory section Java Heap metric Committed Bytes"
		},

		{
			"regex_group": "hp_mm_resv_kb",
			"name": "java_heap_mmap_reserved_bytes",
			"help": "jcmd VM.native_memory section Java Heap metric Mmap Reserved Bytes"
		},
		{
			"regex_group": "hp_mm_comm_kb",
			"name": "java_heap_mmap_committed_bytes",
			"help": "jcmd VM.native_memory section Java Heap metric Mmap Committed Bytes"
		},

		{
			"regex_group": "cl_resv_kb",
			"name": "class_reserved_bytes",
			"help": "jcmd VM.native_memory section Class metric Reserved Bytes"
		},
		{
			"regex_group": "cl_comm_kb",
			"name": "class_committed_bytes",
			"help": "jcmd VM.native_memory section Class metric Committed Bytes"
		},
		{
			"regex_group": "cl_classes_total",
			"name": "class_classes_total",
			"help": "jcmd VM.native_memory section Class metric Classes Total"
		},
		{
			"regex_group": "cl_ic_total",
			"name": "class_instance_classes_total",
			"help": "jcmd VM.native_memory section Class metric Instance Classes Total"
		},
		{
			"regex_group": "cl_ac_total",
			"name": "class_array_classes_total",
			"help": "jcmd VM.native_memory section Class metric Array Classes Total"
		},
		{
			"regex_group": "cl_malloc_kb",
			"name": "class_malloc_bytes",
			"help": "jcmd VM.native_memory section Class metric Malloc Bytes"
		},
		{
			"regex_group": "cl_malloc_total",
			"name": "class_malloc_total",
			"help": "jcmd VM.native_memory section Class metric Malloc Total"
		},
		{
			"regex_group": "cl_mm_resv_kb",
			"name": "class_mmap_reserved_bytes",
			"help": "jcmd VM.native_memory section Class metric Mmap Reserved Bytes"
		},
		{
			"regex_group": "cl_mm_comm_kb",
			"name": "class_mmap_committed_bytes",
			"help": "jcmd VM.native_memory section Class metric Mmap Committed Bytes"
		},

		{
			"regex_group": "cl_me_resv_kb",
			"name": "class_metadata_reserved_bytes",
			"help": "jcmd VM.native_memory section Class subsection Metadata metric Reserved Bytes"
		},
		{
			"regex_group": "cl_me_comm_kb",
			"name": "class_metadata_committed_bytes",
			"help": "jcmd VM.native_memory section Class subsection Metadata metric Committed Bytes"
		},
		{
			"regex_group": "cl_me_used_kb",
			"name": "class_metadata_used_bytes",
			"help": "jcmd VM.native_memory section Class subsection Metadata metric Used Bytes"
		},
		{
			"regex_group": "cl_me_free_kb",
			"name": "class_metadata_free_bytes",
			"help": "jcmd VM.native_memory section Class subsection Metadata metric Free Bytes"
		},
		{
			"regex_group": "cl_me_waste_kb",
			"name": "class_metadata_waste_bytes",
			"help": "jcmd VM.native_memory section Class subsection Metadata metric Waste Bytes"
		},

		{
			"regex_group": "cl_cs_resv_kb",
			"name": "class_class_space_reserved_bytes",
			"help": "jcmd VM.native_memory section Class subsection Class Space metric Reserved Bytes"
		},
		{
			"regex_group": "cl_cs_comm_kb",
			"name": "class_class_space_committed_bytes",
			"help": "jcmd VM.native_memory section Class subsection Class Space metric Committed Bytes"
		},
		{
			"regex_group": "cl_cs_used_kb",
			"name": "class_class_space_used_bytes",
			"help": "jcmd VM.native_memory section Class subsection Class Space metric Used Bytes"
		},
		{
			"regex_group": "cl_cs_free_kb",
			"name": "class_class_space_free_bytes",
			"help": "jcmd VM.native_memory section Class subsection Class Space metric Free Bytes"
		},
		{
			"regex_group": "cl_cs_waste_kb",
			"name": "class_class_space_waste_bytes",
			"help": "jcmd VM.native_memory section Class subsection Class Space metric Waste Bytes"
		},

		{
			"regex_group": "th_resv_kb",
			"name": "thread_reserved_bytes",
			"help": "jcmd VM.native_memory section Thread metric Reserved Bytes"
		},
		{
			"regex_group": "th_comm_kb",
			"name": "thread_committed_bytes",
			"help": "jcmd VM.native_memory section Thread metric Committed Bytes"
		},
		{
			"regex_group": "th_total",
			"name": "thread_total",
			"help": "jcmd VM.native_memory section Thread metric Total"
		},
		{
			"regex_group": "th_s_resv_kb",
			"name": "thread_stack_reserved_bytes",
			"help": "jcmd VM.native_memory section Thread metric Stack Reserved Bytes"
		},
		{
			"regex_group": "th_s_comm_kb",
			"name": "thread_stack_committed_bytes",
			"help": "jcmd VM.native_memory section Thread metric Stack Committed Bytes"
		},
		{
			"regex_group": "th_malloc_kb",
			"name": "thread_malloc_bytes",
			"help": "jcmd VM.native_memory section Thread metric Malloc Bytes"
		},
		{
			"regex_group": "th_malloc_total",
			"name": "thread_malloc_total",
			"help": "jcmd VM.native_memory section Thread metric Malloc Total"
		},
		{
			"regex_group": "th_arena_kb",
			"name": "thread_arena_bytes",
			"help": "jcmd VM.native_memory section Thread metric Arena Bytes"
		},
		{
			"regex_group": "th_arena_total",
			"name": "thread_arena_total",
			"help": "jcmd VM.native_memory section Thread metric Arena Total"
		},

		{
			"regex_group": "co_resv_kb",
			"name": "code_reserved_bytes",
			"help": "jcmd VM.native_memory section Code metric Reserved Bytes"
		},
		{
			"regex_group": "co_comm_kb",
			"name": "code_committed_bytes",
			"help": "jcmd VM.native_memory section Code metric Committed Bytes"
		},
		{
			"regex_group": "co_malloc_kb",
			"name": "code_malloc_bytes",
			"help": "jcmd VM.native_memory section Code metric Malloc Bytes"
		},
		{
			"regex_group": "co_malloc_total",
			"name": "code_malloc_total",
			"help": "jcmd VM.native_memory section Code metric Malloc Total"
		},
		{
			"regex_group": "co_mm_resv_kb",
			"name": "code_mmap_reserved_bytes",
			"help": "jcmd VM.native_memory section Code metric Mmap Reserved Bytes"
		},
		{
			"regex_group": "co_mm_comm_kb",
			"name": "code_mmap_committed_bytes",
			"help": "jcmd VM.native_memory section Code metric Mmap Committed Bytes"
		},

		{
			"regex_group": "gc_resv_kb",
			"name": "gc_reserved_bytes",
			"help": "jcmd VM.native_memory section GC metric Reserved Bytes"
		},
		{
			"regex_group": "gc_comm_kb",
			"name": "gc_committed_bytes",
			"help": "jcmd VM.native_memory section GC metric Committed Bytes"
		},
		{
			"regex_group": "gc_malloc_kb",
			"name": "gc_malloc_bytes",
			"help": "jcmd VM.native_memory section GC metric Malloc Bytes"
		},
		{
			"regex_group": "gc_malloc_total",
			"name": "gc_malloc_total",
			"help": "jcmd VM.native_memory section GC metric Malloc Total"
		},
		{
			"regex_group": "gc_mm_resv_kb",
			"name": "gc_mmap_reserved_bytes",
			"help": "jcmd VM.native_memory section GC metric Mmap Reserved Bytes"
		},
		{
			"regex_group": "gc_mm_comm_kb",
			"name": "gc_mmap_committed_bytes",
			"help": "jcmd VM.native_memory section GC metric Mmap Committed Bytes"
		},

		{
			"regex_group": "cp_resv_kb",
			"name": "compiler_reserved_bytes",
			"help": "jcmd VM.native_memory section Compiler metric Reserved Bytes"
		},
		{
			"regex_group": "cp_comm_kb",
			"name": "compiler_committed_bytes",
			"help": "jcmd VM.native_memory section Compiler metric Committed Bytes"
		},
		{
			"regex_group": "cp_malloc_kb",
			"name": "compiler_malloc_bytes",
			"help": "jcmd VM.native_memory section Compiler metric Malloc Bytes"
		},
		{
			"regex_group": "cp_malloc_total",
			"name": "compiler_malloc_total",
			"help": "jcmd VM.native_memory section Compiler metric Malloc Total"
		},
		{
			"regex_group": "cp_arena_kb",
			"name": "compiler_arena_bytes",
			"help": "jcmd VM.native_memory section Compiler metric Arena Bytes"
		},
		{
			"regex_group": "cp_arena_total",
			"name": "compiler_arena_total",
			"help": "jcmd VM.native_memory section Compiler metric Arena Total"
		},

		{
			"regex_group": "in_resv_kb",
			"name": "internal_reserved_bytes",
			"help": "jcmd VM.native_memory section Internal metric Reserved Bytes"
		},
		{
			"regex_group": "in_comm_kb",
			"name": "internal_committed_bytes",
			"help": "jcmd VM.native_memory section Internal metric Committed Bytes"
		},
		{
			"regex_group": "in_malloc_kb",
			"name": "internal_malloc_bytes",
			"help": "jcmd VM.native_memory section Internal metric Malloc Bytes"
		},
		{
			"regex_group": "in_malloc_total",
			"name": "internal_malloc_total",
			"help": "jcmd VM.native_memory section Internal metric Malloc Total"
		},
		{
			"regex_group": "in_mm_resv_kb",
			"name": "internal_mmap_reserved_bytes",
			"help": "jcmd VM.native_memory section Internal metric Mmap Reserved Bytes"
		},
		{
			"regex_group": "in_mm_comm_kb",
			"name": "internal_mmap_committed_bytes",
			"help": "jcmd VM.native_memory section Internal metric Mmap Committed Bytes"
		},

		{
			"regex_group": "sy_resv_kb",
			"name": "symbol_reserved_bytes",
			"help": "jcmd VM.native_memory section Symbol metric Reserved Bytes"
		},
		{
			"regex_group": "sy_comm_kb",
			"name": "symbol_committed_bytes",
			"help": "jcmd VM.native_memory section Symbol metric Committed Bytes"
		},
		{
			"regex_group": "sy_malloc_kb",
			"name": "symbol_malloc_bytes",
			"help": "jcmd VM.native_memory section Symbol metric Malloc Bytes"
		},
		{
			"regex_group": "sy_malloc_total",
			"name": "symbol_malloc_total",
			"help": "jcmd VM.native_memory section Symbol metric Malloc Total"
		},
		{
			"regex_group": "sy_arena_kb",
			"name": "symbol_arena_bytes",
			"help": "jcmd VM.native_memory section Symbol metric Arena Bytes"
		},
		{
			"regex_group": "sy_arena_total",
			"name": "symbol_arena_total",
			"help": "jcmd VM.native_memory section Symbol metric Arena Total"
		},

		{
			"regex_group": "nm_resv_kb",
			"name": "native_memory_tracking_reserved_bytes",
			"help": "jcmd VM.native_memory section Native Memory Tracking metric Reserved Bytes"
		},
		{
			"regex_group": "nm_comm_kb",
			"name": "native_memory_tracking_committed_bytes",
			"help": "jcmd VM.native_memory section Native Memory Tracking metric Committed Bytes"
		},
		{
			"regex_group": "nm_malloc_kb",
			"name": "native_memory_tracking_malloc_bytes",
			"help": "jcmd VM.native_memory section Native Memory Tracking metric Malloc Bytes"
		},
		{
			"regex_group": "nm_malloc_total",
			"name": "native_memory_tracking_malloc_total",
			"help": "jcmd VM.native_memory section Native Memory Tracking metric Malloc Total"
		},
		{
			"regex_group": "nm_overhead_kb",
			"name": "native_memory_tracking_overhead_bytes",
			"help": "jcmd VM.native_memory section Native Memory Tracking metric Overhead Bytes"
		},

		{
			"regex_group": "sc_resv_kb",
			"name": "shared_class_space_reserved_bytes",
			"help": "jcmd VM.native_memory section Shared Class Space metricReserved Bytes"
		},
		{
			"regex_group": "sc_comm_kb",
			"name": "shared_class_space_committed_bytes",
			"help": "jcmd VM.native_memory section Shared Class Space metricCommitted Bytes"
		},
		{
			"regex_group": "sc_mm_resv_kb",
			"name": "shared_class_space_mmap_reserved_bytes",
			"help": "jcmd VM.native_memory section Shared Class Space metricMmap Reserved Bytes"
		},
		{
			"regex_group": "sc_mm_comm_kb",
			"name": "shared_class_space_mmap_committed_bytes",
			"help": "jcmd VM.native_memory section Shared Class Space metricMmap Committed Bytes"
		},

		{
			"regex_group": "ac_resv_kb",
			"name": "arena_chunk_reserved_bytes",
			"help": "jcmd VM.native_memory section Arena Chunk metricReserved Bytes"
		},
		{
			"regex_group": "ac_comm_kb",
			"name": "arena_chunk_committed_bytes",
			"help": "jcmd VM.native_memory section Arena Chunk metricCommitted Bytes"
		},
		{
			"regex_group": "ac_malloc_kb",
			"name": "arena_chunk_malloc_bytes",
			"help": "jcmd VM.native_memory section Arena Chunk metricMalloc Bytes"
		},

		{
			"regex_group": "lo_resv_kb",
			"name": "logging_reserved_bytes",
			"help": "jcmd VM.native_memory section Logging metric Reserved Bytes"
		},
		{
			"regex_group": "lo_comm_kb",
			"name": "logging_committed_bytes",
			"help": "jcmd VM.native_memory section Logging metric Committed Bytes"
		},
		{
			"regex_group": "lo_malloc_kb",
			"name": "logging_malloc_bytes",
			"help": "jcmd VM.native_memory section Logging metric Malloc Bytes"
		},
		{
			"regex_group": "lo_malloc_total",
			"name": "logging_malloc_total",
			"help": "jcmd VM.native_memory section Logging metric Malloc Total"
		},

		{
			"regex_group": "ar_resv_kb",
			"name": "arguments_reserved_bytes",
			"help": "jcmd VM.native_memory section Arguments metric Reserved Bytes"
		},
		{
			"regex_group": "ar_comm_kb",
			"name": "arguments_committed_bytes",
			"help": "jcmd VM.native_memory section Arguments metric Committed Bytes"
		},
		{
			"regex_group": "ar_malloc_kb",
			"name": "arguments_malloc_bytes",
			"help": "jcmd VM.native_memory section Arguments metric Malloc Bytes"
		},
		{
			"regex_group": "ar_malloc_total",
			"name": "arguments_malloc_total",
			"help": "jcmd VM.native_memory section Arguments metric Malloc Total"
		},

		{
			"regex_group": "mo_resv_kb",
			"name": "module_reserved_bytes",
			"help": "jcmd VM.native_memory section Module metric Reserved Bytes"
		},
		{
			"regex_group": "mo_comm_kb",
			"name": "module_committed_bytes",
			"help": "jcmd VM.native_memory section Module metric Committed Bytes"
		},
		{
			"regex_group": "mo_malloc_kb",
			"name": "module_malloc_bytes",
			"help": "jcmd VM.native_memory section Module metric Malloc Bytes"
		},
		{
			"regex_group": "mo_malloc_total",
			"name": "module_malloc_total",
			"help": "jcmd VM.native_memory section Module metric Malloc Total"
		},

		{
			"regex_group": "sp_resv_kb",
			"name": "safepoint_reserved_bytes",
			"help": "jcmd VM.native_memory section Safepoint metricReserved Bytes"
		},
		{
			"regex_group": "sp_comm_kb",
			"name": "safepoint_committed_bytes",
			"help": "jcmd VM.native_memory section Safepoint metricCommitted Bytes"
		},
		{
			"regex_group": "sp_mm_resv_kb",
			"name": "safepoint_mmap_reserved_bytes",
			"help": "jcmd VM.native_memory section Safepoint metricMmap Reserved Bytes"
		},
		{
			"regex_group": "sp_mm_comm_kb",
			"name": "safepoint_mmap_committed_bytes",
			"help": "jcmd VM.native_memory section Safepoint metricMmap Committed Bytes"
		},

		{
			"regex_group": "sn_resv_kb",
			"name": "synchronization_reserved_bytes",
			"help": "jcmd VM.native_memory section Synchronization metric Reserved Bytes"
		},
		{
			"regex_group": "sn_comm_kb",
			"name": "synchronization_committed_bytes",
			"help": "jcmd VM.native_memory section Synchronization metric Committed Bytes"
		},
		{
			"regex_group": "sn_malloc_kb",
			"name": "synchronization_malloc_bytes",
			"help": "jcmd VM.native_memory section Synchronization metric Malloc Bytes"
		},
		{
			"regex_group": "sn_malloc_total",
			"name": "synchronization_malloc_total",
			"help": "jcmd VM.native_memory section Synchronization metric Malloc Total"
		}
	]
}`

func NewMetricLabelsMap2(data []byte) *metricLabelsMap {

	type metricsAll struct {
		attr []metricAttr `json:"native_memory"`
	}

	var metrics metricsAll

	if err := json.Unmarshal(data, &metrics); err != nil {
		panic(err)
	}

	subsystem := "native_memory"

	m := make(metricLabelsMap, len(metrics.attr)) // TODO capacity ???

	for _, value := range metrics.attr {

		m[subsystem+value.name] = metricAttr{
			value.reGroup,
			value.name,
			value.help,
		}
	}

	return &m
}

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
