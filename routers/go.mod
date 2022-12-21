module cognologix.com/routers

go 1.19

replace cognologix.com/books => ../books

require (
	cognologix.com/service v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

replace cognologix.com/service => ../service
