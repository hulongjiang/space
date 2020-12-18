FROM alpine

RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY ./userRpc /app/userRpc
COPY ./etc/user.yaml /app/etc/config.yaml

CMD ["./userRpc", "-f", "etc/config.yaml"]
