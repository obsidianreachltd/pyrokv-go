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
	"pyrokvgo"
)

func main() {
	kv, err := pyrokvgo.NewPyroKVClient()
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
