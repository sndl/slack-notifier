# slack-notifier
A simple go tool to send notification from your server to slack, can be used as a replacement for notification emails

## Compile
To use the tool you have to build it providing URL to your slack webhook like this:
```
go build -ldflags "-X main.webhook=https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
```

## Usage
To send a message, simply run:
```
./slack-notifier --channel general --message "Hello World!"
```

To get more information run:
```
./slack-notifier help
```
