.PHONY: build

BINARY_NAME=mbaraacom

build:
	npm install tailwindcss @tailwindcss/cli && \
	npx @tailwindcss/cli -i ./resources/css/style.css -o ./resources/css/tailwind.css -m && \
	go mod tidy && \
	go build -ldflags="-w -s" -o ${BINARY_NAME}

dev:
	npx @tailwindcss/cli -i ./resources/css/style.css -o ./resources/css/tailwind.css -m && \
	go build -o ${BINARY_NAME}
	./${BINARY_NAME}

clean:
	go clean
