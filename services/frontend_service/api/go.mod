module github.com/mewil/portal/frontend_service/api

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/didip/tollbooth v4.0.2+incompatible
	github.com/didip/tollbooth_gin v0.0.0-20170928041415-5752492be505
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/google/uuid v1.1.1
	github.com/json-iterator/go v1.1.7 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/mewil/portal/common/logger v0.0.0
	github.com/mewil/portal/common/validation v0.0.0
	github.com/mewil/portal/pb v0.0.0
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/stretchr/testify v1.4.0
	github.com/ugorji/go v1.1.7 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/sys v0.0.0-20191009170203-06d7bd2c5f4f // indirect
	golang.org/x/time v0.0.0-20190921001708-c4c64cad1fd0 // indirect
	google.golang.org/grpc v1.24.0
)

replace github.com/mewil/portal/common/logger v0.0.0 => ../../common/logger

replace github.com/mewil/portal/common/validation v0.0.0 => ../../common/validation

replace github.com/mewil/portal/pb v0.0.0 => ../../pb
