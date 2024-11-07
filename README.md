# Pretty

Example usage:

```go
package main

import (
	"github.com/hyperxpizza/pretty"
)

func main() {
	uglyObject := map[string]any{"ugly": "object","not": "pretty"}
	pretty.Print(uglyObject)
}

// Outputs:
// {
// 	"ugly": "object",
// 	"not": "pretty"
// }


```
