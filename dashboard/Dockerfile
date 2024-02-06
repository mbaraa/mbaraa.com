FROM golang:1.21-alpine as build

WORKDIR /app
COPY . .

RUN apk add make npm nodejs
RUN make build_dashboard

FROM alpine:latest as run

WORKDIR /app

COPY --from=build /app/dashboard/mbaraacom-dashboard ./run
COPY --from=build /app/dashboard/resources ./resources

EXPOSE 3000

CMD ["./run"]
