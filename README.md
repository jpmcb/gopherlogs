<p align="center">
    <img src="https://user-images.githubusercontent.com/23109390/227600598-11e1f583-fca2-4f9e-9626-b26dbdbc1323.png" width=300>
    <img src="https://user-images.githubusercontent.com/23109390/227600551-58c7c47b-3407-4263-8c48-668e5a8743c6.gif" width=450>
</p>

# gopherlogs

[![go report card](https://goreportcard.com/badge/github.com/jpmcb/gopherlogs "go report card")](https://goreportcard.com/report/github.com/jpmcb/gopherlogs)
[![test status](https://github.com/jpmcb/gopherlogs/workflows/Tests/badge.svg?branch=main)](https://github.com/jpmcb/gopherlogs/actions)
[![Apache-2.0 license](https://img.shields.io/github/license/jpmcb/gopherlogs)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/jpmcb/gopherlogs)

_A simple, powerful, and extensible Go logging framework with batteries included 🔋._

`gopherlogs` is ideal for command line applications,
combined logging to system files alongside terminal output,
and much more.

Features:
- Animated logging compatible with concurrent Go routines
- Support for _"emoji"_ style logs
- Dynamic log line replacement
- No external dependencies; what you see is what you get.

---

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

In your go code, you can now import `gopherlogs`:

```go
package main

import (
        "github.com/jpmcb/gopherlogs"
)

func main() {
        // Creates a new logger with options
        l, err := gopherlogs.NewLogger(
            WithLogVerbosity(5),
        )

        // Handle errors from creating a new logger
        if err != nil {
            panic("Could not create new logger")
        }

        // Start logging!
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

