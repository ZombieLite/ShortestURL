FROM golang:latest

COPY ./ ./
RUN go build -o main .
RUN chmod +x main
CMD ("./main")
