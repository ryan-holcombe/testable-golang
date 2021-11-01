FROM golang:1.17.1 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make build BIN=server

FROM alpine:3.14

ARG GIT_COMMIT
ARG VERSION
ARG BUILD_DATE
ARG APP

ENV APP=$APP
ENV GIT_COMMIT=$GIT_COMMIT
ENV VERSION=$VERSION
ENV BUILD_DATE=$BUILD_DATE

LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION
LABEL BUILD_DATE=$BUILD_DATE

WORKDIR /

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/bin/server .
ENTRYPOINT ["/server"]
