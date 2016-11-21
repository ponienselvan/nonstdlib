package executor

import (
	"flag"
	"runtime"
	"testing"
	"time"

	tbnflag "github.com/turbinelabs/nonstdlib/flag"
	"github.com/turbinelabs/test/assert"
	"github.com/turbinelabs/test/log"
)

func TestFromFlags(t *testing.T) {
	log := log.NewNoopLogger()

	flagSet := flag.NewFlagSet("executor", flag.PanicOnError)

	prefixedFlagSet := tbnflag.NewPrefixedFlagSet(flagSet, "exec", "whatever")

	ff := NewFromFlags(prefixedFlagSet)
	assert.NonNil(t, ff)

	ffImpl := ff.(*fromFlags)

	expectedMaxQueueDepth := runtime.NumCPU() * 20
	expectedParallelism := runtime.NumCPU() * 2

	assert.Equal(t, ffImpl.delayType.String(), exponentialDelayType)
	assert.Equal(t, ffImpl.initialDelay, 100*time.Millisecond)
	assert.Equal(t, ffImpl.maxDelay, 30*time.Second)
	assert.Equal(t, ffImpl.maxAttempts, 8)
	assert.Equal(t, ffImpl.maxQueueDepth, expectedMaxQueueDepth)
	assert.Equal(t, ffImpl.parallelism, expectedParallelism)
	assert.Nil(t, ffImpl.executor)

	exec := ff.Make(log)
	assert.SameInstance(t, exec, ffImpl.executor)
	assert.SameInstance(t, ff.Make(log), exec)

	execImpl := exec.(*retryingExec)
	assert.NonNil(t, execImpl.deadlineChan)
	assert.Equal(t, cap(execImpl.execChan), expectedMaxQueueDepth)
	assert.Equal(t, execImpl.parallelism, expectedParallelism)
	assert.Equal(t, execImpl.maxQueueDepth, expectedMaxQueueDepth)
	assert.Equal(t, execImpl.maxAttempts, 8)
	assert.NonNil(t, execImpl.delay)
	assert.Equal(t, execImpl.delay(1), 100*time.Millisecond)
	assert.Equal(t, execImpl.delay(100000), 30*time.Second)
	assert.Equal(t, execImpl.timeout, 5*time.Second)
	assert.SameInstance(t, execImpl.log, log)

	exec.Stop()

	ffImpl.executor = nil

	flagSet.Parse([]string{
		"-exec.delay-type=constant",
		"-exec.delay=1s",
		"-exec.max-delay=5s",
		"-exec.max-attempts=4",
		"-exec.max-queue=128",
		"-exec.parallelism=99",
		"-exec.timeout=100ms",
	})

	assert.Equal(t, ffImpl.delayType.String(), constantDelayType)
	assert.Equal(t, ffImpl.initialDelay, time.Second)
	assert.Equal(t, ffImpl.maxDelay, 5*time.Second)
	assert.Equal(t, ffImpl.maxAttempts, 4)
	assert.Equal(t, ffImpl.maxQueueDepth, 128)
	assert.Equal(t, ffImpl.parallelism, 99)
	assert.Equal(t, ffImpl.timeout, 100*time.Millisecond)

	expectedMaxQueueDepth = 128
	expectedParallelism = 99

	exec = ff.Make(nil)
	assert.SameInstance(t, exec, ffImpl.executor)

	execImpl = exec.(*retryingExec)
	assert.NonNil(t, execImpl.deadlineChan)
	assert.Equal(t, cap(execImpl.execChan), expectedMaxQueueDepth)
	assert.Equal(t, execImpl.parallelism, expectedParallelism)
	assert.Equal(t, execImpl.maxQueueDepth, expectedMaxQueueDepth)
	assert.Equal(t, execImpl.maxAttempts, 4)
	assert.NonNil(t, execImpl.delay)
	assert.Equal(t, execImpl.delay(1), 1*time.Second)
	assert.Equal(t, execImpl.delay(100000), 1*time.Second)
	assert.Equal(t, execImpl.timeout, 100*time.Millisecond)
	assert.Nil(t, execImpl.log)

	exec.Stop()
}
