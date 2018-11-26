# build stage
FROM golang:alpine as builder

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/armoredboar/account-api/

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN apk add --no-cache git
RUN go get -d -v github.com/gin-gonic/gin
RUN go install -i github.com/armoredboar/account-api/cmd/account-api

#ENTRYPOINT /go/bin/account-api

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
# Run the outyet command by default when the container starts.

ENV AB_EMAIL_SMTP=${AB_EMAIL_SMTP}
ENV AB_EMAIL_PORT=${AB_EMAIL_PORT}
ENV AB_EMAIL_NOREPLY=${AB_EMAIL_NOREPLY}
ENV AB_EMAIL_NOREPLY_NAME=${AB_EMAIL_NOREPLY_NAME}
ENV AB_EMAIL_NOREPLY_PASSWORD=${AB_EMAIL_NOREPLY_PASSWORD}

COPY --from=builder /go/bin/account-api /app
ENTRYPOINT ./app
LABEL Name=account-api Version=0.0.1
# Document that the service listens on port 3500.
EXPOSE 3500