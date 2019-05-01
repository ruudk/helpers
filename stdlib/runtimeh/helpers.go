package runtimeh

import (
	"runtime"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	BlockProfileKey = "BlockProfile"

	BreakpointKey = "Breakpoint"

	CPUProfileKey = "CPUProfile"

	CallerKey = "Caller"

	CallersKey = "Callers"

	CallersFramesKey = "CallersFrames"

	FuncForPCKey = "FuncForPC"

	GCKey = "GC"

	GOARCHKey = "GOARCH"

	GOMAXPROCSKey = "GOMAXPROCS"

	GOOSKey = "GOOS"

	GOROOTKey = "GOROOT"

	GoexitKey = "Goexit"

	GoroutineProfileKey = "GoroutineProfile"

	GoschedKey = "Gosched"

	KeepAliveKey = "KeepAlive"

	LockOSThreadKey = "LockOSThread"

	MemProfileKey = "MemProfile"

	MutexProfileKey = "MutexProfile"

	NumCPUKey = "NumCPU"

	NumCgoCallKey = "NumCgoCall"

	NumGoroutineKey = "NumGoroutine"

	ReadMemStatsKey = "ReadMemStats"

	ReadTraceKey = "ReadTrace"

	SetBlockProfileRateKey = "SetBlockProfileRate"

	SetCPUProfileRateKey = "SetCPUProfileRate"

	SetCgoTracebackKey = "SetCgoTraceback"

	SetFinalizerKey = "SetFinalizer"

	SetMutexProfileFractionKey = "SetMutexProfileFraction"

	StackKey = "Stack"

	StartTraceKey = "StartTrace"

	StopTraceKey = "StopTrace"

	ThreadCreateProfileKey = "ThreadCreateProfile"

	UnlockOSThreadKey = "UnlockOSThread"

	VersionKey = "Version"
)

func New() hctx.Map {
	return hctx.Map{

		BlockProfileKey: BlockProfile,

		BreakpointKey: Breakpoint,

		CPUProfileKey: CPUProfile,

		CallerKey: Caller,

		CallersKey: Callers,

		CallersFramesKey: CallersFrames,

		FuncForPCKey: FuncForPC,

		GCKey: GC,

		GOARCHKey: GOARCH,

		GOMAXPROCSKey: GOMAXPROCS,

		GOOSKey: GOOS,

		GOROOTKey: GOROOT,

		GoexitKey: Goexit,

		GoroutineProfileKey: GoroutineProfile,

		GoschedKey: Gosched,

		KeepAliveKey: KeepAlive,

		LockOSThreadKey: LockOSThread,

		MemProfileKey: MemProfile,

		MutexProfileKey: MutexProfile,

		NumCPUKey: NumCPU,

		NumCgoCallKey: NumCgoCall,

		NumGoroutineKey: NumGoroutine,

		ReadMemStatsKey: ReadMemStats,

		ReadTraceKey: ReadTrace,

		SetBlockProfileRateKey: SetBlockProfileRate,

		SetCPUProfileRateKey: SetCPUProfileRate,

		SetCgoTracebackKey: SetCgoTraceback,

		SetFinalizerKey: SetFinalizer,

		SetMutexProfileFractionKey: SetMutexProfileFraction,

		StackKey: Stack,

		StartTraceKey: StartTrace,

		StopTraceKey: StopTrace,

		ThreadCreateProfileKey: ThreadCreateProfile,

		UnlockOSThreadKey: UnlockOSThread,

		VersionKey: Version,
	}
}

// BlockProfile returns n, the number of records in the current blocking profile.
// If len(p) &gt;= n, BlockProfile copies the profile into p and returns n, true.
// If len(p) &lt; n, BlockProfile does not change p and returns n, false.
//
// Most clients should use the runtime/pprof package or
// the testing package&#39;s -test.blockprofile flag instead
// of calling BlockProfile directly.
var BlockProfile = runtime.BlockProfile

