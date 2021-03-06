BitClient
===

BitClient is a Golang Api client for Atlassian Bitbucket REST api (v1.0).

<!-- toc -->
- [Overview](#overview)
- [Installation](#installation)
- [Getting Started](#getting-started)
- [Supported versions](#supported-version)

<!-- tocstop -->

## Overview
This library allow you to call Bitbucket api endpoint from you go programs, converting json responses to Go structs (see [types.go](types.go) for the complete list).

It's currently integrated in a command line utility [BitAdmin](https://github.com/daeMOn63/bitadmin), aimed to ease and speed up administration of Bitbucket repositories.

## Installation
Make sure you have a working Go environment.  Go version 1.2+ is supported.  [See
the install instructions for Go](http://golang.org/doc/install.html).

To install bitclient, simply run:
```
$ go get github.com/daeMOn63/bitclient
```

Dependencies are managed using [Dep](https://github.com/golang/dep). Make sure to follow the [installation instructions](https://github.com/golang/dep#setup) to setup ```dep``` on your workstation.

To install the dependencies, run from the project root:
```
$ dep ensure
```

## Getting started
The following sample will connect to the bitbucket api and retrieve a list of all projects.

```go
package main
import (
    "fmt"
    "os"
    "github.com/lotos2512/bitclient"
)

func main() {
	
	cfg := bitclient.BitConf{
		Url: "https://bitbucket.org/",
		UserName: "",
		Password: "",
    }
	
    client := bitclient.NewBitClient(cfg)

    requestParams := bitclient.PagedRequest{
        Limit: 10,
        Start: 0,
    }

    projectsResponse, err := client.GetProjects(requestParams)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    for _, project := range projectsResponse.Values {
        fmt.Printf("Project : %d - %s\n", project.Id, project.Key)
    }
}
```

## Supported versions
The library is currently developed according to the documentation of [Atlassian Bitbucket v4.13.1](https://developer.atlassian.com/static/rest/bitbucket-server/4.13.1/bitbucket-rest.html)
Other versions have not been tested yet.
