package controllers

import (
  "net/http"
  "log"
)

var GetIndex = func(w http.ResponseWriter, r *http.Request) {
	log.Println("serving index page")
	w.Write([]byte(`
<!DOCTYPE html>
<html>
    <head>
        <title>Invoicer Web</title>
        <script src="statics/jquery-1.12.4.min.js"></script>
        <script src="statics/invoicer-cli.js"></script>
        <link href="statics/style.css" rel="stylesheet">
    </head>
    <body>
	<h1>Invoicer Web</h1>
        <p class="desc-invoice"></p>
        <div class="invoice-details">
        </div>
        <h3>Request an invoice by ID</h3>
        <form id="invoiceGetter" method="GET">
            <label>ID :</label>
            <input id="invoiceid" type="text" />
            <input type="submit" />
        </form>
        <form id="invoiceDeleter" method="DELETE">
            <label>Delete this invoice</label>
            <input type="submit" />
        </form>


        <h1>New Invoice</h1>
        <form method="POST" action="/invoice">
        	<label>is_paid:</label><br />
        	<input type="text" name="is_paid"><br />
        	<label>amount:</label><br />
        	<input type="text" name="amount"><br />
        	<input type="submit">
        </form>

    </body>
</html>`))
}
