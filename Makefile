.PHONY: build

BINARY_NAME=mbaraacom

build:
	cd tailwindcss && \
	npx @tailwindcss/cli -i ../resources/css/style.css -o ../resources/css/tailwind.css -m && \
	cd .. && \
	go mod tidy && \
	go build -ldflags="-w -s" -o ${BINARY_NAME}


# install inotify-tools
dev:
	cd tailwindcss && \
	npx @tailwindcss/cli -i ../resources/css/style.css -o ../resources/css/tailwind.css --watch & \
	while true; do \
	  go build -o ${BINARY_NAME}; \
	  ./${BINARY_NAME} & \
	  PID=$$!; \
	  echo "PID=$$PID"; \
	  inotifywait -r -e modify ./**/*; \
	  kill $$PID; \
	done

clean:
	go clean
