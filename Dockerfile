FROM gcr.io/distroless/static

WORKDIR /app

COPY ./build/server /app/server

ENV PORT 8080

CMD [ "/app/server" ]