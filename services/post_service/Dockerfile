# The first stage populates the module cache based on the go.{mod,sum} files
# and builds the portal application backend
FROM golang:1.13-alpine AS build-golang
RUN apk add --update \
    git \
    gcc \
    libc-dev
WORKDIR /go/src/github.com/mewil/portal/post_service
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY pb /go/src/github.com/mewil/portal/pb
COPY common /go/src/github.com/mewil/portal/common
COPY post_service/ /go/src/github.com/mewil/portal/post_service/
RUN go mod download
RUN go install .
RUN adduser -D -g '' user

# The second stage, uses a fresh scratch image to reduce the image size and not
# ship the Go compiler in production, here we copy the statically compiled Go
# binary and use it as our entrypoint
FROM scratch AS post_service
LABEL Service="post"
LABEL Author="Michael Wilson"
COPY --from=build-golang /etc/passwd /etc/passwd
COPY --from=build-golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-golang /go/bin/post_service /bin/post_service
USER user
ENTRYPOINT ["/bin/post_service"]
EXPOSE 8003