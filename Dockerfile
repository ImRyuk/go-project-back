FROM golang:alpine AS build-env

ENV GO111MODULE=on

ENV CGO_ENABLED=0
ENV GOOS=$GOOS
ENV GOARCH=$GOARCH

WORKDIR ./go_app
COPY ./go_app .
RUN go mod download \
    && go build -o /build/buildedApp main/main.go

FROM alpine

WORKDIR /main
COPY --from=build-env build/buildedApp .

ENTRYPOINT [ "./buildedApp" ]
