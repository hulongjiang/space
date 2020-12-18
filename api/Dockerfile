FROM alpine

ENV TZ Asia/Shanghai

WORKDIR /app
COPY ./api /app/api
COPY ./etc/ /app/etc/
RUN chmod +x api

CMD ["./api", "-f", "etc/api-api.yaml"]