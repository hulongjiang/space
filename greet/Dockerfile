FROM alpine

ENV TZ Asia/Shanghai

WORKDIR /app
COPY ./greetApi /app/greetApi
COPY ./etc/ /app/etc/

CMD ["./greetApi", "-f", "etc/greet-api.yaml"]