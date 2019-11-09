module github.com/micro-in-cn/starter-kit

go 1.13

require (
	github.com/casbin/casbin/v2 v2.1.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/hb-go/micro-plugins v1.14.0
	github.com/jinzhu/gorm v1.9.11
	github.com/kataras/iris/v12 v12.0.1 // indirect
	github.com/micro/go-micro v1.14.0
	github.com/micro/go-plugins v1.4.0
	github.com/micro/micro v1.14.0
	github.com/rakyll/statik v0.1.6
	github.com/sarulabs/di v2.0.0+incompatible
)

replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d
