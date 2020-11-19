package metrics

import (
	"github.com/iot-arch/adaptors/dummy/octopus/pkg/metrics/limb"
)

// alias limb package
var (
	RegisterLimbMetrics    = limb.RegisterMetrics
	GetLimbMetricsRecorder = limb.GetMetricsRecorder
)
