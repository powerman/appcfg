//go:build generate

//go:generate mkdir -p .buildcache/bin
//go:generate -command GOINSTALL env "GOBIN=$PWD/.buildcache/bin" go install

package appcfg

//go:generate GOINSTALL github.com/golangci/golangci-lint/cmd/golangci-lint@v1.48.0
//go:generate GOINSTALL github.com/mattn/goveralls@v0.0.11
//go:generate GOINSTALL github.com/cheekybits/genny@v1.0.1-0.20200709201058-3e22f1a88ff2
