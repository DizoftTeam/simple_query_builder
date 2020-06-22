# Simple Query Builder

![](https://badges.fyi/github/latest-tag/DizoftTeam/simple_query_builder)
![](https://badges.fyi/github/license/DizoftTeam/simple_query_builder)
![](https://badges.fyi/github/stars/DizoftTeam/simple_query_builder)

Simple query builder for build SQL queries

## Example

Simple example below

```go
package main

import (
  "github.com/DizoftTeam/sqb"

  "fmt"
)

func main() {
	query := sqb.NewQuery("users as u").
		Select("u.name, u.email").
    Where("and", "u.name = 'Ruslan'").
    Generate()

  // TODO: Here you can use generated query as you want

	fmt.Printf("SQL: %v\n", query)
}
```
