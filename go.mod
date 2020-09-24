module github.com/go-clang/gen

go 1.13

require (
	github.com/go-clang/clang-v3.9 v0.0.0-20190823090603-8e83bb44d7e2
	github.com/stretchr/testify v1.6.1
	github.com/termie/go-shutil v0.0.0-20140729215957-bcacb06fecae
	golang.org/x/tools v0.0.0-20200922173257-82fe25c37531
)

replace github.com/go-clang/clang-v3.9 => ../clang-v3.9
