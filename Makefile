.PHONY: build

BINARY_NAME=mbaraacom

build:
	tailwindcss-cli -i ./resources/css/style.css -o ./resources/css/tailwind.css -m && \
	go mod tidy && \
	go build -ldflags="-w -s" -o ${BINARY_NAME}


# install inotify-tools
dev:
	tailwindcss-cli -i ./resources/css/style.css -o ./resources/css/tailwind.css -m && \
	while true; do \
	  go build -o ${BINARY_NAME}; \
	  ./${BINARY_NAME} & \
	  PID=$$!; \
	  echo "PID=$$PID"; \
	  inotifywait -r -e modify ./**/*; \
	  kill $$PID; \
	done

download-tailwindcss-binary:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v4.0.5/tailwindcss-linux-x64
	chmod +x tailwindcss-linux-x64
	mv tailwindcss-linux-x64 tailwindcss-cli

clean:
	go clean
