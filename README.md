# utils
go common utils

[![Actions](https://github.com/lyf-coder/utils/workflows/CI/badge.svg)](https://github.com/lyf-coder/utils/actions?query=workflow%3ACI)
[![GoDoc](https://godoc.org/github.com/lyf-coder/utils?status.svg)](https://godoc.org/github.com/lyf-coder/utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/lyf-coder/utils)](https://goreportcard.com/report/github.com/lyf-coder/utils)

## Install

```console
go get github.com/lyf-coder/utils
```
## Usage
    import (
        "github.com/lyf-coder/utils"
    )
    type People struct {
    	Name string
    	Age  int
    }
    
    type Student struct {
    	Name  string
    	Grade int
    }
    
    s := Student{}
    s.Name = "jack"
    s.Grade = 100
    p := People{}
    CopyStruct(s, &p) // s is src struct, p is target struct
    p.Name == s.Name == "jack" // p.Name will be same as s.Name is jack    	