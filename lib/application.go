package lib

import (
	"fmt"
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
		fmt.Fprintf(w, "非法")
	}

	//fmt.Println(r.Method)
	body, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(body))
	ret := this.Handler(string(body))
	fmt.Fprintf(w, ret)
}

func (this *Application) Start(host string) {
	err := http.ListenAndServe(host, this)

	if err != nil {
		log.Fatal(err)
	}
}

func (this *Application) Verify(r *http.Request) bool {
	return false
}
