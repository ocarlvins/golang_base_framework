package print

import (
	"fmt"
	"net/http"
)

func PrintToFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Printing to file")

}
