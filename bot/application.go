package bot

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"github.com/dueros/bot-sdk-go/bot/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Application struct {
	AppId              string
	DisableCertificate bool
	DisableVerifyJson  bool
	Handler            func(rawRequest string) string
}

func (this *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	r = r.WithContext(context.WithValue(r.Context(), "requestBody", body))

	if !this.Verify(w, r) {
		return
	}

	ret := this.Handler(string(body))
	w.Write([]byte(ret))
}

func (this *Application) Start(host string) {
	err := http.ListenAndServe(host, this)

	if err != nil {
		log.Fatal(err)
	}
}

func (this *Application) Verify(w http.ResponseWriter, r *http.Request) bool {
	if !this.DisableVerifyJson && !verifyJSON(w, r, this.AppId) {
		return false
	}

	if !this.DisableCertificate && !validateRequest(w, r) {
		return false
	}
	return true
}

func HTTPError(w http.ResponseWriter, logMsg string, err string, errCode int) {
	if logMsg != "" {
		log.Println(logMsg)
	}

	http.Error(w, err, errCode)
}

// Decode the JSON request and verify it.
func verifyJSON(w http.ResponseWriter, r *http.Request, appId string) bool {
	var req *model.Request

	body := r.Context().Value("requestBody").([]byte)

	if err := json.Unmarshal(body, &req); err != nil {
		HTTPError(w, err.Error(), "Bad Request", 400)
		return false
	}

	// Check the timestamp
	if !req.VerifyTimestamp() && r.URL.Query().Get("_dev") == "" {
		HTTPError(w, "Request too old to continue (>180s).", "Bad Request", 400)
		return false
	}

	// Check the app id
	if !req.VerifyBotID(appId) {
		HTTPError(w, "DuerOS BotID mismatch!", "Bad Request", 400)
		return false
	}
	return true
}

// Run all mandatory DuerOS security checks on the request.
func validateRequest(w http.ResponseWriter, r *http.Request) bool {
	// Check for debug bypass flag
	devFlag := r.URL.Query().Get("_dev")

	isDev := devFlag != ""

	if !isDev {
		isRequestValid := IsValidRequest(w, r)
		if !isRequestValid {
			return false
		}
	}
	return true
}

// IsValidRequest handles all the necessary steps to validate that an incoming http.Request has actually come from
// the DuerOS service. If an error occurs during the validation process, an http.Error will be written to the provided http.ResponseWriter.
// The required steps for request validation can be found on this page:
// https://dueros.baidu.com/didp/doc/dueros-bot-platform/dbp-deploy/authentication_markdown
func IsValidRequest(w http.ResponseWriter, r *http.Request) bool {
	certURL := r.Header.Get("SignatureCertUrl")

	// Verify certificate URL
	if !verifyCertURL(certURL) {
		HTTPError(w, "Invalid cert URL: "+certURL, "Not Authorized", 401)
		return false
	}

	// Fetch certificate data
	certContents, err := readCert(certURL)
	if err != nil {
		HTTPError(w, err.Error(), "Not Authorized", 401)
		return false
	}

	// Decode certificate data
	block, _ := pem.Decode(certContents)
	if block == nil {
		HTTPError(w, "Failed to parse certificate PEM.", "Not Authorized", 401)
		return false
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		HTTPError(w, err.Error(), "Not Authorized", 401)
		return false
	}

	// Check the certificate date
	if time.Now().Unix() < cert.NotBefore.Unix() || time.Now().Unix() > cert.NotAfter.Unix() {
		HTTPError(w, "DuerOS certificate expired.", "Not Authorized", 401)
		return false
	}

	// Check the certificate alternate names
	foundName := false
	for _, altName := range cert.Subject.Names {
		if altName.Value == "dueros-api.baidu.com" {
			foundName = true
		}
	}

	if !foundName {
		HTTPError(w, "DuerOS certificate invalid.", "Not Authorized", 401)
		return false
	}

	// Verify the key
	publicKey := cert.PublicKey
	encryptedSig, _ := base64.StdEncoding.DecodeString(r.Header.Get("Signature"))

	// Make the request body SHA1 and verify the request with the public key
	//var bodyBuf bytes.Buffer
	hash := sha1.New()
	_, err = io.Copy(hash, bytes.NewReader(r.Context().Value("requestBody").([]byte)))
	if err != nil {
		HTTPError(w, err.Error(), "Internal Error", 500)
		return false
	}

	err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), encryptedSig)
	if err != nil {
		HTTPError(w, "Signature match failed.", "Not Authorized", 401)
		return false
	}

	return true
}

func readCert(certURL string) ([]byte, error) {
	cert, err := http.Get(certURL)
	if err != nil {
		return nil, errors.New("Could not download DuerOS cert file.")
	}
	defer cert.Body.Close()
	certContents, err := ioutil.ReadAll(cert.Body)
	if err != nil {
		return nil, errors.New("Could not read DuerOS cert file.")
	}

	return certContents, nil
}

func verifyCertURL(path string) bool {
	link, _ := url.Parse(path)

	if link.Scheme != "https" {
		return false
	}

	if link.Host != "duer.bdstatic.com" {
		return false
	}

	return true
}
