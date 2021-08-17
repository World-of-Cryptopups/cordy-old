package fauna

import (
	"fmt"
	"os"

	f "github.com/fauna/faunadb-go/v4/faunadb"
)

func Client() *f.FaunaClient {
	fmt.Println(os.Getenv("FAUNA_SECRET"))

	return f.NewFaunaClient(os.Getenv("FAUNA_SECRET"))
}