// Breakpoint executes a breakpoint trap.
var Breakpoint = runtime.Breakpoint

// CPUProfile panics.
// It formerly provided raw access to chunks of
// a pprof-format profile generated by the runtime.
// The details of generating that format have changed,
// so this functionality has been removed.
//
// Deprecated: use the runtime/pprof package,
// or the handlers in the net/http/pprof package,
// or the testing package&#39;s -test.cpuprofile flag instead.
var CPUProfile = runtime.CPUProfile

// Caller reports file and line number information about function invocations on
// the calling goroutine&#39;s stack. The argument skip is the number of stack frames
// to ascend, with 0 identifying the caller of Caller.  (For historical reasons the
// meaning of skip differs between Caller and Callers.) The return values report the
// program counter, file name, and line number within the file of the corresponding
// call. The boolean ok is false if it was not possible to recover the information.
var Caller = runtime.Caller

// Callers fills the slice pc with the return program counters of function invocations
// on the calling goroutine&#39;s stack. The argument skip is the number of stack frames
// to skip before recording in pc, with 0 identifying the frame for Callers itself and
// 1 identifying the caller of Callers.
// It returns the number of entries written to pc.
//
// To translate these PCs into symbolic information such as function
// names and line numbers, use CallersFrames. CallersFrames accounts
// for inlined functions and adjusts the return program counters into
// call program counters. Iterating over the returned slice of PCs
// directly is discouraged, as is using FuncForPC on any of the
// returned PCs, since these cannot account for inlining or return
// program counter adjustment.
// go:noinline
var Callers = runtime.Callers

// CallersFrames takes a slice of PC values returned by Callers and
// prepares to return function/file/line information.
// Do not change the slice until you are done with the Frames.
var CallersFrames = runtime.CallersFrames

// FuncForPC returns a *Func describing the function that contains the
// given program counter address, or else nil.
//
// If pc represents multiple functions because of inlining, it returns
// the a *Func describing the innermost function, but with an entry
// of the outermost function.
var FuncForPC = runtime.FuncForPC

// GC runs a garbage collection and blocks the caller until the
// garbage collection is complete. It may also block the entire
// program.
var GC = runtime.GC

var GOARCH = runtime.GOARCH

// GOMAXPROCS sets the maximum number of CPUs that can be executing
// simultaneously and returns the previous setting. If n &lt; 1, it does not
// change the current setting.
// The number of logical CPUs on the local machine can be queried with NumCPU.
// This call will go away when the scheduler improves.
var GOMAXPROCS = runtime.GOMAXPROCS

var GOOS = runtime.GOOS

// GOROOT returns the root of the Go tree. It uses the
// GOROOT environment variable, if set at process start,
// or else the root used during the Go build.
var GOROOT = runtime.GOROOT

// Goexit terminates the goroutine that calls it. No other goroutine is affected.
// Goexit runs all deferred calls before terminating the goroutine. Because Goexit
// is not a panic, any recover calls in those deferred functions will return nil.
//
// Calling Goexit from the main goroutine terminates that goroutine
// without func main returning. Since func main has not returned,
// the program continues execution of other goroutines.
// If all other goroutines exit, the program crashes.
var Goexit = runtime.Goexit

// GoroutineProfile returns n, the number of records in the active goroutine stack profile.
// If len(p) &gt;= n, GoroutineProfile copies the profile into p and returns n, true.
// If len(p) &lt; n, GoroutineProfile does not change p and returns n, false.
//
// Most clients should use the runtime/pprof package instead
// of calling GoroutineProfile directly.
var GoroutineProfile = runtime.GoroutineProfile

// Gosched yields the processor, allowing other goroutines to run. It does not
// suspend the current goroutine, so execution resumes automatically.
var Gosched = runtime.Gosched

