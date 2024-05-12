package controller

import (
	"net/http"
	"azrielrisywan/be-assignment-payment/config"
	"github.com/gin-gonic/gin"
)

type TransactionListByUserRequest struct {
	IdUsers int `json:"idUser"`
}

type Transaction struct {
	IdUsers int `json:"idUser"`
	IdAccountFrom int `json:"idAccountFrom"`
	IdAccountTo *int `json:"idAccountTo"`
	Amount int `json:"amount"`
	PaymentType string `json:"paymentType"`
	TransactionDate string `json:"transactionDate"`
}

func TransactionListByUser(ctx *gin.Context) {
	var req TransactionListByUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.SetupDatabase()

	sqlQuery := `SELECT 
					i_id_users, 
					i_id_accounts_from, 
					i_id_accounts_to, 
					v_amount,
					CASE
						WHEN i_payments = '1' THEN 'SEND'
						WHEN i_payments = '2' THEN 'WITHDRAW'
					END AS paymentType, 
					d_created_at 
				FROM be_assignment.payments 
				WHERE i_id_users = $1 
				ORDER BY d_created_at DESC`
	rows, err := db.Query(sqlQuery, req.IdUsers)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		var idAccountTo *int
		err := rows.Scan(&transaction.IdUsers, &transaction.IdAccountFrom, &transaction.IdAccountTo, &transaction.Amount, &transaction.PaymentType, &transaction.TransactionDate)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if idAccountTo != nil {
			transaction.IdAccountTo = idAccountTo
		}
		transactions = append(transactions, transaction)
	}

	ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
}