package fauna

import (
	"os"

	f "github.com/fauna/faunadb-go/v4/faunadb"
)

func Client() *f.FaunaClient {
	return f.NewFaunaClient(os.Getenv("FAUNA_SECRET"))
}
