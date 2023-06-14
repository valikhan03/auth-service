FROM golang:1.18

WORKDIR /app/
ADD ./ /app
RUN go install
RUN go build -o index

ENV HOST=127.0.0.1
ENV PORT=8840
ENV DB_PASSWORD=v.9cUrhp

ENTRYPOINT [ "/app/index" ]