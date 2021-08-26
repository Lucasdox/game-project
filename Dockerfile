FROM golang:1.16-alpine as builder

WORKDIR /build

RUN apk --no-cache add git tzdata

RUN git config --global url."git@github.com:".insteadOf https://github.com/

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o game-project ./cmd/game-project/.

FROM alpine
WORKDIR /app
COPY --from=builder /build/game-project .
COPY ./data ./data
ADD build/package/docker/entrypoint.sh /

EXPOSE 8080

ENTRYPOINT ["sh", "/entrypoint.sh"]
