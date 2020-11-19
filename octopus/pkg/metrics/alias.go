package metrics

import (
	"github.com/iot-arch/adaptors/dummy/pkg/metrics/limb"
)

// alias limb package
var (
	RegisterLimbMetrics    = limb.RegisterMetrics
	GetLimbMetricsRecorder = limb.GetMetricsRecorder
)
