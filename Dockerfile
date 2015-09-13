FROM alpine:3.2

EXPOSE 55555

ADD artifacts/iris /app/iris
ENTRYPOINT [ "/app/iris" ]
