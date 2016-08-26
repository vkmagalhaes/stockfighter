GOVERSION=1.7
IMAGE ?= turbo
PORT ?= 8080

# FILE_ARGS = -category.file="$(PWD)/data/category.json" -artist.file="$(PWD)/data/artist.json" -brand.file="$(PWD)/data/brand.json" -seller.file="$(PWD)/data/seller.json" -redirect.file="$(PWD)/data/redirect.json"

run:
	go run main.go $(FILE_ARGS)

run_chock_a_block:
	go run cli/main.go -puzzle="chockablock" -api-key="$(STOCKFIGHTER_KEY)" -account="$(STOCKFIGHTER_ACCOUNT)" -venue="IBLEX"  -symbol="CHYU" -shares.per-trade=100 -shares.goal=100000 -price.min=100 -price.drop=50 -price.start=5000

# build:
# 	go build -v -o dist/slug
#
# test:
# 	go test -cover -race $$(go list ./... | grep -v /vendor/)
#
# bench:
# 	go test -cover -benchmem -race -bench=. $$(go list ./... | grep -v /vendor/)

clean:
	rm -rf main dist/
