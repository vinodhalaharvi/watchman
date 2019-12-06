run: build
	go run main.go watch_files.go  \
		classify.go  text_to_speech.go \
		slack.go

copy_image_file:
	rm -f n02085620_1205.jpg 
	wget https://cdn.pixabay.com/photo/2016/12/13/05/15/puppy-1903313_960_720.jpg


build: 
	go build *.go
