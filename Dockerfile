FROM golang:alpine AS build

WORKDIR /app
COPY . .


RUN apk add --no-cache npm nodejs make

RUN make

FROM alpine:latest AS run

WORKDIR /app

COPY --from=build /app/mbaraacom ./mbaraacom
COPY --from=build /app/resources ./resources

EXPOSE 3000

CMD ["./mbaraacom"]
