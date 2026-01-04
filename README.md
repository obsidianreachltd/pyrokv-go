# PyroKV Go Library

## Overview

This library provides client access to a PyroKV key-value database.

## Installing

```bash
go get github.com/obsidianreachltd/pyrokv-go
```

## Example

```go
package main

import (
	pyrokv "github.com/obsidianreachltd/pyrokv-go"
)

func main() {
	kv, err := pyrokv.NewPyroKVClient()
	if err != nil {
		panic(err)
	}
	defer kv.Close()
	if err := kv.Set("my_key", "my_value"); err != nil {
		panic(err)
	}
	var res string
	if err := kv.Get("my_key", &res); err != nil {
		panic(err)
	}
	println("Value for 'my_key':", res)
}
```

### Supported Types and Encoding

This library supports encoding standard types (int, string, bool, etc.) as well as any JSON marshalable object (e.g. a struct).
