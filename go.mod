module github.com/micro-in-cn/starter-kit

go 1.13

require (
	cloud.google.com/go v0.50.0
	contrib.go.opencensus.io/exporter/jaeger v0.2.0
	contrib.go.opencensus.io/exporter/stackdriver v0.12.1
	github.com/ajg/form v1.5.1 // indirect
	github.com/alibaba/sentinel-golang v0.3.0
	github.com/astaxie/beego v1.12.0
	github.com/casbin/casbin/v2 v2.1.2
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/fasthttp-contrib/websocket v0.0.0-20160511215533-1f3b11f56072 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/gin-gonic/gin v1.5.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1
	github.com/gopherjs/gopherjs v0.0.0-20190430165422-3e4dfb77656c // indirect
	github.com/gorilla/mux v1.7.3
	github.com/hb-go/micro-plugins/v3 v3.0.0-20200830143901-151e7b6a0a8d
	github.com/hb-go/pkg v0.0.2
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/jinzhu/gorm v1.9.11
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/iris/v12 v12.0.1
	github.com/labstack/echo/v4 v4.1.15
	github.com/micro-in-cn/x-gateway v0.0.5-0.20200218143744-d53812e7b074
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.4.1-0.20200413220547-71d4253927cc
	github.com/micro/go-micro/v3 v3.0.0-beta.0.20200902122854-6bdf33c4eede
	github.com/micro/go-plugins/logger/zap/v2 v2.3.0
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.0.0-20200317215710-66384449b09c
	github.com/micro/go-plugins/transport/tcp/v2 v2.0.0-20200317215710-66384449b09c
	github.com/micro/micro/v2 v2.4.0
	github.com/micro/micro/v3 v3.0.0-beta
	github.com/micro/services v0.10.0 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.0
	github.com/rakyll/statik v0.1.6
	github.com/rs/cors v1.7.0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/stretchr/testify v1.5.1
	github.com/uber/jaeger-client-go v2.21.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible
	github.com/valyala/fasthttp v1.6.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	go.opencensus.io v0.22.2
	go.uber.org/dig v1.9.0
	go.uber.org/zap v1.13.0
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	google.golang.org/grpc v1.27.0
	xorm.io/xorm v0.8.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/micro/micro/v2 v2.4.0 => github.com/hb-chen/micro/v2 v2.0.0-20200414123212-977f933825b7

replace github.com/micro/micro/v3 v3.0.0-beta => github.com/hb-chen/micro/v3 v3.0.0-20200903160822-d37c7b88b9ae
