FROM alpine

ENV TZ Asia/Shanghai

WORKDIR /app
COPY ./api /app/api
COPY ./etc/ /app/etc/

CMD ["./api", "-f", "etc/api-api.yaml"]