package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
	"github.com/joho/godotenv"
)

func Secret(user, realm string) string {
	if user == "felippe" {
		val := os.Getenv("FELIPPE_PASSWORD")
		fmt.Println(val)
		return val
	}
	return ""
}

func main() {
	godotenv.Load()
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <diretorio> <porta>")
		os.Exit(1)
	}
	httpDir := os.Args[1]
	httpPort := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("meuserver.com", Secret)
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))
	fmt.Println("Subindo servidor na porta", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, nil))
}
