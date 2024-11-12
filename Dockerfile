FROM golang:1.20-alpine as Build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main 

FROM amd64/golang:1.20-alpine as Final

COPY --from=Build /app/main main

RUN chmod +x ./main

ENTRYPOINT ["./main"]