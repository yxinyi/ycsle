module server

go 1.15

require (
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/name5566/leaf v0.0.0-20200516012428-8592b1abbbbe
	csmsg v0.0.0-00010101000000-000000000000
)

replace csmsg => ../csmsg
