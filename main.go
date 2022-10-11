package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const BaseURL = "https://api.veriphone.io/v2"
const ApiKey = "916C4BBDBD044FC0873378A06FDEE51E"

func main() {
	//repo := newVerifyPhoneRepo(os.Getenv("API"))
	//verify, err := repo.validatePhone(context.Background(), "15123577723", "US")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(verify)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHTTPFunction)
	r.HandleFunc("/verify/{phone}", VerifyHTTPFunction)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func HomeHTTPFunction(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", r.Method)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"Message":"App running"}`)
}

func VerifyHTTPFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	phone := vars["phone"]
	log.Printf("%s", vars["phone"])

	repo := newVerifyPhoneRepo(ApiKey)
	verify, err := repo.validatePhone(context.Background(), phone, "US")
	if err != nil {
		log.Printf("%v", err)
	}

	json.NewEncoder(w).Encode(verify)
}

type verifyPhoneRepo struct {
	Key    string
	client *http.Client
}

func newVerifyPhoneRepo(key string) *verifyPhoneRepo {
	return &verifyPhoneRepo{
		Key:    key,
		client: http.DefaultClient,
	}
}

type verifyResponse struct {
	Status              string `json:"status"`
	Phone               string `json:"phone"`
	PhoneValid          bool   `json:"phone_valid"`
	PhoneType           string `json:"phone_type"`
	PhoneRegion         string `json:"phone_region"`
	Country             string `json:"country"`
	CountryCode         string `json:"country_code"`
	CountryPrefix       string `json:"country_prefix"`
	InternationalNumber string `json:"international_number"`
	LocalNumber         string `json:"local_number"`
	E164                string `json:"e164"`
	Carrier             string `json:"carrier"`
}

func (r *verifyPhoneRepo) validatePhone(ctx context.Context, phone, defaultCountry string) (verifyResponse, error) {
	url := fmt.Sprintf("%s/verify?phone=%s&key=%s&default_country=%s", BaseURL, phone, r.Key, defaultCountry)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return verifyResponse{}, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return verifyResponse{}, err
	}
	defer res.Body.Close()
	var result verifyResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return verifyResponse{}, err
	}
	return result, nil
}
