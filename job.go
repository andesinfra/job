package job

import "context"

type Runnable interface {
	Run(ctx context.Context) error
}

func Schedule(x Runnable, cron string, onErr func(error)) error {
	x.Run(context.Background())
	return nil
}
