
bin/viamalpr: *.go cmd/module/*.go go.*
	go build -o bin/viam-alpr.AppImage cmd/module/cmd.go
	dylibbundler -od -b -s /usr/local/lib -x ./bin/viam-alpr.AppImage
	# install_name_tool -add_rpath @executable_path/../libs bin/viam-alpr.AppImage -> for development faster as it doesn't collect all the libs but simply updates the path

dev: *.go cmd/module/*.go go.*
	go build -o bin/viam-alpr.AppImage cmd/module/cmd.go
	install_name_tool -add_rpath @executable_path/../libs bin/viam-alpr.AppImage

bin/remoteserver: *.go cmd/remote/*.go go.*
	go build -o bin/remoteserver cmd/remote/cmd.go

lint:
	gofmt -w -s .

updaterdk:
	go get go.viam.com/rdk@latest
	go mod tidy

module: bin/viamalpr
	tar czf module.tar.gz bin libs openalpr