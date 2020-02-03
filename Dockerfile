FROM golang

WORKDIR /app
COPY ./go-time /app

CMD ["/app/go-time"]