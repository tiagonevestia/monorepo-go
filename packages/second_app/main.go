package main

import (
	"fmt"

	"github.com/tiagonevestia/monorepo-go/packages/second_app/router"
)

func main() {
	fmt.Println("Run main second router")
	_ = router.GetEngine().Run(":8081")
}
