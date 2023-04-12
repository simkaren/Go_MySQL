FROM golang:1.18.5

WORKDIR /opt/go_mysql
COPY . .
RUN go build -o a src/main.go
CMD /opt/go_mysql/a