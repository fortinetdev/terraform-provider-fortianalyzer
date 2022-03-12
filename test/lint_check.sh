p=$(dirname "$PWD");
export GOPATH=${p%/*/*/*/*}"/"
make -C ../  fmt
golint ../faz
golint ../vendor/github.com/fortinetdev/forti-sdk-go/fortianalyzer/auth
golint ../vendor/github.com/fortinetdev/forti-sdk-go/fortianalyzer/config
golint ../vendor/github.com/fortinetdev/forti-sdk-go/fortianalyzer/request
golint ../vendor/github.com/fortinetdev/forti-sdk-go/fortianalyzer/sdkcore
make -C ../  build

