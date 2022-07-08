FROM gcr.io/distroless/static

COPY ./build/server /app/server

ENV PORT 443

CMD [ "/app/server" ]