package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func generate(w http.ResponseWriter, r *http.Request) {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	cmd := dir + "/build-web.sh"
	out, err := exec.Command(cmd).Output()

	if err != nil {
		log.Print(err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Problem Generating Satis Repository\n%s\n%s", err, string(out[:]))
		return
	}

	log.Print(string(out[:]))

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func main() {

	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	http.HandleFunc("/generate", generate)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
