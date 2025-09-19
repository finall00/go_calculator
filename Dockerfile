FROM golang:1.24-alpine AS builder


WORKDIR  /app

# COPY go.mod go.sum ./
COPY go.mod  ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /calculator .

FROM scratch


WORKDIR /

COPY --from=builder /calculator /calculator

ENTRYPOINT [ "/calculator" ]
