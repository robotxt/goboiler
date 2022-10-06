FROM golang:1.18

WORKDIR /application

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go install github.com/githubnemo/CompileDaemon@latest
ENTRYPOINT CompileDaemon --build="go build -o executable" --command=./executable

