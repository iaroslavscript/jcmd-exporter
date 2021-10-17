package main

import "regexp"

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
