package controller

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/razorpay/razorpay-go"
)

type PageVariables struct {
	OrderId string
	Email   string
	Name    string
	Amount  int
	Contact string
}

func App(c *gin.Context){
     page := &PageVariables{}
	 page.Amount = 1000
	 page.Email = "ratheeshgopinadhkumar@gmail.com"
	 page.Name = "Ratheesh G Kumar"
	 page.Contact = "8762334325"
	 client := razorpay.NewClient(os.Getenv("RAZOR_PAY_KEY_ID"),os.Getenv("RAZOR_PAY_KEY_SECRET"))

	 data := map[string]interface{}{
		   "amount": page.Amount,
		   "currency":"INR",
		   "receipt":"some_receipt_id",
	 }
	 body, err := client.Order.Create(data,nil)
	 fmt.Println("/////////////receipt",body)
	 if err != nil{
		fmt.Println("Problem getting the repository information",err)
		os.Exit(1)
	 }

	 value := body["id"]
	 str := value.(string)
	 fmt.Println("str//////////",str)
	 HomePageVar := PageVariables{
		OrderId: str,
		Amount: page.Amount/100,
		Email: page.Email,
		Name: page.Name,
		Contact: page.Contact,
	 }
	 c.HTML(200,"app.html",HomePageVar)
}

func PaymentStatus(c *gin.Context){
	paymentid := c.Query("paymentid")
	orderid := c.Query("orderid")
	signature := c.Query("signature")

	fmt.Println("paymentid",paymentid)
	fmt.Println("orderid",orderid)
	fmt.Println("signature",signature)

}
func PaymentFailure(c*gin.Context){

}
