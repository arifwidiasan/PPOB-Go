package helper

import (
	"fmt"
	"net/smtp"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	"github.com/leekchan/accounting"
)

func SendEmail(email string, transaction model.Transaction, virtual_account model.Virtual_account) error {
	ac := accounting.Accounting{Symbol: "Rp. ", Precision: 2, Thousand: ".", Decimal: ","}
	price := ac.FormatMoney(transaction.Price)

	link_payment := fmt.Sprintf("%s%s%s%d", "https://upay.exzork.me/v1/payments/", transaction.CodeTransaction, "/", transaction.Price)

	server := "smtp-mail.outlook.com"
	port := 587
	user := "upay-app@outlook.com"
	from := user
	pass := "budipekerti123"
	dest := email

	auth := LoginAuth(user, pass)

	to := []string{dest}

	msg := []byte("From: " + from + "\n" +
		"To: " + dest + "\n" +
		"Subject: Upay - Transaksi " + transaction.CodeTransaction + "\n\n" +
		"Berikut Detail Transaksi Anda \n\n" +
		"Kode Transaksi \t: " + transaction.CodeTransaction + "\n" +
		"Nama Produk \t: " + transaction.Product.Name + "\n" +
		"Charge Number  : " + transaction.ChargeNumber + "\n" +
		"Harga\t\t\t: " + price + "\n" +
		"Pembayaran \t  : Virtual Account - " + transaction.Payment_Method.CodeBank + "\n\n" +
		"Dimohon untuk segera melakukan pembarayan ke nomor berikut : \n\n" + virtual_account.VANumber + "\n\n" +
		"Atau bisa juga dengan klik tautan berikut untuk melakukan pembayaran : \n\n" + link_payment + "\n\n" +
		"Terima kasih telah melakukan transaksi di Upay !")

	endpoint := fmt.Sprintf("%v:%v", server, port)
	err := smtp.SendMail(endpoint, auth, from, to, msg)
	if err != nil {
		return err
	}

	return nil
}
