### Go Module Tracing
Tracing modules

### Installation
```sh
go get github.com/fajarardiyanto/flt-go-tracer
```

###### Upgrading to the latest version
```sh
go get -u github.com/fajarardiyanto/flt-go-tracer
```

###### Upgrade or downgrade with tag version if available
```sh
go get -u github.com/fajarardiyanto/flt-go-tracer@v1.0.0
```

### Usage
```go
package main

import (
	"context"
	"github.com/fajarardiyanto/flt-go-tracer/interfaces"
	"github.com/fajarardiyanto/flt-go-tracer/lib"
	jaeger2 "github.com/fajarardiyanto/flt-go-tracer/lib/jaeger"
)

func main() {
	jaeger := lib.NewLib()
	_, closer := jaeger.LoadJaeger("Module", interfaces.JaegerConfig{
		Endpoint: "6831",
		LogSpan:  true,
	}).InitTracer()

	defer closer.Close()

	span, _ := jaeger2.CreateRootSpan(context.Background(), "TESTING")
	defer span.Finish()

	sp := jaeger2.CreateSubChildSpan(span, "HELLO")
	defer sp.Finish()

	jaeger2.LogRequest(sp, "Example Module")
}


```

#### Run Example
```sh
make help
```

#### Tips
Maybe it would be better to do some basic code scanning before pushing to the repository.
```sh
# for *.nix users just run gosec.sh
# curl is required
# more information https://github.com/securego/gosec
make scan
```