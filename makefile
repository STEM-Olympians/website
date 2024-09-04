build:
	go build server/main.go &&\
		mv main bin/main

clean:
	rm bin/main
