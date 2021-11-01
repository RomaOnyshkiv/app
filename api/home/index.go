package home

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

var user = User{
	Name: "Stranger",
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s, welome!", user.Name)
	log.Info("HomePage opened")
}
