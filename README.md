# watchman
AI watchman that watches for events and notifies you

## Run 
```make run```

## Now what? 
Copy file from web in the current directory and the watchman will notify you ```wget https://cdn.pixabay.com/photo/2016/12/13/05/15/puppy-1903313_960_720.jpg``` or if you a cat person ```wget https://cdn.pixabay.com/photo/2017/02/20/18/03/cat-2083492_960_720.jpg```.

If you want slack notification, then change these in slack.go
token       = "YOUR_SLACK_TOKEN"
channelName = "YOUR_SLACK_CHANNEL"

