module client

go 1.15

require (
	csmsg v0.0.0
	queue v0.0.0
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hajimehoshi/ebiten/v2 v2.1.2
	github.com/name5566/leaf v0.0.0-20200516012428-8592b1abbbbe
	golang.org/x/image v0.0.0-20210220032944-ac19c3e999fb
)

replace csmsg => ../csmsg
replace queue => ../util/queue