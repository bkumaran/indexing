package system

//#cgo LDFLAGS: -lsigar
//#include <sigar.h>
import "C"

import (
	"errors"
	"fmt"
)

type SystemStats struct {
	handle *C.sigar_t
	pid    C.sigar_pid_t
}

//
// Open a new handle
//
func NewSystemStats() (*SystemStats, error) {

	var handle *C.sigar_t

	if err := C.sigar_open(&handle); err != C.SIGAR_OK {
		return nil, errors.New(fmt.Sprintf("Fail to open sigar.  Error code = %v", err))
	}

	h := &SystemStats{}
	h.handle = handle
	h.pid = C.sigar_pid_get(handle)

	return h, nil
}

//
// Close handle
//
func (h *SystemStats) Close() {
	C.sigar_close(h.handle)
}

//
// Get CPU percentage
//
func (h *SystemStats) ProcessCpuPercent() (float64, error) {

	var cpu C.sigar_proc_cpu_t
	if err := C.sigar_proc_cpu_get(h.handle, h.pid, &cpu); err != C.SIGAR_OK {
		return float64(0), errors.New(fmt.Sprintf("Fail to get CPU.  Err=%v", C.sigar_strerror(h.handle, err)))
	}

	return float64(cpu.percent), nil
}
