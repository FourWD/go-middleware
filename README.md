SETUP GOENV For Github
go env -w GO111MODULE=on
go env -w GOPRIVATE=github.com/HinekoTech/go-middleware

go get -u github.com/HinekoTech/go-middleware