// KeepAlive marks its argument as currently reachable.
// This ensures that the object is not freed, and its finalizer is not run,
// before the point in the program where KeepAlive is called.
//
// A very simplified example showing where KeepAlive is required:
// 	type File struct { d int }
// 	d, err := syscall.Open(&#34;/file/path&#34;, syscall.O_RDONLY, 0)
// 	// ... do something if err != nil ...
// 	p := &amp;File{d}
// 	runtime.SetFinalizer(p, func(p *File) { syscall.Close(p.d) })
// 	var buf [10]byte
// 	n, err := syscall.Read(p.d, buf[:])
// 	// Ensure p is not finalized until Read returns.
// 	runtime.KeepAlive(p)
// 	// No more uses of p after this point.
//
// Without the KeepAlive call, the finalizer could run at the start of
// syscall.Read, closing the file descriptor before syscall.Read makes
// the actual system call.
var KeepAlive = runtime.KeepAlive

// LockOSThread wires the calling goroutine to its current operating system thread.
// The calling goroutine will always execute in that thread,
// and no other goroutine will execute in it,
// until the calling goroutine has made as many calls to
// UnlockOSThread as to LockOSThread.
// If the calling goroutine exits without unlocking the thread,
// the thread will be terminated.
//
// All init functions are run on the startup thread. Calling LockOSThread
// from an init function will cause the main function to be invoked on
// that thread.
//
// A goroutine should call LockOSThread before calling OS services or
// non-Go library functions that depend on per-thread state.
var LockOSThread = runtime.LockOSThread

// MemProfile returns a profile of memory allocated and freed per allocation
// site.
//
// MemProfile returns n, the number of records in the current memory profile.
// If len(p) &gt;= n, MemProfile copies the profile into p and returns n, true.
// If len(p) &lt; n, MemProfile does not change p and returns n, false.
//
// If inuseZero is true, the profile includes allocation records
// where r.AllocBytes &gt; 0 but r.AllocBytes == r.FreeBytes.
// These are sites where memory was allocated, but it has all
// been released back to the runtime.
//
// The returned profile may be up to two garbage collection cycles old.
// This is to avoid skewing the profile toward allocations; because
// allocations happen in real time but frees are delayed until the garbage
// collector performs sweeping, the profile only accounts for allocations
// that have had a chance to be freed by the garbage collector.
//
// Most clients should use the runtime/pprof package or
// the testing package&#39;s -test.memprofile flag instead
// of calling MemProfile directly.
var MemProfile = runtime.MemProfile

// MutexProfile returns n, the number of records in the current mutex profile.
// If len(p) &gt;= n, MutexProfile copies the profile into p and returns n, true.
// Otherwise, MutexProfile does not change p, and returns n, false.
//
// Most clients should use the runtime/pprof package
// instead of calling MutexProfile directly.
var MutexProfile = runtime.MutexProfile

// NumCPU returns the number of logical CPUs usable by the current process.
//
// The set of available CPUs is checked by querying the operating system
// at process startup. Changes to operating system CPU allocation after
// process startup are not reflected.
var NumCPU = runtime.NumCPU

// NumCgoCall returns the number of cgo calls made by the current process.
var NumCgoCall = runtime.NumCgoCall

// NumGoroutine returns the number of goroutines that currently exist.
var NumGoroutine = runtime.NumGoroutine

// ReadMemStats populates m with memory allocator statistics.
//
// The returned memory allocator statistics are up to date as of the
// call to ReadMemStats. This is in contrast with a heap profile,
// which is a snapshot as of the most recently completed garbage
// collection cycle.
var ReadMemStats = runtime.ReadMemStats

// ReadTrace returns the next chunk of binary tracing data, blocking until data
// is available. If tracing is turned off and all the data accumulated while it
// was on has been returned, ReadTrace returns nil. The caller must copy the
// returned data before calling ReadTrace again.
// ReadTrace must be called from one goroutine at a time.
var ReadTrace = runtime.ReadTrace

