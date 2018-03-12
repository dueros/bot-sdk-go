package bot

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Application struct {
	enableCertificate    bool
	enableVerifyTimespan bool
	Handler              func(rawRequest string) string
}

func (this *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !this.Verify(r) {
		w.WriteHeader(403) // TODO
		return
	}

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	ret := this.Handler(string(body))
	w.Write([]byte(ret))
}

func (this *Application) Start(host string) {
	err := http.ListenAndServe(host, this)

	if err != nil {
		log.Fatal(err)
	}
}

func (this *Application) Verify(r *http.Request) bool {
	return true
}
