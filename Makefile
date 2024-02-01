BINARY_NAME=mbaraacom

build_tailwindcss:
	cd tailwindcss && \
	npx tailwindcss -i ../website/resources/css/style.css -o ../website/resources/css/tailwind.css -m && \
	npx tailwindcss -i ../dashboard/resources/css/style.css -o ../dashboard/resources/css/tailwind.css -m && \
	cd ..

build_dashboard: build_tailwindcss
	cd ./dashboard && \
	go mod tidy && \
	go build -ldflags="-w -s" -o ${BINARY_NAME}-dashboard

build_website: build_tailwindcss
	cd ./website && \
	go mod tidy && \
	go build -ldflags="-w -s" -o ${BINARY_NAME}-website

# install inotify-tools
dev_website:
	cd ./website && \
	export `xargs < .env` && \
	cd tailwindcss && \
	npx tailwindcss -i ../resources/css/style.css -o ../resources/css/tailwind.css --watch & \
	while true; do \
	  go build -o ${BINARY_NAME}; \
	  ./${BINARY_NAME} & \
	  PID=$$!; \
	  echo "PID=$$PID"; \
	  inotifywait -r -e modify ./**/*; \
	  kill $$PID; \
	done

# install inotify-tools
dev_dashboard:
	cd ./dashboard && \
	export `xargs < .env` && \
	cd tailwindcss && \
	npx tailwindcss -i ../resources/css/style.css -o ../resources/css/tailwind.css --watch & \
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
