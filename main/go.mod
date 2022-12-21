module cognologix.com/main

go 1.19

replace cognologix.com/routers => ../routers

require cognologix.com/routers v0.0.0-00010101000000-000000000000

require (
	cognologix.com/service v0.0.0-00010101000000-000000000000 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
)

replace cognologix.com/service => ../service
