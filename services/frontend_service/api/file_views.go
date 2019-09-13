package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mewil/portal/common/validation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetFile handles GET /v1/file/:file_id
func (s *FrontendSvc) GetFile(newFileSvcClient FileSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileID := c.Param("file_id")
		if err := validation.ValidUUID(fileID); err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid file id")
			return
		}
		stats, err := s.FileSvcGetFileStats(c.Request.Context(), newFileSvcClient, fileID)
		if match := c.Request.Header.Get("If-None-Match"); match != "" {
			if strings.Contains(match, stats.GetETag()) {
				ResponseValid(c, http.StatusNotModified, "file not modified", gin.H{})
				return
			}
		}
		body, err := s.FileSvcGetFile(c.Request.Context(), newFileSvcClient, fileID)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			c.Header("Etag", stats.GetETag())
			c.Header("Cache-Control", "public,max-age=86400")
			c.Header("Content-Length", fmt.Sprintf("%d", len(body)))
			c.Writer.WriteHeader(http.StatusOK)
			c.Writer.Write(body)
			c.Writer.Flush()
		case codes.InvalidArgument:
			ResponseError(c, http.StatusBadRequest, "please provide a valid file id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}
