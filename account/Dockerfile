FROM golang:latest as builder

#Définit le répertoire de travail pour le builder.
WORKDIR /go/src/app

ENV GO111MODULE=on

# Get Reflex for live reload in dev
#RUN go install github.com/cespare/reflex@latest
#RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/cosmtrek/air@latest


COPY go.mod .
COPY go.sum .

RUN go mod download
 
COPY . .

RUN go build -o ./run  cmd/main.go

#une image Alpine Linux minimale est utilisée comme image de base pour l'image finale.
FROM alpine:latest

#Installe les certificats CA pour la gestion des connexions HTTPS.
RUN apk --no-cache add ca-certificates
WORKDIR /root/

#Copy executable from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8080

#CMD ["./run"]
 
