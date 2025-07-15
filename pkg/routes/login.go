package routes

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func AuthJSTORHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Starting Auth Handler: %s", r.Host)
	
	for _, value := range r.Cookies() {
		log.Printf("Header %s: %s", value.Name, value.Value)
	}

	// Get UUID cookie
	uuidCookie, err := r.Cookie("UUID")
	var uuid string
	if err != nil {
		log.Printf("Failed to get cookie: %s", err)
	} else {
		log.Printf("Got Cookie"	)
		log.Printf("Cookie: %s", uuidCookie.Value)
		uuid = uuidCookie.Value
	}
	log.Printf("Cookie UUID: %s", uuid)
	toQuery := r.URL.Query().Get("to")
	log.Printf("Query: %s", toQuery)

	newURL, err := url.Parse(toQuery)
	if err != nil {
		log.Printf("Failed to parse: %s", err)
		http.Error(w, "failed to parse to query", http.StatusBadRequest)
		return
	}
	if newURL.Path != "/about" || newURL.Path == "about" {
		newURL.Path = "/"
	}
    newCookie := http.Cookie{
        Name:     "uuid",
        Value:    uuid,
        Path:     "/",
        MaxAge:   3600 * 24,
        Secure:   true,
        SameSite: http.SameSiteNoneMode,
		Domain: ".jstor.org",
    }

	// Add the session UUID as a URL fragment
	finalURL := fmt.Sprintf("%s#%s", newURL.String(), uuid)

	http.SetCookie(w, &newCookie)

	// Redirect to the redirect URL
	w.Header().Set("Location", finalURL)
	w.WriteHeader(http.StatusFound)
}
