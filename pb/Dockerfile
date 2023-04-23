FROM alpine:latest

ARG PB_VERSION=0.12.1

RUN apk add --no-cache \
    unzip \
    ca-certificates

# download and unzip PocketBase
COPY . ./

EXPOSE 8090

# start PocketBase
CMD ["./pocketbase", "serve", "--http=0.0.0.0:8090"]
