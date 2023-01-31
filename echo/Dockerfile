FROM golang

WORKDIR /app

COPY . ./

RUN go mod tidy
RUN go build .
EXPOSE 1323
CMD ["./echo"]