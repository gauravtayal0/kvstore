# kvstore
Key value in memory store in golang

# Usage

```go
package main

import (
	"github.com/gauravtayal0/kvstore"
)

func main() {
	kvs := kvstore.New()

	#set function usage
	kvs.Set("test", "1")

	#get function usage
	kv, err := kvs.Get("test")

	#getValue function usage
	value, err := kvs.GetValue("test")

	#delete key function usage
	kvs.Delete("test")
}
```