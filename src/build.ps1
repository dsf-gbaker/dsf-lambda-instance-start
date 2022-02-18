$env:GOOS="linux"
$env:CGO_ENABLED="0"
$env:GOARCH="amd64"
go build -o ../dist/main main.go
cd ../dist
build-lambda-zip -output main.zip main