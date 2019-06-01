package controllers

import(
  "encoding/json"
  "net/http"
  // "log"
  "fmt"
  "strconv"
  "github.com/gorilla/mux"

  // u "github.com/shekodn/sdn-invoicer/utils"

  "github.com/shekodn/sdn-invoicer/models"
)

var CreateInvoice = func(w http.ResponseWriter, r *http.Request) {

  var iv models.Invoice

  isPaid, err := strconv.ParseBool(r.FormValue("is_paid"))

  if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      w.Write([]byte(fmt.Sprintf("type mismatch in var IsPaid")))
      return
  }

  iv.IsPaid = isPaid

  amount, err := strconv.Atoi(r.FormValue("amount"))

  if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      w.Write([]byte(fmt.Sprintf("type mismatch in var Amount")))
      return
  }

  iv.Amount = amount

  if r.FormValue("is_description") != "" {

      description := r.FormValue("is_description")

      if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(fmt.Sprintf("type mismatch in var description")))
      }

      iv.Description = description
  }

  // xss := "<script type='text/javascript'>alert('xss');</script>"

  // iv := &models.Invoice{IsPaid: isPaid, Amount: amount, Amoun}

  resp := iv.Create() //Create account
  if resp["status"] != true {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(fmt.Sprintf("Error: %s", resp["message"])))
    return
  }

  w.WriteHeader(http.StatusCreated)

  jsonInvoice, err := json.Marshal(iv)
    if err != nil {
      // httpError(w, r, http.StatusInternalServerError, "failed to retrieve invoice id %d: %s", vars["id"], err)
      w.Write([]byte(fmt.Sprintf("Error: %s", "help")))
      return
    }

  w.Write(jsonInvoice)


  // _ := u.PrettyPrint(iv)
  // w.Write([]byte(fmt.Sprintf("%s", data)))
  // fmt.Printf("%+v", iv.Description)
  // fmt.Printf("%#v", iv.Description) //with type


  // // w.Header().Add("Content-Type", "text/html")
  // w.Write([]byte(fmt.Sprintf("<script type='text/javascript'>alert('xss');</script>")))
  // w.Write([]byte(fmt.Sprintf("created invoice %d ", iv.ID)))
}


var GetInvoice = func(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.Atoi(vars["id"])
  iv := models.GetInvoice(id)

  if iv.ID == 0 {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(fmt.Sprintf("No invoice id %s", vars["id"])))
    return
  }

  // w.Header().Add("Content-Type", "text/html")
  // w.Header().Add("Content-Security-Policy", "default-src 'self'; child-src 'self;")
	// w.Header().Add("Access-Control-Allow-Origin", "*")
  // jsonInvoice, _ := json.Marshal(iv)
	w.WriteHeader(http.StatusOK)
  // w.Write([]byte(fmt.Sprintf("Is Paid: %t", iv.IsPaid)))
  // w.Write([]byte(fmt.Sprintf("Description: %s", iv.Description)))
  // fmt.Println(u.PrettyPrint(iv))
  w.Write([]byte(fmt.Sprintf(`
    <h1> Invoice #%x </h1>
    <p> Is Paid: %t </p>
    <p> Amount: %x </p>
    <p> Description: %s </p>
  `, iv.ID, iv.IsPaid, iv.Amount, iv.Description)))
}




  // iv := &models.Invoice{}
  // id, _ := strconv.Atoi(vars["id"])
  // iv.db.First(&iv, id)
  //
  // fmt.Printf("%+v\n", iv)
  //
  // if i1.ID == 0 {
  //   // httpError(w, r, http.StatusNotFound, "No invoice id %s", vars["id"])
  //   w.WriteHeader(http.StatusNotFound)
  //   w.Write([]byte(fmt.Sprintf("No invoice id")))
  //   return
  // }






  // err := json.NewDecoder(r.Body).Decode(iv)
  // fmt.Println(r.Body)
	// if err != nil {
  //   fmt.Println("ERR: ", err)
  //   w.WriteHeader(http.StatusBadRequest)
  //   w.Write([]byte(fmt.Sprintf("Check Params!")))
	// 	return
	// }
  //
	// resp := iv.Create()
  // w.WriteHeader(http.StatusCreated)
  // w.Write([]byte(fmt.Sprintf("created invoice %d", iv.ID)))
  // fmt.Println("resp")
  // fmt.Println(resp)


// vars := mux.Vars(r)
// log.Println("getting invoice id", vars["id"])
// var i1 Invoice
// id, _ := strconv.Atoi(vars["id"])
// iv.db.First(&i1, id)
// fmt.Printf("%+v\n", i1)
// if i1.ID == 0 {
//   httpError(w, r, http.StatusNotFound, "No invoice id %s", vars["id"])
//   return
// }
// iv.db.Where("invoice_id = ?", i1.ID).Find(&i1.Charges)
// jsonInvoice, err := json.Marshal(i1)
// if err != nil {
//   httpError(w, r, http.StatusInternalServerError, "failed to retrieve invoice id %d: %s", vars["id"], err)
//   return
// }
// w.Header().Add("Content-Type", "application/json")
// w.Header().Add("Access-Control-Allow-Origin", "*")
// w.WriteHeader(http.StatusOK)
// w.Write(jsonInvoice)
// al := appLog{Message: fmt.Sprintf("retrieved invoice %d", i1.ID), Action: "get-invoice"}
// al.log(r)
