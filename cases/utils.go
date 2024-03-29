package cases

import (
	"runtime"
	"sync"
	"time"
)

type statistics struct {
	startMutex   sync.Mutex
	endMutex     sync.Mutex
	memoryStatus *runtime.MemStats
	startTime    time.Time
}

func (s *statistics) start() {
	s.startMutex.Lock()
	runtime.ReadMemStats(s.memoryStatus)
	s.startTime = time.Now()
}

func (s *statistics) end() (memoryAllocBytes uint64, elapsed time.Duration) {
	s.endMutex.Lock()
	s.startMutex.Unlock()
	var endMemoryStatus runtime.MemStats
	runtime.ReadMemStats(&endMemoryStatus)
	defer func() {
		s.startMutex = sync.Mutex{}
		s.endMutex = sync.Mutex{}
		s.memoryStatus = &runtime.MemStats{}
		s.startTime = time.Now()
	}()
	return endMemoryStatus.TotalAlloc - s.memoryStatus.TotalAlloc, time.Since(s.startTime)
}

func newStatistics() *statistics {
	return &statistics{
		startMutex:   sync.Mutex{},
		endMutex:     sync.Mutex{},
		memoryStatus: &runtime.MemStats{},
		startTime:    time.Time{},
	}
}
