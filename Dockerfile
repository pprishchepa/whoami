FROM golang:1.19 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o whoami main.go

FROM scratch
COPY --from=builder /app/whoami /whoami
ENTRYPOINT ["/whoami"]