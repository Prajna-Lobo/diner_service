FROM golang:1.19

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /diner_service app/main.go

EXPOSE 8080

#CMD [ "/diner_service" ]
ENTRYPOINT [ "/diner_service", "-fpath=./app/log.csv"]