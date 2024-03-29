FROM pandoc/core:2.17.0 as pandoc-builder

# Backend build ################################################################
# https://hub.docker.com/_/golang
FROM golang:1.19.2-alpine3.16 as backend-builder

RUN apk add --no-cache git
WORKDIR /go/src/com.nguyenonline/formipro
COPY . .
RUN go get -v -t .
RUN go build -o app

# App image ####################################################################
FROM nguyen99/alpine-latex:20221013190916

COPY --from=pandoc-builder \
  /usr/local/bin/pandoc \
  /usr/local/bin/

# Reinstall any system packages required for runtime.
RUN apk --no-cache add \
        gmp \
        libffi \
        lua5.3 \
        lua5.3-lpeg

ENV APP_HOME /go/src/com.nguyenonline/formipro
WORKDIR $APP_HOME

COPY --from=backend-builder /go/src/com.nguyenonline/formipro/app .
COPY assets assets
RUN mkdir tmp

EXPOSE 8080

CMD ["./app"]
