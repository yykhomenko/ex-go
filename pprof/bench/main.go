package main

import (
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	runtime.GC()

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	log.Println(mem.Alloc)
	log.Println(mem.TotalAlloc)
	log.Println(mem.HeapAlloc)
	log.Println(mem.HeapSys)
}
