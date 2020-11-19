// +build test

package envtest

import (
	"github.com/iot-arch/octopus-adaptors/dummy/pkg/util/log/zap"
)

var log = zap.WrapAsLogr(zap.NewDevelopmentLogger()).WithName("test-env")
