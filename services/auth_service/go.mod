module github.com/mewil/portal/auth_service

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/didip/tollbooth v4.0.2+incompatible // indirect
	github.com/didip/tollbooth_gin v0.0.0-20170928041415-5752492be505 // indirect
	github.com/gin-gonic/gin v1.4.0 // indirect
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/mewil/portal/common/database v0.0.0
	github.com/mewil/portal/common/grpc_utils v0.0.0
	github.com/mewil/portal/common/logger v0.0.0
	github.com/mewil/portal/common/validation v0.0.0
	github.com/mewil/portal/pb v0.0.0
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/time v0.0.0-20190921001708-c4c64cad1fd0 // indirect
	google.golang.org/grpc v1.24.0
)

replace github.com/mewil/portal/common/logger v0.0.0 => ../common/logger

replace github.com/mewil/portal/common/grpc_utils v0.0.0 => ../common/grpc_utils

replace github.com/mewil/portal/common/database v0.0.0 => ../common/database

replace github.com/mewil/portal/common/validation v0.0.0 => ../common/validation

replace github.com/mewil/portal/pb v0.0.0 => ../pb
