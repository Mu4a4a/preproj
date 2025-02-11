FROM golang:1.22.4

RUN go version

# Установка рабочей директории
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN rm -rf preproj
RUN go build -o preproj ./cmd/grpc/main.go
RUN chmod +x /app/preproj


CMD ["./preproj"]