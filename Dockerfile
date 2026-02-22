FROM golang:1.25-alpine AS builder
WORKDIR /app

# Copier les fichiers de dépendances
COPY go.mod ./
# (On copiera go.sum plus tard quand il sera généré)

# Copier le code source
COPY . .

# Compiler l'application
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# Image finale
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copier le binaire
COPY --from=builder /app/main .
# Copier le dossier public contenant ton frontend HTML
COPY --from=builder /app/public ./public

EXPOSE 3000
CMD ["./main"]