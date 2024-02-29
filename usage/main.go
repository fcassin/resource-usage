package usage

// #cgo CFLAGS: -g -Wall
// #include "usage.h"
import "C"
import (
	"encoding/binary"
	"time"
)

type ResourceUsage struct {
	Utime time.Duration /* user CPU time used */
	Stime time.Duration /* system CPU time used */

	Maxrss   uint64 /* maximum resident set size */
	Ixrss    uint64 /* integral shared memory size */
	Idrss    uint64 /* integral unshared data size */
	Isrss    uint64 /* integral unshared stack size */
	Minflt   uint64 /* page reclaims (soft page faults) */
	Majflt   uint64 /* page faults (hard page faults) */
	Nswap    uint64 /* swaps */
	Inblock  uint64 /* block input operations */
	Outblock uint64 /* block output operations */
	Msgsnd   uint64 /* IPC messages sent */
	Msgrcv   uint64 /* IPC messages received */
	Nsignals uint64 /* signals received */
	Nvcsw    uint64 /* voluntary context switches */
	Nivcsw   uint64 /* involuntary context switches */
}

func GetResourceUsage() (resourceUsage ResourceUsage) {
	var crusage C.struct_rusage = C.ReadUsage()
	return parseCRUsage(crusage)
}

func parseCRUsage(crusage C.struct_rusage) (rusage ResourceUsage) {
	rusage.Utime = time.Duration(int64(crusage.ru_utime.tv_sec)*int64(time.Second)) +
		time.Duration(int64(crusage.ru_utime.tv_usec)*int64(time.Microsecond))

	rusage.Stime = time.Duration(int64(crusage.ru_stime.tv_sec)*int64(time.Second)) +
		time.Duration(int64(crusage.ru_stime.tv_usec)*int64(time.Microsecond))

	// NOTE: Can we dynamically determine whether we should use BigEndian or LittleEndian?
	rusage.Maxrss = binary.LittleEndian.Uint64(crusage.anon0[:])
	rusage.Ixrss = binary.LittleEndian.Uint64(crusage.anon1[:])
	rusage.Idrss = binary.LittleEndian.Uint64(crusage.anon2[:])
	rusage.Isrss = binary.LittleEndian.Uint64(crusage.anon3[:])
	rusage.Minflt = binary.LittleEndian.Uint64(crusage.anon4[:])
	rusage.Majflt = binary.LittleEndian.Uint64(crusage.anon5[:])
	rusage.Nswap = binary.LittleEndian.Uint64(crusage.anon6[:])
	rusage.Inblock = binary.LittleEndian.Uint64(crusage.anon7[:])
	rusage.Outblock = binary.LittleEndian.Uint64(crusage.anon8[:])
	rusage.Msgsnd = binary.LittleEndian.Uint64(crusage.anon9[:])
	rusage.Msgrcv = binary.LittleEndian.Uint64(crusage.anon10[:])
	rusage.Nsignals = binary.LittleEndian.Uint64(crusage.anon11[:])
	rusage.Nvcsw = binary.LittleEndian.Uint64(crusage.anon12[:])
	rusage.Nivcsw = binary.LittleEndian.Uint64(crusage.anon13[:])

	return rusage
}
