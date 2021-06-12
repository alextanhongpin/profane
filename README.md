# Profane

[![Go Reference](https://pkg.go.dev/badge/github.com/alextanhongpin/profane.svg)](https://pkg.go.dev/github.com/alextanhongpin/profane)

Profane checks for profanity and also provides an option to replace them.

```go
package main

import (
	"log"

	"github.com/alextanhongpin/profane"
)

func main() {
	profane := profane.New()
	log.Println(profane.ReplaceGarbled("hello @ssh0l3"))              // hello $@!#%
	log.Println(profane.ReplaceStars("hello @ssh0l3"))                // hello a*****e
	log.Println(profane.ReplaceVowels("hello @ssh0l3"))               // hello *ssh*l*
	log.Println(profane.ReplaceCustom("hello @ssh0l3", "[CENSORED]")) // hello [CENSORED]
}
```

## Installation

```bash
$ go get github.com/alextanhongpin/profane
```

## References
- https://www.cs.cmu.edu/~biglou/resources/bad-words.txt
- https://github.com/chucknorris-io/swear-words
