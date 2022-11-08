FROM golang:1.19

WORKDIR /
COPY . .
RUN go mod download
EXPOSE 5001

CMD ["go","run","main.go"]