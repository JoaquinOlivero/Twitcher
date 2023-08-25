FROM jrottenberg/ffmpeg:4.3-alpine AS ffmpeg
FROM golang:1.21.0-alpine3.18 AS go

ADD api/ /app/api/

RUN apk add build-base

RUN mkdir -p /app/bin

WORKDIR /app/api

RUN cp *.ttf /app/bin

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o /app/bin/Twitcher

FROM node:20.5.1-alpine3.18

COPY --from=ffmpeg / /
COPY --from=go /app/bin /app/bin
ADD src/ /app/src/
ADD proto/ /app/proto/
COPY start.sh/ /app/start.sh

RUN apk add build-base

WORKDIR /app/src

RUN npm install && npm run build

WORKDIR /

RUN chmod +x /app/bin/Twitcher
RUN chmod +x /app/start.sh

CMD /app/./start.sh