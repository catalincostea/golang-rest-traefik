# FROM golang:1.21
FROM golang:1.21-alpine as builder

RUN apk add --no-cache \
    curl \
    bash \
    unzip && \
    curl -fsSL https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_linux_amd64.zip -o terraform.zip && \
    unzip terraform.zip && \
    mv terraform /usr/local/bin/ && \
    rm terraform.zip

WORKDIR /app/api
COPY . .
RUN go mod tidy
RUN go build -o main .

FROM alpine:latest
RUN apk add --no-cache bash

# Copy the built Go binary and Terraform from the builder stage
COPY --from=builder /usr/local/bin/terraform /usr/local/bin/terraform
COPY --from=builder /app/api/main /usr/local/bin/main

# For terraform dev purposes only
COPY --from=builder /app/ /app

EXPOSE 3000
WORKDIR /app/api

# CMD ["tail", "-f", "/dev/null"]
CMD ["./main"]
