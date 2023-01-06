# gopherlogs

_A simple, powerful, and extensible Go logging framework with batteries included._

`gopherlogs` is ideal for command line applications,
combined logging to system files alongside terminal output,
and much more.

Features:
- Animated logging compatible with concurrent Go routines
- Support for "emoji" style logs
- Dynamic log line replacement
- No external dependencies; what you see is what you get.

## Usage

Install as a dependency in your Go module:

```
module example.com

go 1.19

require github.com/jpmcb/gopherlogs v0.1.0
```

And tidy up your modules and `go.sum` with:

```
$ go mod tidy
```

In your go code, you can now import gopherlogs and use it!

```go
package main

import (
        "github.com/jpmcb/gopherlogs"
)

func main() {
        l := gopherlogs.NewLogger()
        l.Info("Hello world")
}
```

When you run your Go program, it will look something like this:

```
$ go run main.go
Hello world
```

## Why another Go logging framework?

This library is heavily inspired by the original logging library
from [Tanzu Community Edition's](https://github.com/vmware-tanzu/community-edition)
`unmanaged-cluster` CLI (which is no longer being maintained and is effectively abandoned).

I worked on that project for a few years, and we created the first iterations of this
logging framework because we couldn't find anything suitable enough
that was still very delightful to experience.

That experience is at the heart of this library; an amazing user experience
for both the _end user_ and the _developer_.
In some ways, this logging library is a fork of some of the best pieces of Go code
that came out of that project.
And in the end, I believed it deserved to see abit more lite of day.

Shout out and huge kudos to the original authors: @stmcginnis, @joshrosso, and @jpmcb (me).
Your work has inspired me and continues to bring me joy!

