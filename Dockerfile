FROM golang:1.22-alpine3.18
WORKDIR /app
COPY . .
RUN go build -o main main.go

EXPOSE 8080
CMD [ "/app/main" ]