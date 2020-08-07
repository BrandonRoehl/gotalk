.PHONY: test

test:
	go test -benchmem -bench .

serve:
	godoc -http=:6060
