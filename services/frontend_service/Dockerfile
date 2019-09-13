# The first stage populates the module cache based on the go.{mod,sum} files
# and builds the portal application backend
FROM golang:1.13-alpine AS build-golang
RUN apk add --update \
    git \
    gcc \
    libc-dev
WORKDIR /go/src/github.com/mewil/portal
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY ./backend /go/src/github.com/mewil/portal
RUN go mod download
RUN go install .
RUN adduser -D -g '' user

# The second stage, uses a node image to build the poral application frontend
FROM node:12.11.0-alpine AS build-node
RUN mkdir -p /usr/app
WORKDIR /usr/app
ARG NODE_ENV
ENV NODE_ENV $NODE_ENV
COPY ./frontend /usr/app
RUN yarn install
RUN yarn run build

# The third stage, uses a fresh scratch image to reduce the image size and not
# ship the Go compiler or NodeJS in production, here we copy the statically
# compiled Go binary and use it as our entrypoint
FROM scratch AS portal
LABEL Author="Michael Wilson"
COPY --from=build-node /usr/app/build/ /app/
COPY --from=build-golang /etc/passwd /etc/passwd
COPY --from=build-golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-golang /go/bin/portal /bin/portal
USER user
ENTRYPOINT ["/bin/portal"]
EXPOSE 8000