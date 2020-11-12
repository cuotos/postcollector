package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)


func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(getPostHandler(os.Stdout))

	fmt.Println("listening for requests on 0.0.0.0:3000")
	if err := http.ListenAndServe("0.0.0.0:3000", r); err != nil {
		panic(err)
	}

}

func getPostHandler(output io.Writer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		d, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Fprintf(output, `### START ###
%v
### END ###
`, string(d))
	}
}