// SetBlockProfileRate controls the fraction of goroutine blocking events
// that are reported in the blocking profile. The profiler aims to sample
// an average of one blocking event per rate nanoseconds spent blocked.
//
// To include every blocking event in the profile, pass rate = 1.
// To turn off profiling entirely, pass rate &lt;= 0.
var SetBlockProfileRate = runtime.SetBlockProfileRate

// SetCPUProfileRate sets the CPU profiling rate to hz samples per second.
// If hz &lt;= 0, SetCPUProfileRate turns off profiling.
// If the profiler is on, the rate cannot be changed without first turning it off.
//
// Most clients should use the runtime/pprof package or
// the testing package&#39;s -test.cpuprofile flag instead of calling
// SetCPUProfileRate directly.
var SetCPUProfileRate = runtime.SetCPUProfileRate

// SetCgoTraceback records three C functions to use to gather
// traceback information from C code and to convert that traceback
// information into symbolic information. These are used when printing
// stack traces for a program that uses cgo.
//
// The traceback and context functions may be called from a signal
// handler, and must therefore use only async-signal safe functions.
// The symbolizer function may be called while the program is
// crashing, and so must be cautious about using memory.  None of the
// functions may call back into Go.
//
// The context function will be called with a single argument, a
// pointer to a struct:
//
// 	struct {
// 		Context uintptr
// 	}
//
// In C syntax, this struct will be
//
// 	struct {
// 		uintptr_t Context;
// 	};
//
// If the Context field is 0, the context function is being called to
// record the current traceback context. It should record in the
// Context field whatever information is needed about the current
// point of execution to later produce a stack trace, probably the
// stack pointer and PC. In this case the context function will be
// called from C code.
//
// If the Context field is not 0, then it is a value returned by a
// previous call to the context function. This case is called when the
// context is no longer needed; that is, when the Go code is returning
// to its C code caller. This permits the context function to release
// any associated resources.
//
// While it would be correct for the context function to record a
// complete a stack trace whenever it is called, and simply copy that
// out in the traceback function, in a typical program the context
// function will be called many times without ever recording a
// traceback for that context. Recording a complete stack trace in a
// call to the context function is likely to be inefficient.
//
// The traceback function will be called with a single argument, a
// pointer to a struct:
//
// 	struct {
// 		Context    uintptr
// 		SigContext uintptr
// 		Buf        *uintptr
// 		Max        uintptr
// 	}
//
// In C syntax, this struct will be
//
// 	struct {
// 		uintptr_t  Context;
// 		uintptr_t  SigContext;
// 		uintptr_t* Buf;
// 		uintptr_t  Max;
// 	};
//
// The Context field will be zero to gather a traceback from the
// current program execution point. In this case, the traceback
// function will be called from C code.
//
// Otherwise Context will be a value previously returned by a call to
// the context function. The traceback function should gather a stack
// trace from that saved point in the program execution. The traceback
// function may be called from an execution thread other than the one
// that recorded the context, but only when the context is known to be
// valid and unchanging. The traceback function may also be called
// deeper in the call stack on the same thread that recorded the
// context. The traceback function may be called multiple times with
// the same Context value; it will usually be appropriate to cache the
// result, if possible, the first time this is called for a specific
// context value.
//
// If the traceback function is called from a signal handler on a Unix
// system, SigContext will be the signal context argument passed to
// the signal handler (a C ucontext_t* cast to uintptr_t). This may be
// used to start tracing at the point where the signal occurred. If
// the traceback function is not called from a signal handler,
// SigContext will be zero.
//
// Buf is where the traceback information should be stored. It should
// be PC values, such that Buf[0] is the PC of the caller, Buf[1] is
// the PC of that function&#39;s caller, and so on.  Max is the maximum
// number of entries to store.  The function should store a zero to
// indicate the top of the stack, or that the caller is on a different
// stack, presumably a Go stack.
//
// Unlike runtime.Callers, the PC values returned should, when passed
// to the symbolizer function, return the file/line of the call
// instruction.  No additional subtraction is required or appropriate.
//
// On all platforms, the traceback function is invoked when a call from
// Go to C to Go requests a stack trace. On linux/amd64, linux/ppc64le,
// and freebsd/amd64, the traceback function is also invoked when a
// signal is received by a thread that is executing a cgo call. The
// traceback function should not make assumptions about when it is
// called, as future versions of Go may make additional calls.
//
// The symbolizer function will be called with a single argument, a
// pointer to a struct:
//
// 	struct {
// 		PC      uintptr // program counter to fetch information for
// 		File    *byte   // file name (NUL terminated)
// 		Lineno  uintptr // line number
// 		Func    *byte   // function name (NUL terminated)
// 		Entry   uintptr // function entry point
// 		More    uintptr // set non-zero if more info for this PC
// 		Data    uintptr // unused by runtime, available for function
// 	}
//
// In C syntax, this struct will be
//
// 	struct {
// 		uintptr_t PC;
// 		char*     File;
// 		uintptr_t Lineno;
// 		char*     Func;
// 		uintptr_t Entry;
// 		uintptr_t More;
// 		uintptr_t Data;
// 	};
//
// The PC field will be a value returned by a call to the traceback
// function.
//
// The first time the function is called for a particular traceback,
// all the fields except PC will be 0. The function should fill in the
// other fields if possible, setting them to 0/nil if the information
// is not available. The Data field may be used to store any useful
// information across calls. The More field should be set to non-zero
// if there is more information for this PC, zero otherwise. If More
// is set non-zero, the function will be called again with the same
// PC, and may return different information (this is intended for use
// with inlined functions). If More is zero, the function will be
// called with the next PC value in the traceback. When the traceback
// is complete, the function will be called once more with PC set to
// zero; this may be used to free any information. Each call will
// leave the fields of the struct set to the same values they had upon
// return, except for the PC field when the More field is zero. The
// function must not keep a copy of the struct pointer between calls.
//
// When calling SetCgoTraceback, the version argument is the version
// number of the structs that the functions expect to receive.
// Currently this must be zero.
//
// The symbolizer function may be nil, in which case the results of
// the traceback function will be displayed as numbers. If the
// traceback function is nil, the symbolizer function will never be
// called. The context function may be nil, in which case the
// traceback function will only be called with the context field set
// to zero.  If the context function is nil, then calls from Go to C
// to Go will not show a traceback for the C portion of the call stack.
//
// SetCgoTraceback should be called only once, ideally from an init function.
var SetCgoTraceback = runtime.SetCgoTraceback

