FROM node:18-alpine as frontend

WORKDIR /build

COPY rvnx_doener_map/package.json ./
RUN npm install

COPY rvnx_doener_map .
RUN apk --no-cache add jq && \
    npm run fix-fa-package && \
    npm run build

# TODO: find official svelte way
RUN cp ./static/* ./build

FROM golang:1.18-alpine AS backend

WORKDIR /build

COPY rvnx_doener_service/go.mod rvnx_doener_service/go.sum ./
RUN go mod download

COPY rvnx_doener_service ./
RUN rm -f ./frontend/keep
COPY --from=frontend /build/build ./frontend

RUN go build -ldflags "-s -w" -o /backend

FROM alpine

RUN adduser -DH rvnx && mkdir /app && chown rvnx /app
WORKDIR /app

ADD CHECKS .

USER rvnx

COPY --from=backend --chown=rvnx /backend backend

ENV GIN_MODE=release

ENTRYPOINT /app/backend