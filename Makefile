run: build
	go run main.go watch_files.go  \
		classify.go  text_to_speech.go \
		slack.go

copy_image_file:
	rm -f n02085620_1205.jpg 
	cp /Users/vinodhalaharvi/Downloads/images/only20/n02085620_1205.jpg  .

build: 
	go build *.go
