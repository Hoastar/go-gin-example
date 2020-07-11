module github.com/hoastar/go-gin-example

go 1.13

replace (
	github.com/hoastar/go-gin-example/conf => ./conf
	github.com/hoastar/go-gin-example/middleware => ./middleware
	github.com/hoastar/go-gin-example/models => ./models
	github.com/hoastar/go-gin-example/pkg/e => ./pkg/e
	github.com/hoastar/go-gin-example/pkg/setting => ./pkg/setting
	github.com/hoastar/go-gin-example/pkg/util => ./pkg/util
	github.com/hoastar/go-gin-example/routers => ./routers
)

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.1
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.57.0 // indirect
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/hoastar/go-gin-example/models v0.0.0-00010101000000-000000000000
	github.com/hoastar/go-gin-example/pkg/e v0.0.0-00010101000000-000000000000
	github.com/hoastar/go-gin-example/pkg/setting v0.0.0-00010101000000-000000000000
	github.com/hoastar/go-gin-example/pkg/util v0.0.0-00010101000000-000000000000
	github.com/hoastar/go-gin-example/routers v0.0.0-00010101000000-000000000000
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	//github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/swaggo/swag v1.6.7
	github.com/urfave/cli/v2 v2.2.0 // indirect
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sys v0.0.0-20200620081246-981b61492c35 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/tools v0.0.0-20200622150058-fcc5b64fe1f1 // indirect
	google.golang.org/protobuf v1.24.0
	gopkg.in/ini.v1 v1.57.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
