package main

import (
	"fmt"

	"github.com/tiagonevestia/monorepo-go/packages/main_app/router"
)

func main() {
	fmt.Println("Run main application router")
	_ = router.GetEngine().Run(":8080")
}

func Check() {
	fmt.Println("Check")
}