// SetFinalizer sets the finalizer associated with obj to the provided
// finalizer function. When the garbage collector finds an unreachable block
// with an associated finalizer, it clears the association and runs
// finalizer(obj) in a separate goroutine. This makes obj reachable again,
// but now without an associated finalizer. Assuming that SetFinalizer
// is not called again, the next time the garbage collector sees
// that obj is unreachable, it will free obj.
//
// SetFinalizer(obj, nil) clears any finalizer associated with obj.
//
// The argument obj must be a pointer to an object allocated by calling
// new, by taking the address of a composite literal, or by taking the
// address of a local variable.
// The argument finalizer must be a function that takes a single argument
// to which obj&#39;s type can be assigned, and can have arbitrary ignored return
// values. If either of these is not true, SetFinalizer may abort the
// program.
//
// Finalizers are run in dependency order: if A points at B, both have
// finalizers, and they are otherwise unreachable, only the finalizer
// for A runs; once A is freed, the finalizer for B can run.
// If a cyclic structure includes a block with a finalizer, that
// cycle is not guaranteed to be garbage collected and the finalizer
// is not guaranteed to run, because there is no ordering that
// respects the dependencies.
//
// The finalizer is scheduled to run at some arbitrary time after the
// program can no longer reach the object to which obj points.
// There is no guarantee that finalizers will run before a program exits,
// so typically they are useful only for releasing non-memory resources
// associated with an object during a long-running program.
// For example, an os.File object could use a finalizer to close the
// associated operating system file descriptor when a program discards
// an os.File without calling Close, but it would be a mistake
// to depend on a finalizer to flush an in-memory I/O buffer such as a
// bufio.Writer, because the buffer would not be flushed at program exit.
//
// It is not guaranteed that a finalizer will run if the size of *obj is
// zero bytes.
//
// It is not guaranteed that a finalizer will run for objects allocated
// in initializers for package-level variables. Such objects may be
// linker-allocated, not heap-allocated.
//
// A finalizer may run as soon as an object becomes unreachable.
// In order to use finalizers correctly, the program must ensure that
// the object is reachable until it is no longer required.
// Objects stored in global variables, or that can be found by tracing
// pointers from a global variable, are reachable. For other objects,
// pass the object to a call of the KeepAlive function to mark the
// last point in the function where the object must be reachable.
//
// For example, if p points to a struct that contains a file descriptor d,
// and p has a finalizer that closes that file descriptor, and if the last
// use of p in a function is a call to syscall.Write(p.d, buf, size), then
// p may be unreachable as soon as the program enters syscall.Write. The
// finalizer may run at that moment, closing p.d, causing syscall.Write
// to fail because it is writing to a closed file descriptor (or, worse,
// to an entirely different file descriptor opened by a different goroutine).
// To avoid this problem, call runtime.KeepAlive(p) after the call to
// syscall.Write.
//
// A single goroutine runs all finalizers for a program, sequentially.
// If a finalizer must run for a long time, it should do so by starting
// a new goroutine.
var SetFinalizer = runtime.SetFinalizer

