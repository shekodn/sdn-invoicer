package main

import (
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"os"
	// "strconv"
	// "time"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "github.com/wader/gormstore"
	"go.mozilla.org/mozlog"

  "github.com/shekodn/sdn-invoicer/controllers"
)

func init() {
  // init logger
  mozlog.Logger.LoggerName = "invoice"
  log.SetFlags(0)
}


func main() {
  // register routes
  r := mux.NewRouter()
  r.HandleFunc("/", controllers.GetIndex).Methods("GET")
  // r.HandleFunc("/__heartbeat__", getHeartbeat).Methods("GET")
  r.HandleFunc("/invoice/{id:[0-9]+}", controllers.GetInvoice).Methods("GET")
  r.HandleFunc("/invoice", controllers.CreateInvoice).Methods("POST")
  // r.HandleFunc("/invoice/{id:[0-9]+}", controllers.PutInvoice).Methods("PUT")
  // r.HandleFunc("/invoice/{id:[0-9]+}", controllers.DeleteInvoice).Methods("DELETE")
  // r.HandleFunc("/invoice/delete/{id:[0-9]+}", controllers.DeleteInvoice).Methods("GET")
  // r.HandleFunc("/__version__", getVersion).Methods("GET")

  port := os.Getenv("PORT")

  if port == "" {
    port = "8000"
  }
  fmt.Println("Listening in " + port)
  err := http.ListenAndServe(":" + port, r)
  if err != nil {
    log.Println("Error:", err)
  }


}
