package part

import (
	"path/filepath"
	"runtime"
	"time"
)

var (
	LogFlushWait = 200 * time.Millisecond
)

var (
	_, b, _, _ = runtime.Caller(0)
	RootPath   = filepath.Join(filepath.Dir(b), "/")
)
