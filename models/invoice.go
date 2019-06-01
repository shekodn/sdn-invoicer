package models

import (
  "fmt"

  "github.com/jinzhu/gorm"

  u "github.com/shekodn/sdn-invoicer/utils"
)

type Invoice struct {
  gorm.Model
  IsPaid      bool      `json:"is_paid"`
  Amount      int       `json:"amount"`
  Description string    `gorm:"default:'Default description'" json:"is_description"`
}

  // PaymentDate time.Time `json:"payment_date"`
  // DueDate     time.Time `json:"due_date"`
  // Charges     []Charge  `json:"charges"`


func typeof(v interface{}) string {
    return fmt.Sprintf("%T", v)
}

func (iv *Invoice) Validate() (map[string] interface{}, bool) {

    if iv.Amount <= 0 {
        fmt.Println("here", iv.Amount)
        return u.Message(false, "Amount should be greater than 0"), false
    }

    return u.Message(true, "success"), true
}


func (iv *Invoice) Create() (map[string] interface {}) {

  if resp, ok := iv.Validate(); !ok {
    return resp
  }

  GetDB().Create(iv)

  resp := u.Message(true, "success")
  resp["iv"] = iv
  return resp

}



func GetInvoice(id int) (*Invoice) {
    iv := &Invoice{}
    err := GetDB().First(&iv, id).Error

    if err != nil {
        fmt.Println("Error:", err)
        return iv
    }

    GetDB().First(&iv, id)
    return iv

}

// func GetContacts(user uint) ([]*Contact) {
//     contacts := make([]*Contact, 0)
//     err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
//
//     if err != nil {
//         fmt.Println(err)
//         return nil
//     }
//
//     return contacts
// }




// func (iv *Invoice) Get(id int) {
//
//   // iv.db.First(&iv, id)
//
//   a := GetDB().First(&iv, id)
//
//
//   fmt.Printf("!!!!!!%+v\n", a)
//
//   // GetDB().Where("invoice_id = ?", &iv.ID).Find
//
//   // return AuxIv
// //   jsonInvoice, err := json.Marshal(i1)
// // // if err != nil {
// // //   httpError(w, r, http.StatusInternalServerError, "failed to retrieve invoice id %d: %s", vars["id"], err)
// // //   return
// // // }
// }
