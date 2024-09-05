package main

import (
	"RestAPI/server"
)

func main() {
	/* OLD
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
		fmt.Fprintf(w, "Hello there %s", "visitor")
	})
	srv := http.Server{
		Addr: ":8080",
	}
	*/
	srv := server.New(":8080")

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
