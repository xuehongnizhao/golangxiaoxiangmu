package log

import (
	"dashboard/context"
	"testing"
)

func Test_Start(t *testing.T) {
	ctx := context.InitContext()
	InitLog(ctx)
}
