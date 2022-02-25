FROM golang:1.17 AS builder

COPY ${PWD} /app
WORKDIR /app


RUN CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o /app/appbin *.go


FROM scratch
LABEL MAINTAINER Daniel Pickens <daniel@gmail.com>

COPY --from=builder /app /app

WORKDIR /app

EXPOSE 8000
EXPOSE 8443

CMD ["./appbin"]
