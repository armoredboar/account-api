#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
ENV AB_EMAIL_SMTP=${AB_EMAIL_SMTP}
ENV AB_EMAIL_PORT=${AB_EMAIL_PORT}
ENV AB_EMAIL_NOREPLY=${AB_EMAIL_NOREPLY}
ENV AB_EMAIL_NOREPLY_NAME=${AB_EMAIL_NOREPLY_NAME}
ENV AB_EMAIL_NOREPLY_PASSWORD=${AB_EMAIL_NOREPLY_PASSWORD}
COPY --from=builder /go/bin/app /app
ENTRYPOINT ./app
LABEL Name=account-api Version=0.0.1
EXPOSE 3000
