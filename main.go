package main

import (
	"log"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

func main() {
	// OAuth2 yöneticisi oluşturma
	manager := manage.NewDefaultManager()

	// Token saklama deposu olarak bellek tabanlı bir yapı kullanma
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// Client bilgilerini saklamak için bellek tabanlı bir yapı kullanma
	clientStore := store.NewClientStore()
	clientStore.Set("client_id", &models.Client{
		ID:     "client_id",
		Secret: "client_secret",
		Domain: "http://localhost:9094",
	})
	manager.MapClientStorage(clientStore)

	// OAuth2 sunucusu oluşturma
	srv := server.NewDefaultServer(manager)

	// Yetkilendirme endpoint'i
	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	// Token endpoint'i
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("OAuth2 Server is running at 9096 port.")
	log.Fatal(http.ListenAndServe(":9096", nil))
}
