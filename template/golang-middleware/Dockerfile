FROM openfaas/of-watchdog:0.7.2 as watchdog

FROM golang:1.13.1-alpine3.10 as build

RUN apk --no-cache add git gcc libc-dev

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

ENV CGO_ENABLED=1

RUN mkdir -p /go/src/handler
WORKDIR /go/src/handler
COPY . .


ARG GO111MODULE="on"
ARG GOPROXY=""

# Run go test, gofmt, go vet
RUN chmod +x test.sh \
    && ./test.sh \
    && go build --ldflags "-s -w" -a -installsuffix cgo -o handler .

FROM alpine:3.10
# Add non root user and certs
RUN apk --no-cache add ca-certificates \
    && addgroup -S app && adduser -S -g app app \
    && mkdir -p /home/app \
    && chown app /home/app

WORKDIR /home/app

COPY --from=build /go/src/handler/handler    .
COPY --from=build /usr/bin/fwatchdog         .
COPY --from=build /go/src/handler/function/  .

RUN chown -R app /home/app

USER app

ENV fprocess="./handler"
ENV mode="http"
ENV upstream_url="http://127.0.0.1:8082"

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["./fwatchdog"]
