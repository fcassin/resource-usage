package usage

// #cgo CFLAGS: -g -Wall
// #include "usage.h"
import "C"
import (
	"encoding/binary"
	"time"
)

type ResourceUsage struct {
	ruUtime time.Duration /* user CPU time used */
	ruStime time.Duration /* system CPU time used */

	ruMaxrss   uint64 /* maximum resident set size */
	ruIxrss    uint64 /* integral shared memory size */
	ruIdrss    uint64 /* integral unshared data size */
	ruIsrss    uint64 /* integral unshared stack size */
	ruMinflt   uint64 /* page reclaims (soft page faults) */
	ruMajflt   uint64 /* page faults (hard page faults) */
	ruNswap    uint64 /* swaps */
	ruInblock  uint64 /* block input operations */
	ruOutblock uint64 /* block output operations */
	ruMsgsnd   uint64 /* IPC messages sent */
	ruMsgrcv   uint64 /* IPC messages received */
	ruNsignals uint64 /* signals received */
	ruNvcsw    uint64 /* voluntary context switches */
	ruNivcsw   uint64 /* involuntary context switches */
}

func GetResourceUsage() (resourceUsage ResourceUsage) {
	var crusage C.struct_rusage = C.ReadUsage()
	return parseCRUsage(crusage)
}

func parseCRUsage(crusage C.struct_rusage) (rusage ResourceUsage) {
	rusage.ruUtime = time.Duration(int64(crusage.ru_utime.tv_sec)*int64(time.Second)) +
		time.Duration(int64(crusage.ru_utime.tv_usec)*int64(time.Microsecond))

	rusage.ruStime = time.Duration(int64(crusage.ru_stime.tv_sec)*int64(time.Second)) +
		time.Duration(int64(crusage.ru_stime.tv_usec)*int64(time.Microsecond))

	// NOTE: Can we dynamically determine whether we should use BigEndian or LittleEndian?
	rusage.ruMaxrss = binary.LittleEndian.Uint64(crusage.anon0[:])
	rusage.ruIxrss = binary.LittleEndian.Uint64(crusage.anon1[:])
	rusage.ruIdrss = binary.LittleEndian.Uint64(crusage.anon2[:])
	rusage.ruIsrss = binary.LittleEndian.Uint64(crusage.anon3[:])
	rusage.ruMinflt = binary.LittleEndian.Uint64(crusage.anon4[:])
	rusage.ruMajflt = binary.LittleEndian.Uint64(crusage.anon5[:])
	rusage.ruNswap = binary.LittleEndian.Uint64(crusage.anon6[:])
	rusage.ruInblock = binary.LittleEndian.Uint64(crusage.anon7[:])
	rusage.ruOutblock = binary.LittleEndian.Uint64(crusage.anon8[:])
	rusage.ruMsgsnd = binary.LittleEndian.Uint64(crusage.anon9[:])
	rusage.ruMsgrcv = binary.LittleEndian.Uint64(crusage.anon10[:])
	rusage.ruNsignals = binary.LittleEndian.Uint64(crusage.anon11[:])
	rusage.ruNvcsw = binary.LittleEndian.Uint64(crusage.anon12[:])
	rusage.ruNivcsw = binary.LittleEndian.Uint64(crusage.anon13[:])

	return rusage
}
