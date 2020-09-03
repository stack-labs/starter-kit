package profile

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v3/auth/noop"
	"github.com/micro/go-micro/v3/broker"
	"github.com/micro/go-micro/v3/broker/http"
	"github.com/micro/go-micro/v3/client"
	"github.com/micro/go-micro/v3/config"
	memStream "github.com/micro/go-micro/v3/events/stream/memory"
	"github.com/micro/go-micro/v3/registry"
	"github.com/micro/go-micro/v3/registry/etcd"
	"github.com/micro/go-micro/v3/router"
	regRouter "github.com/micro/go-micro/v3/router/registry"
	"github.com/micro/go-micro/v3/runtime/local"
	"github.com/micro/go-micro/v3/server"
	"github.com/micro/go-micro/v3/store/file"

	//inAuth "github.com/micro/micro/v3/internal/auth"
	mProfile "github.com/micro/micro/v3/profile"
	microAuth "github.com/micro/micro/v3/service/auth"
	microBroker "github.com/micro/micro/v3/service/broker"
	microClient "github.com/micro/micro/v3/service/client"
	microConfig "github.com/micro/micro/v3/service/config"
	microEvents "github.com/micro/micro/v3/service/events"
	"github.com/micro/micro/v3/service/logger"
	microRegistry "github.com/micro/micro/v3/service/registry"
	microRouter "github.com/micro/micro/v3/service/router"
	microRuntime "github.com/micro/micro/v3/service/runtime"
	microServer "github.com/micro/micro/v3/service/server"
	microStore "github.com/micro/micro/v3/service/store"
)

func init() {
	mProfile.Register("dev", Dev)
}

// Dev profile to run develop env
var Dev = &mProfile.Profile{
	Name: "dev",
	Setup: func(ctx *cli.Context) error {
		microAuth.DefaultAuth = noop.NewAuth()
		microRuntime.DefaultRuntime = local.NewRuntime()
		microStore.DefaultStore = file.NewStore()
		microConfig.DefaultConfig, _ = config.NewConfig()
		setBroker(http.NewBroker())
		setRegistry(etcd.NewRegistry())
		setupJWTRules()

		var err error
		microEvents.DefaultStream, err = memStream.NewStream()
		if err != nil {
			logger.Fatalf("Error configuring stream: %v", err)
		}

		return nil
	},
}

func setRegistry(reg registry.Registry) {
	microRegistry.DefaultRegistry = reg
	microRouter.DefaultRouter = regRouter.NewRouter(router.Registry(reg))
	microServer.DefaultServer.Init(server.Registry(reg))
	microClient.DefaultClient.Init(client.Registry(reg))
}

func setBroker(b broker.Broker) {
	microBroker.DefaultBroker = b
	microClient.DefaultClient.Init(client.Broker(b))
	microServer.DefaultServer.Init(server.Broker(b))
}

func setupJWTRules() {
	//for _, rule := range inAuth.SystemRules {
	//	if err := microAuth.DefaultAuth.Grant(rule); err != nil {
	//		logger.Fatal("Error creating default rule: %v", err)
	//	}
	//}
}
