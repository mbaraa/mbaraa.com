FROM golang:bookworm AS build

WORKDIR /app
COPY . .


RUN apt update && \
    apt install -y curl make &&\
    make download-tailwindcss-binary && \
    mv tailwindcss-cli /usr/bin/tailwindcss-cli

RUN ls -la /usr/bin | grep tailwindcss-cli

RUN make

FROM debian:bookworm-slim AS run

WORKDIR /app

COPY --from=build /app/mbaraacom ./mbaraacom
COPY --from=build /app/resources ./resources

EXPOSE 3000

CMD ["./mbaraacom"]
