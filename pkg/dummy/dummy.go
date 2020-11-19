package dummy

import (
	"golang.org/x/sync/errgroup"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/iot-arch/adaptors/dummy/pkg/adaptor"
	api "github.com/iot-arch/adaptors/dummy/pkg/adaptor/api/v1alpha1"
	"github.com/iot-arch/adaptors/dummy/pkg/adaptor/connection"
	"github.com/iot-arch/adaptors/dummy/pkg/adaptor/log"
	"github.com/iot-arch/adaptors/dummy/pkg/adaptor/registration"
	"github.com/iot-arch/adaptors/dummy/pkg/metadata"
	"github.com/iot-arch/adaptors/dummy/pkg/util/critical"
)

// +kubebuilder:rbac:groups=devices.edge.cattle.io,resources=dummyspecialdevices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=devices.edge.cattle.io,resources=dummyspecialdevices/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=devices.edge.cattle.io,resources=dummyprotocoldevices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=devices.edge.cattle.io,resources=dummyprotocoldevices/status,verbs=get;update;patch

func Run() error {
	log.Info("Starting")

	var stop = ctrl.SetupSignalHandler()
	var ctx = critical.Context(stop)
	eg, ctx := errgroup.WithContext(ctx)
	stop = ctx.Done()
	eg.Go(func() error {
		// start adaptor to receive requests from Limb
		return connection.Serve(metadata.Endpoint, adaptor.NewService(), stop)
	})
	eg.Go(func() error {
		// register adaptor to Limb
		return registration.Register(ctx, api.RegisterRequest{
			Name:     metadata.Name,
			Version:  metadata.Version,
			Endpoint: metadata.Endpoint,
		})
	})
	return eg.Wait()
}
