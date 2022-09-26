FROM golang:1.18

WORKDIR /app
ADD /. /app


RUN go install

RUN go build -o index



ENTRYPOINT [ "/app/index" ]