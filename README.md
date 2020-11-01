# Go client for the Spontit API [![PkgGoDev](https://pkg.go.dev/badge/github.com/FryDay/spontit-go)](https://pkg.go.dev/github.com/FryDay/spontit-go)

Spontit-go is an unofficial [Spontit](https://spontit.com/) client for the [Go](http://www.golang.org/) programming language.

## Installation
`go get github.com/FryDay/spontit-go`

## Usage
Create a `.env` file in the root of your project and add the following lines:
```shell
SPONTIT_USERID=your_user_id
SPONTIT_KEY=your_private_key
```

Then to create a Spontit client:
```go
package main

import (
	"log"

	"github.com/FryDay/spontit-go"
)

func main() {
	spontitClient, err := spontit.NewClient()
	if err != nil {
		log.Fatal(err)
    }

    // Do things with the client here...
}
```

## Progress

### API

- [x] Push
- [x] Categories
- [ ] Channel
  - [x] Create
  - [ ] Get
  - [ ] Change
  - [x] Delete
  - [ ] Profile Image
- [ ] Channels
- [ ] Followers
