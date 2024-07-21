package utility

import (
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("faedfsedj7fyki7rm76rmi87i8mmi88mi")) //os.Getenv("SESSION_SECRET")))
