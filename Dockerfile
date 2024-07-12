FROM golang:1.22.4-alpine

WORKDIR /journey

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .
COPY .env /journey/bin/journey


WORKDIR /journey/cmd/journey

RUN go build -o /journey/bin/journey .

# Adicione um comando para verificar se o binário está presente
RUN ls -l /journey/bin



EXPOSE 8080
RUN ls -l /journey
ENTRYPOINT [ "/journey/bin/journey" ]
