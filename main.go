package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Syahreza-Ferdian/Restologi-Payment-Backend/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/veritrans/go-midtrans"
)

func createTransaction(c *gin.Context) {
	var req model.TransactionRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received request: %+v", req)

	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	custAddress := &midtrans.CustAddress{
		FName:       req.CustomerDetails.FirstName,
		LName:       req.CustomerDetails.LastName,
		Phone:       req.CustomerDetails.Phone,
		Address:     "Sigura-gura Pride",
		City:        "Malang",
		Postcode:    "16000",
		CountryCode: "IDN",
	}

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  uuid.New().String(),
			GrossAmt: int64(req.TransactionDetails.GrossAmount),
		},
		CustomerDetail: &midtrans.CustDetail{
			FName:    req.CustomerDetails.FirstName,
			LName:    req.CustomerDetails.LastName,
			Email:    req.CustomerDetails.Email,
			Phone:    req.CustomerDetails.Email,
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)

	if err != nil {
		log.Printf("Error getting token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := model.TransactionResponse{
		Token:       snapTokenResp.Token,
		SnapUrl:     snapTokenResp.RedirectURL,
		OrderID:     snapReq.TransactionDetails.OrderID,
		GrossAmount: int64(req.TransactionDetails.GrossAmount),
	}

	c.JSON(http.StatusOK, resp)
}

func main() {
	r := setupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

var Handler = setupRouter()

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	r.POST("/charge", createTransaction)
	return r
}
