// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

// +build linux

package utils

import (
	"runtime"

	"github.com/DataDog/datadog-agent/pkg/util"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

// MemoryMonitor monitors memory cgroup usage
type MemoryMonitor = util.MemoryController

// MemoryMonitor instantiates a new memory monitor
func NewMemoryMonitor(memoryPressureLevel string, memoryThreshold int) (*MemoryMonitor, error) {
	var memoryMonitors []util.MemoryMonitor
	if memoryPressureLevel != "" {
		memoryMonitors = append(memoryMonitors, util.MemoryPressureMonitor(func() {
			log.Infof("Memory pressure reached level '%s', triggering garbage collector", memoryPressureLevel)
			runtime.GC()
		}, memoryPressureLevel))
	}

	if memoryThreshold > 0 {
		memoryMonitors = append(memoryMonitors, util.MemoryThresholdMonitor(func() {
			log.Infof("Memory pressure above %d% threshold, triggering garbage collector", memoryThreshold)
			runtime.GC()
		}, uint64(memoryThreshold), false))
	}

	return util.NewMemoryController(memoryMonitors...)
}
