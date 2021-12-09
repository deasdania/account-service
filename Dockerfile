FROM golang:alpine

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git

WORKDIR /app/auth-service

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . . 

# Install the air binary so we get live code-reloading when we save files
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

EXPOSE 8080

CMD ["air"]

CMD ["go","run","/app/auth-service/main.go"]
