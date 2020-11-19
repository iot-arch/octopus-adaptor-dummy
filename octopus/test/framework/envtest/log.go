// +build test

package envtest

import (
	"github.com/iot-arch/adaptors/dummy/octopus/pkg/util/log/zap"
)

var log = zap.WrapAsLogr(zap.NewDevelopmentLogger()).WithName("test-env")
