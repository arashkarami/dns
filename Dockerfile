FROM golang:1.16-alpine

WORKDIR /
COPY . .
RUN go mod download
EXPOSE 5001

CMD ["go","run","main.go"]