package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (ce *EchoController) SimulatePaymentController(c echo.Context) error {
	code_transaction := c.Param("code_transaction")

	price := c.Param("price")
	price_int, _ := strconv.Atoi(price)

	type Payload struct {
		Amount int `json:"amount"`
	}

	data := Payload{
		Amount: price_int,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshall")
	}

	body := bytes.NewReader(payloadBytes)
	request := fmt.Sprintf("%s%s%s", "https://api.xendit.co/callback_virtual_accounts/external_id=", code_transaction, "/simulate_payment")

	req, err := http.NewRequest("POST", request, body)
	if err != nil {
		return fmt.Errorf("error request")
	}
	req.Header.Set("Authorization", "Basic eG5kX2RldmVsb3BtZW50X09HMkhLdk40UGF3dUFrb1NxOFVmc0JvUDJYTXdDMlZIM05pNWh3dWI3OHk0THZjVzFPelYwdWFrMVJvQ3ZiOg==")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error do request")
	}
	defer resp.Body.Close()

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
	})
}
