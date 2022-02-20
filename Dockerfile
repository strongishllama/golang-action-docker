FROM golang:1.17 as builder

WORKDIR /action
COPY . /action

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o action main.go

FROM gcr.io/distroless/static

COPY --from=builder /action/action /action

ENTRYPOINT ["/action"]
