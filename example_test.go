package url_test

import (
	"fmt"
	"log"

	"github.com/akthrmsx/url"
)

func ExampleParse() {
	u, err := url.Parse("https://github.com/akthrmsx")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u)
	// Output:
	// https://github.com/akthrmsx
}