// SetMutexProfileFraction controls the fraction of mutex contention events
// that are reported in the mutex profile. On average 1/rate events are
// reported. The previous rate is returned.
//
// To turn off profiling entirely, pass rate 0.
// To just read the current rate, pass rate &lt; 0.
// (For n&gt;1 the details of sampling may change.)
var SetMutexProfileFraction = runtime.SetMutexProfileFraction

// Stack formats a stack trace of the calling goroutine into buf
// and returns the number of bytes written to buf.
// If all is true, Stack formats stack traces of all other goroutines
// into buf after the trace for the current goroutine.
var Stack = runtime.Stack

// StartTrace enables tracing for the current process.
// While tracing, the data will be buffered and available via ReadTrace.
// StartTrace returns an error if tracing is already enabled.
// Most clients should use the runtime/trace package or the testing package&#39;s
// -test.trace flag instead of calling StartTrace directly.
var StartTrace = runtime.StartTrace

// StopTrace stops tracing, if it was previously enabled.
// StopTrace only returns after all the reads for the trace have completed.
var StopTrace = runtime.StopTrace

// ThreadCreateProfile returns n, the number of records in the thread creation profile.
// If len(p) &gt;= n, ThreadCreateProfile copies the profile into p and returns n, true.
// If len(p) &lt; n, ThreadCreateProfile does not change p and returns n, false.
//
// Most clients should use the runtime/pprof package instead
// of calling ThreadCreateProfile directly.
var ThreadCreateProfile = runtime.ThreadCreateProfile

// UnlockOSThread undoes an earlier call to LockOSThread.
// If this drops the number of active LockOSThread calls on the
// calling goroutine to zero, it unwires the calling goroutine from
// its fixed operating system thread.
// If there are no active LockOSThread calls, this is a no-op.
//
// Before calling UnlockOSThread, the caller must ensure that the OS
// thread is suitable for running other goroutines. If the caller made
// any permanent changes to the state of the thread that would affect
// other goroutines, it should not call this function and thus leave
// the goroutine locked to the OS thread until the goroutine (and
// hence the thread) exits.
var UnlockOSThread = runtime.UnlockOSThread

// Version returns the Go tree&#39;s version string.
// It is either the commit hash and date at the time of the build or,
// when possible, a release tag like &#34;go1.3&#34;.
var Version = runtime.Version
