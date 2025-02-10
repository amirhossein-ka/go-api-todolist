FROM golang:1.19-alpine AS builder
ARG HTTP_PROXY
ARG HTTPS_PROXY

WORKDIR /app/

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN go build -o ./bin/app .

FROM alpine AS runner

ENV GOOS linux
ENV CGO_ENABLED 0

COPY --from=builder /app/bin/app /usr/bin/app
USER nobody:nobody
CMD [ "/usr/bin/app" ]
