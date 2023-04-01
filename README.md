# simplemux

# Usage
```
package main

import (
	"net/http"

	"github.com/Acebond/simplemux"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	mux := &simplemux.Mux{}
	mux.HandlerFunc(http.MethodGet, "/test", testHandler)
	http.ListenAndServe(":8080", mux)
}

```
