# A lightweight cloudflare go library
## WIP!!!
Ever find yourself needing a lightweight cloudflare go library to manage your dns and nothing else???? Well this is for you! With very simple syntax.

```go
package main 

import (
  "log"
  "fmt"
  
  "github.com/windowscmd/cloudflare-go"
)

fucn main(){
  
  //For auth with email and a token
  api, err := cloudflare.New(os.Getenv("CF_TOKEN"), os.Getenv("CF_EMAIL"))
  
  //For auth with just an API key
  //api_with_key, err := cloudflare.New(os.Getenv("CF_API_KEY"))
  
  if err != nil {
    log.Fatal(err)
  }
  
}
```
