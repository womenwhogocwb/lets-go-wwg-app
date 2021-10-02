FROM golang:latest
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o build/api cmd/main.go
EXPOSE 3000
CMD ["/app/build/api"]
