FROM golang

WORKDIR /app
COPY ./go-time /app
COPY ./entrypoint.sh /app
RUN chmod 777 /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]