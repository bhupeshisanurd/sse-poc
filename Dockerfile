FROM golang
WORKDIR /src
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o server server/server.go

CMD ["./server/server"]
