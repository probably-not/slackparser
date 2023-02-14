# build stage
FROM golang:1.18 AS build-env
WORKDIR /go/src/github.com/probably-not/go-module-small
COPY . .

## Get Dependencies
COPY go.mod go.sum ./
RUN go mod download && go get -d -v ./...

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -tags netgo,osusergo -ldflags '-extldflags "-static"' -o go-module-small

# final stage
FROM gcr.io/distroless/static:latest
WORKDIR /app

COPY --from=build-env /go/src/github.com/probably-not/go-module-small/configs /app/configs
COPY --from=build-env /go/src/github.com/probably-not/go-module-small /app/

ENTRYPOINT ["./go-module-small"]
