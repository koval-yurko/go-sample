module github.com/koval-yurko/go-sample

go 1.21.4

replace (
	github.com/koval-yurko/go-sample/client => ./client
	github.com/koval-yurko/go-sample/test => ./test
)

require github.com/koval-yurko/go-sample/client v0.0.4
