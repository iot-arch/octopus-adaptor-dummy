// +build test

package content

import (
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/iot-arch/octopus-adaptors/dummy/pkg/util/converter"
)

func ToRawExtension(content interface{}) *runtime.RawExtension {
	if content == nil {
		return nil
	}
	switch t := content.(type) {
	case []byte:
		return &runtime.RawExtension{Raw: t}
	case string:
		return &runtime.RawExtension{Raw: []byte(t)}
	default:
		return &runtime.RawExtension{Raw: converter.TryMarshalJSON(content)}
	}
}
