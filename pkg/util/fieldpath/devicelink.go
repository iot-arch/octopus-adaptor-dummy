package fieldpath

import (
	"github.com/pkg/errors"

	"github.com/iot-arch/octopus-adaptors/dummy/pkg/util/converter"
	edgev1alpha1 "github.com/rancher/octopus/api/v1alpha1"
)

// ExtractDeviceLinkFieldPathAsBytes is extracts the field from the given DeviceLink
// and returns it as a byte array.
func ExtractDeviceLinkFieldPathAsBytes(link *edgev1alpha1.DeviceLink, fieldPath string) ([]byte, error) {
	if link == nil {
		return nil, errors.New("link is nil")
	}

	var status = link.Status
	var str string
	switch fieldPath {
	case "status.nodeHostName":
		str = status.NodeHostName
	case "status.nodeInternalIP":
		str = status.NodeInternalIP
	case "status.nodeInternalDNS":
		str = status.NodeInternalDNS
	case "status.nodeExternalIP":
		str = status.NodeExternalIP
	case "status.nodeExternalDNS":
		str = status.NodeExternalDNS
	default:
		return ExtractObjectFieldPathAsBytes(link, fieldPath)
	}
	return converter.UnsafeStringToBytes(str), nil
}
