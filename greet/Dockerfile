FROM alpine

RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY ./greetApi /app/greetApi
COPY ./etc/ /app/etc/

CMD ["./greetApi", "-f", "etc/greet-api.yaml"]