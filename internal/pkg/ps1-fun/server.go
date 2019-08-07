package ps1_fun

import (
	"context"
	"net/http"
	"os"
	"strings"
)


func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("oh hey there"))
}

func Serve(ctx context.Context) error {

	http.HandleFunc("/", root)

	var port = ":8080"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")

		if !strings.HasPrefix(port, ":") {
			port = ":" + port
		}
	}


	if err := http.ListenAndServe(port, nil); err != nil {
		return err
	}

	return nil
}
