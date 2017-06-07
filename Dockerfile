# Base image
FROM golang:alpine
MAINTAINER platel.kevin@gmail.com

# Define app name
# Use run-docker.sh rather than building directly the Dockerfile
ENV APP_NAME fizzbuzz-api

# Sepcify the current directory to use. '.' will now refer to that directory
WORKDIR /app

# Copy only the binary, so  should have been build before
COPY fizzbuzz-api /app/

# Explicite which port the container expose
# Note: this not insure the port will be bind to the local machine when running,
#       you should use the '-p' option or specify the port in the docker-compose.yml file
EXPOSE 8080

# Command to execute by the container
ENTRYPOINT /app/fizzbuzz-api
