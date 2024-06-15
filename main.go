package main

import (
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	midclient := midtrans.NewClient()
	midclient.ServerKey = serverKey
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	orderID := uuid.New().String()
	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: req.Amount,
		},
		CustomerDetail: &midtrans.CustDetail{
			FName: "",
			LName: "",
			Email: "",
			Phone: "",
		},
	}

	snapResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := model.TransactionResponse{
		Token:       snapResp.Token,
		SnapUrl:     snapResp.RedirectURL,
		OrderID:     orderID,
		GrossAmount: req.Amount,
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

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	r.POST("/create-transaction", createTransaction)
	return r
}
