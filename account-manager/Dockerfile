#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=accountmanager Version=0.0.1
EXPOSE 3001

# Gunakan image Golang sebagai basis
FROM golang:1.18-alpine

# Setel direktori kerja dalam container
WORKDIR /app

# Salin file modul Go ke dalam container
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Salin seluruh kode sumber ke dalam container
COPY . .

# Bangun aplikasi
RUN go build -o /account-manager

# # Tentukan port yang akan diekspos
# EXPOSE 3001

# Tentukan perintah untuk menjalankan aplikasi
CMD [ "/account-manager" ]
