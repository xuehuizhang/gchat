module com.gchat/connect

go 1.18

replace com.gchat.pkg v0.0.1 => ../pkg

require (
	com.gchat.pkg v0.0.1
	github.com/gorilla/websocket v1.5.0
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
)
