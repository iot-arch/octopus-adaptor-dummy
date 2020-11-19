package v1alpha1

import mqttapi "github.com/iot-arch/adaptors/dummy/octopus/pkg/mqtt/api"

// DummyDeviceExtension defines the desired state of device extension.
type DummyDeviceExtension struct {
	// Specifies the MQTT settings.
	// +optional
	MQTT *mqttapi.MQTTOptions `json:"mqtt,omitempty"`
}
