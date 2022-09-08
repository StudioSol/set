dependency:
	go mod vendor

test:
	echo "" > coverage.txt
	for d in $(shell go list ./... | grep -v vendor); do \
		go test -mod=mod -race -v -coverprofile=profile.out -covermode=atomic $$d || exit 1; \
		[ -f profile.out ] && cat profile.out >> coverage.txt && rm profile.out; \
	done
