package ps1_fun

import (
	"context"
	"net/http"
	"os"
	"strings"
)


func root(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`
echo PS1=\"coolprompt baby\> \" >> ~/.ps1.fun
echo "source ~/.ps1.fun" >> ~/.bashrc

`))
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
