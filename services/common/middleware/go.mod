module github.com/mewil/portal/common/middleware

go 1.13

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/mewil/portal/common/logger v0.0.0
)

replace github.com/mewil/portal/common/logger v0.0.0 => ../logger
