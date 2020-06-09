FROM golang:1.8.1-alpine

RUN go mod tidy

ENV SOURCE ${pwd}
COPY . ${SOURCE}

RUN cd ${SOURCE} && CGP_ENABLED=0 go build

WORKDIR ${SOURCE}
CMD ${SOURCE}Gin-Web

EXPOSE 8091
