package handler

import (
	"os"
	"path/filepath"

	api "github.com/iot-arch/adaptors/dummy/octopus/pkg/adaptor/api/v1alpha1"
)

func NewPanicsCleanupSocketHandler(endpoint string) func(interface{}) {
	return func(r interface{}) {
		var socketPath = filepath.Join(api.AdaptorPath, endpoint)
		if pi, err := os.Stat(socketPath); err == nil && !pi.IsDir() && pi.Mode()&os.ModeSocket != 0 {
			_ = os.RemoveAll(socketPath)
		}
	}
}
