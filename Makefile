gobuild:
	go build -a -installsuffix cgo -o main .

run:
	gin --appPort 8080 --port 5000 -i

db:
	goose postgres "user=larrya dbname=iv-images sslmode=disable" redo

dkbuild:
	docker build -t larrya/images .