# BUG_REPRO

The following failures were observed while validating the initial project state.
Each section records what failed, how to reproduce it, and the complete command output.
They are preserved intentionally; only failing build gates are omitted from the generated Dockerfile.

## Failure 1: Go test (.)

- Observed problem: `Go test (.)` failed in the initial project state.
- Working directory: `.`
- Command: `cd /app && GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go test -count=1 ./...`
- Exit status: `1`

```text
?   	mountainrescue/cmd/app	[no test files]
?   	mountainrescue/web	[no test files]
--- FAIL: TestSearchMissingKeywordProvidesStarterKit (0.00s)
panic: assignment to entry in nil map [recovered]
	panic: assignment to entry in nil map

goroutine 12 [running]:
testing.tRunner.func1.2({0x1e7820, 0x289ab0})
	/usr/local/go/src/testing/testing.go:1631 +0x1c4
testing.tRunner.func1()
	/usr/local/go/src/testing/testing.go:1634 +0x33c
panic({0x1e7820?, 0x289ab0?})
	/usr/local/go/src/runtime/panic.go:770 +0x124
mountainrescue/internal/rescue.(*Service).Search(0x400001a380, {0x40001141f2, 0x12})
	/app/internal/rescue/service.go:40 +0xdc
mountainrescue/internal/rescue.NewHTTPServer.func2({0x28ba08, 0x4000012e00}, 0x40?)
	/app/internal/rescue/http.go:24 +0xa4
net/http.HandlerFunc.ServeHTTP(0x4000108f70?, {0x28ba08?, 0x4000012e00?}, 0x1b6cb0?)
	/usr/local/go/src/net/http/server.go:2171 +0x38
net/http.(*ServeMux).ServeHTTP(0x22413c?, {0x28ba08, 0x4000012e00}, 0x4000000360)
	/usr/local/go/src/net/http/server.go:2688 +0x1a4
mountainrescue/internal/rescue.TestSearchMissingKeywordProvidesStarterKit(0x40001329c0)
	/app/internal/rescue/service_test.go:36 +0x120
testing.tRunner(0x40001329c0, 0x2508d8)
	/usr/local/go/src/testing/testing.go:1689 +0xec
created by testing.(*T).Run in goroutine 1
	/usr/local/go/src/testing/testing.go:1742 +0x318
FAIL	mountainrescue/internal/rescue	0.006s
FAIL
```

## Architecture reproduction

### linux/amd64
- Go toolchain version: exit `0`
- Node.js version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/app): exit `0`
- Frontend build (web): exit `0`
### linux/arm64
- Go toolchain version: exit `0`
- Node.js version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/app): exit `0`
- Frontend build (web): exit `0`
