# build
FROM golang:1.16 as builder

WORKDIR /workspace

COPY src/go.mod go.mod
COPY src/main.go main.go

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o main main.go

# run
FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /workspace/main .

ENTRYPOINT ["/main"]
