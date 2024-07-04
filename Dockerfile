# Imposta l'immagine base ufficiale di Go come ambiente di costruzione
FROM golang:1.21.6 AS builder

# Imposta il working directory all'interno del container
WORKDIR /app

# Copia i file go.mod e go.sum per gestire le dipendenze
COPY go.mod go.sum ./
RUN go mod download

# Copia il codice sorgente nella directory corrente del container
COPY . .

# Compila l'applicazione in un eseguibile standalone
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .
#RUN go get -u github.com/go-delve/delve/cmd/dlv
# Fase di runtime: Usa un'immagine Docker leggera
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copia l'eseguibile dallo stage di build
COPY --from=builder /app/server .

# Espone la porta 8080
#EXPOSE 8884
#EXPOSE 2345

# Comando eseguito all'avvio del container
CMD ["./server"]
#CMD ["dlv", "debug", "--headless", "--listen=:2345", "--api-version=2", "--accept-multiclient"]