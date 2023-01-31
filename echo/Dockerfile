FROM golang

WORKDIR /app

COPY . ./

RUN go mod tidy
RUN go build .
CMD ["./echo"]