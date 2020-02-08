FROM golang:1.13-alpine as builder
WORKDIR /root
COPY ./  ./

RUN export GO111MODULE=on && CGO_ENABLED=0 GOOS=linux go build -o build/gpserver server/main.go

FROM node:12-alpine as node-builder
WORKDIR /root
COPY web .
RUN npm rebuild node-sass && npm install && npm run build:prod


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /root/build/gpserver ./
COPY --from=node-builder /root/dist dist

#EXPOSE 8080
ENTRYPOINT ["/root/gpserver"]
