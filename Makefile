all: test



test:
	go get github.com/onsi/gomega
	go get github.com/onsi/ginkgo/ginkgo
	ginkgo -r
