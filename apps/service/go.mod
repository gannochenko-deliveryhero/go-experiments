module github.com/gannochenko-deliveryhero/go-experiments/apps/service/cmd/main.go

go 1.18

replace github.com/gannochenko-deliveryhero/go-experiments/apps/service/internal/util v0.0.0 => ./internal/util

require (
	github.com/gannochenko-deliveryhero/go-experiments/apps/service/internal/util v0.0.0
)
