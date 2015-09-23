package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("(?P<Date>.*) (?P<Milliseconds>[0-9]+) jvm.JvmMetrics: Context=jvm, ProcessName=(?P<ProcessName>.*), SessionId=(?P<SessionId>.*), Hostname=(?P<Hostname>.*), MemNonHeapUsedM=(?P<MemNonHeapUsedM>.*), MemNonHeapCommittedM=(?P<MemNonHeapCommittedM>.*), MemNonHeapMaxM=(?P<MemNonHeapMaxM>.*), MemHeapUsedM=(?P<MemHeapUsedM>.*), MemHeapCommittedM=(?P<MemHeapCommittedM>.*), MemHeapMaxM=(?P<MemHeapMaxM>.*), MemMaxM=(?P<MemMaxM>.*), GcCountCopy=(?P<GcCountCopy>[0-9]+), GcTimeMillisCopy=(?P<GcTimeMillisCopy>[0-9]+), GcCountMarkSweepCompact=(?P<GcCountMarkSweepCompact>[0-9]+), GcTimeMillisMarkSweepCompact=(?P<GcTimeMillisMarkSweepCompact>[0-9]+), GcCount=(?P<GcCount>[0-9]+), GcTimeMillis=(?P<GcTimeMillis>[0-9]+), ThreadsNew=(?P<ThreadsNew>[0-9]+), ThreadsRunnable=(?P<ThreadsRunnable>[0-9]+), ThreadsBlocked=(?P<ThreadsBlocked>[0-9]+), ThreadsWaiting=(?P<ThreadsWaiting>[0-9]+), ThreadsTimedWaiting=(?P<ThreadsTimedWaiting>[0-9]+), ThreadsTerminated=(?P<ThreadsTerminated>[0-9]+), LogFatal=(?P<LogFatal>[0-9]+), LogError=(?P<LogError>[0-9]+), LogWarn=(?P<LogWarn>[0-9]+), LogInfo=(?P<LogInfo>[0-9]+)")
	n1 := re.SubexpNames()
	r2 := re.FindAllStringSubmatch("2015/09/23 14:56:33 1443020193960 jvm.JvmMetrics: Context=jvm, ProcessName=NodeManager, SessionId=null, Hostname=vagrant-ubuntu-trusty-64, MemNonHeapUsedM=32.380257, MemNonHeapCommittedM=33.421875, MemNonHeapMaxM=-9.536743E-7, MemHeapUsedM=10.907906, MemHeapCommittedM=33.125, MemHeapMaxM=966.6875, MemMaxM=966.6875, GcCountCopy=212, GcTimeMillisCopy=291, GcCountMarkSweepCompact=1, GcTimeMillisMarkSweepCompact=28, GcCount=213, GcTimeMillis=319, ThreadsNew=0, ThreadsRunnable=15, ThreadsBlocked=0, ThreadsWaiting=11, ThreadsTimedWaiting=35, ThreadsTerminated=0, LogFatal=0, LogError=0, LogWarn=2, LogInfo=64", -1)[0]

	md := map[string]string{}
	for i, n := range r2 {
		md[n1[i]] = n
	}

	fmt.Printf("LogWarn is %s\n", md)
	fmt.Printf("Milliseconds is %s\n", md["Milliseconds"])
	fmt.Printf("LogWarn is %s\n", md["LogWarn"])
	fmt.Printf("LogInfo is %s\n", md["LogInfo"])
	fmt.Printf("MemNonHeapUsedM is %s\n", md["MemNonHeapUsedM"])

}
