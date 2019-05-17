# Stage 0: build environment
FROM golang:alpine as builder
LABEL maintainer "Masaharu TASHIRO <masatsr.kit@gmail.com>"

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR /go/src/mshrtsr/go-gracefully-killed-in-docker
WORKDIR ${APP_DIR}
COPY . $APP_DIR

# Compile 
RUN CGO_ENABLED=0 go build -ldflags '-d -w -s' -o app ./

# Stage 1: deploy environment
FROM alpine:latest
LABEL maintainer "Masaharu TASHIRO <masatsr.kit@gmail.com>"

ENV APP_DIR /go/src/mshrtsr/go-gracefully-killed-in-docker
WORKDIR ${APP_DIR}
COPY --from=builder ${APP_DIR}/app ${APP_DIR}/app

CMD ["./app"]
