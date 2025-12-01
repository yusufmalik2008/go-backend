<<<<<<< HEAD
.PHONY: templ tailwind dev test

templ:
	templ generate

tailwind:
	go run scripts/build-tailwind.go

dev: templ tailwind
	air   # hot reload pake cosmtrek/air

test:
	go test -v ./...

metrics:
	open http://localhost:9090
=======
.PHONY: templ tailwind dev test

templ:
	templ generate

tailwind:
	go run scripts/build-tailwind.go

dev: templ tailwind
	air   # hot reload pake cosmtrek/air

test:
	go test -v ./...

metrics:
	open http://localhost:9090
>>>>>>> d70ebf2dc433eb0293c604a2c53a9f3bed25aca9
