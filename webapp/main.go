package webapp

import (
    "net/http"
    "github.com/magiconair/properties"
)

func Run() {
    createDatabaseTable()
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.ListenAndServe(":8080", nil)
}

func HealthCheck() {
    healthCheck := "/etc/conf.d/ot-go-webapp/healthcheck.properties"
    healthVaules := properties.MustLoadFiles([]string{healthCheck}, properties.UTF8, true)

    healthy := healthVaules.GetString("healthy", "healthy")
    livecheck := healthVaules.GetString("livecheck", "livecheck")

    mux := http.NewServeMux()

    if healthy == "true" {
        mux.HandleFunc("/healthy", returnCode200)
    } else {
        mux.HandleFunc("/healthy", returnCode404)
    }

    if livecheck == "true" {
        mux.HandleFunc("/livecheck", returnCode200)
    } else {
        mux.HandleFunc("/livecheck", returnCode404)
    }
    return mux
}

func returnCode200(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("☄ HTTP status code returned!"))
}

func returnCode404(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("☄ HTTP status code returned!"))
}
