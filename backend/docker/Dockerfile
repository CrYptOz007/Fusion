# Official Golang image
FROM golang:latest

ENV GO111MODULE=on
ENV ENVIRONMENT=local

# Working directory
WORKDIR /var/www/app/backend

RUN go install github.com/beego/bee/v2@latest
RUN apt-get update && apt-get install -y ipmitool

EXPOSE 4000

CMD ["bee", "run"]
