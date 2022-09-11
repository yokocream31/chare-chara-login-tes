# For development
FROM golang:1.19

RUN mkdir /api
WORKDIR /api

ENV ENV=development

COPY go.mod go.sum ./
RUN go mod tidy

RUN go install github.com/cosmtrek/air@v1.27.3
CMD ["air", "-c", ".air.toml"]
