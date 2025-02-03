FROM golang:1.21-alpine as build

WORKDIR /app
COPY . .

RUN apk add --no-cache make npm nodejs &&\
    make

FROM alpine:latest as run

WORKDIR /app

COPY --from=build /app/mbaraacom ./mbaraacom
COPY --from=build /app/resources ./resources

EXPOSE 3000

CMD ["./mbaraacom"]
