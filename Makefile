build:
	./tailwindcss -i publuc/input.css -o public/output.css --minify
	go build -o app 

run: build
	./app
