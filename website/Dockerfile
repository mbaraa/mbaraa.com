FROM golang:1.21-alpine as build

WORKDIR /app
COPY . .

RUN apk add make npm nodejs
RUN make build_website

FROM alpine:latest as run

WORKDIR /app

COPY --from=build /app/website/mbaraacom-website ./run
COPY --from=build /app/website/resources ./resources

EXPOSE 3000

CMD ["./run"]
