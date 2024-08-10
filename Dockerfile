# Fase de construcción
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copia los archivos necesarios para la compilación
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Compila la aplicación
RUN GOOS=linux GOARCH=amd64 go build -o bootstrap main.go

# Fase final: crea la imagen para AWS Lambda
FROM amazonlinux:2

WORKDIR /var/task

# Copia el archivo binario compilado y el archivo .env desde la fase de construcción
COPY --from=builder /app/bootstrap .
COPY --from=builder /app/.env .

# Establece el punto de entrada para Lambda
ENTRYPOINT ["/var/task/bootstrap"]
