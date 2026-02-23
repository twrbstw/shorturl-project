package worker

import "context"

type Worker interface {
	Process(ctx context.Context)
}
