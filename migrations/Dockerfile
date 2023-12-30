# FROM alpine:3.17.0

# RUN apk update && \
#     apk upgrade && \
#     apk add bash && \
#     rm -rf /var/cache/apk/*

FROM golang:latest

ARG ENV=development

WORKDIR /migration

# ADD https://github.com/pressly/goose/releases/download/v3.7.0/goose_linux_x86_64 /bin/goose

COPY . .
# COPY migration-main.go main.go
COPY ".env.$ENV" .env
# COPY cmd/main.go main.go

RUN go build -o bin main.go

ENTRYPOINT ["./bin"]

# CMD [ "sh", "-c", "ls migrations/ && ./mig" ]

# ENTRYPOINT ["/bin/goose"]	
