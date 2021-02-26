gobuild:
	go build -a -installsuffix cgo -o main .

run:
	gin --appPort 8080 --port 5000 -i

mstatus:
	goose status

dkbuild:
	docker build -t larrya/images .