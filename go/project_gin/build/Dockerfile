FROM ccr.ccs.tencentyun.com/grepcode/ubuntu:20.04 AS builder
ENV WORKDIR=/go/src/grepcode.cn/app/
WORKDIR ${WORKDIR}
COPY .  ${WORKDIR}
RUN mkdir -p ./bin && make compile

FROM ccr.ccs.tencentyun.com/grepcode/debian:stable AS runtime

LABEL maintainer="houming818@qq.com"

WORKDIR /app

RUN mkdir -p /data/app && mkdir -p /data/log

COPY --from=builder /go/src/grepcode.cn/app/target /app/bin
COPY ./etc /app/etc
ENTRYPOINT ["./bin/main", "-config", "/app/etc/config.yaml"]
