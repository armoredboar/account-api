# This Dockerfile is executed in 2 steps. The first step is a build step.
# An image is created and the code is compiled. To compile the application,
# several extra files are required, like golang compiler, git, some kernel
# modules and more. This produces an image with a size of ~500MB, which is
# a little too much for this app.
# The solution is a second build stage, named 'Final Stage' here. This
# stage creates a new image and copies the builded app from the first image.
# Since we do not need any extra components or modules to run or build our
# application, because it is already compiled to alpine OS, the final image
# size is something less than 20MB. A huge diference when uploading the
# image and saving registry disk storage. 

# Build Stage
FROM golang:alpine as builder

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/armoredboar/account-api/

# Build the project inside the container.
# Install git that is required by 'go get' command that follows.
RUN apk add --no-cache git
# The account-api project depends on gin-gonic webserver framework to run.
RUN go get -d -v github.com/gin-gonic/gin
# Installs the account-api project. Go default output is /go/bin/<ProjectName>
# The -i flag indicates Go to install the dependencies as well, like gin-gonic.
RUN go install -i github.com/armoredboar/account-api/cmd/account-api

# Final Stage
FROM alpine:latest

# Add the required environment variables to the image.
# The variables will be set by docker-compose.
ENV AB_EMAIL_SMTP=${AB_EMAIL_SMTP}
ENV AB_EMAIL_PORT=${AB_EMAIL_PORT}
ENV AB_EMAIL_NOREPLY=${AB_EMAIL_NOREPLY}
ENV AB_EMAIL_NOREPLY_NAME=${AB_EMAIL_NOREPLY_NAME}
ENV AB_EMAIL_NOREPLY_PASSWORD=${AB_EMAIL_NOREPLY_PASSWORD}

# Copy the compiled account-api executable from the previous image.
# The executable account-api is renamed to app.
COPY --from=builder /go/bin/account-api /app
# Set the running command when this image is executed. In this case, execute our app.
ENTRYPOINT ./app
# Create a label to this image.
LABEL Name=account-api Version=0.0.1
# Document that the service listens on port 3500.
EXPOSE 3500