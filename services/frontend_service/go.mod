module github.com/mewil/portal/frontend_service

go 1.13

require (
	github.com/gin-contrib/static v0.0.0-20190913125243-df30d4057ba1
	github.com/gin-contrib/zap v0.0.0-20190911144541-f473495929db
	github.com/gin-gonic/gin v1.4.0
	github.com/mewil/portal/common/logger v0.0.0
	github.com/mewil/portal/common/middleware v0.0.0
	github.com/mewil/portal/common/validation v0.0.0
	github.com/mewil/portal/frontend_service/api v0.0.0
	github.com/mewil/portal/pb v0.0.0
	go.uber.org/zap v1.10.0
)

replace github.com/mewil/portal/common/logger v0.0.0 => ../common/logger

replace github.com/mewil/portal/common/validation v0.0.0 => ../common/validation

replace github.com/mewil/portal/common/middleware v0.0.0 => ../common/middleware

replace github.com/mewil/portal/frontend_service/api v0.0.0 => ./api

replace github.com/mewil/portal/pb v0.0.0 => ../pb
