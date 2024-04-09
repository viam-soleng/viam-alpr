
bin/viamalpr: *.go cmd/module/*.go go.*
	go build -o bin/viam-alpr.AppImage cmd/module/cmd.go

bin/remoteserver: *.go cmd/remote/*.go go.*
	go build -o bin/remoteserver cmd/remote/cmd.go

lint:
	gofmt -w -s .

updaterdk:
	go get go.viam.com/rdk@latest
	go mod tidy

module: bin/viamalpr
	tar czf module.tar.gz bin/*