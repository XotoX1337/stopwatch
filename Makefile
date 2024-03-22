check:
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	gocyclo -over 10 .
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	gosec -quiet --severity high ./...
	go install github.com/sonatype-nexus-community/nancy@latest
	go list -json -deps ./... | nancy sleuth