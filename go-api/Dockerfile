# For production
FROM golang:1.18 as production

RUN mkdir /api
WORKDIR /api

ENV ENV=production

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v

EXPOSE 8080
CMD ["./main"]

# For development
FROM golang:1.18 as development

RUN mkdir /api
WORKDIR /api

ENV ENV=development

COPY go.mod go.sum ./
RUN go mod tidy

RUN go install github.com/cosmtrek/air@v1.27.3
CMD ["air", "-c", ".air.toml"]
