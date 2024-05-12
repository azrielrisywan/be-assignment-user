package controller

import (
	"net/http"
	"azrielrisywan/be-assignment-payment/config"
	"github.com/gin-gonic/gin"
	"log"
)

type PaymentRequest struct {
	IdUsers int `json:"idUser"`
	IdAccountFrom int `json:"idAccountFrom"`
	IdAccountTo int `json:"idAccountTo"`
	Amount int `json:"amount"`
}

func SendPayment(ctx *gin.Context) {
	var req PaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if condition := req.IdAccountFrom == req.IdAccountTo; condition {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account from and account to cannot be the same"})
		return
	}

	// Check if account from belongs to user
	isAccountFromBelongsToUser, errCheckAccount := checkAccountUser(req.IdUsers, req.IdAccountFrom)
	if errCheckAccount != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errCheckAccount.Error()})
		return
	}
	if !isAccountFromBelongsToUser {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account from does not belong to user"})
		return
	}

	// Check if balance is enough
	isBalanceEnough, errCheckBalance := checkBalance(req.IdAccountFrom, req.Amount)
	if errCheckBalance != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errCheckBalance.Error()})
		return
	}
	if !isBalanceEnough {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Balance is not enough"})
		return
	}

	// Reduce balance from account
	_, err := reduceBalance(req.IdAccountFrom, req.Amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Insert payment record
	db := config.SetupDatabase()

	sqlQuery := "INSERT INTO be_assignment.payments (i_id_users, i_id_accounts_from, i_id_accounts_to, v_amount, i_payments) VALUES ($1, $2, $3, $4, '1')"
	_, err = db.Exec(sqlQuery, req.IdUsers, req.IdAccountFrom, req.IdAccountTo, req.Amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Add balance to account
	_, err = addBalance(req.IdAccountTo, req.Amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Payment success"})

}

var emailCtxKey = "email"

func WithdrawPayment(ctx *gin.Context) {
	var req PaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if account from belongs to user
	isAccountFromBelongsToUser, errCheckAccount := checkAccountUser(req.IdUsers, req.IdAccountFrom)
	if errCheckAccount != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errCheckAccount.Error()})
		return
	}
	if !isAccountFromBelongsToUser {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account from does not belong to user"})
		return
	}

	// Check if balance is enough
	isBalanceEnough, errCheckBalance := checkBalance(req.IdAccountFrom, req.Amount)
	if errCheckBalance != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errCheckBalance.Error()})
		return
	}
	if !isBalanceEnough {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Balance is not enough"})
		return
	}

	// Reduce balance from account
	_, err := reduceBalance(req.IdAccountFrom, req.Amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Insert payment record
	db := config.SetupDatabase()

	sqlQuery := "INSERT INTO be_assignment.payments (i_id_users, i_id_accounts_from, i_id_accounts_to, v_amount, i_payments) VALUES ($1, $2, $3, $4, '2')"
	_, err = db.Exec(sqlQuery, req.IdUsers, req.IdAccountFrom, nil, req.Amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Payment success"})


}

func reduceBalance(idAccount int, reductionAmount int) (int, error) {
	db := config.SetupDatabase()

	sqlQuery := "UPDATE be_assignment.accounts SET v_balance = v_balance - $1 WHERE i_id = $2"

	_, err := db.Exec(sqlQuery, reductionAmount, idAccount)
	log.Println("reductionAmount : ", reductionAmount, " idAccount : ", idAccount)

	if err != nil {
		return 0, err
	}

	return reductionAmount, nil
}

func addBalance(idAccount int, additionAmount int) (int, error) {
	db := config.SetupDatabase()

	sqlQuery := "UPDATE be_assignment.accounts SET v_balance = v_balance + $1 WHERE i_id = $2"

	_, err := db.Exec(sqlQuery, additionAmount, idAccount)
	log.Println("additionAmount : ", additionAmount, " idAccount : ", idAccount)

	if err != nil {
		return 0, err
	}

	return additionAmount, nil
}

func checkAccountUser(idUser int, idAccount int) (bool, error) {
	db := config.SetupDatabase()

	sqlQuery := "SELECT i_id FROM be_assignment.accounts WHERE i_id_users = $1 AND i_id = $2"

	rows, err := db.Query(sqlQuery, idUser, idAccount)
	if err != nil {
		return false, err
	}

	if rows.Next() {
		return true, nil
	}

	return false, nil
}

func checkBalance(idAccount int, amount int) (bool, error) {
	db := config.SetupDatabase()

	sqlQuery := "SELECT v_balance FROM be_assignment.accounts WHERE i_id = $1"

	rows, err := db.Query(sqlQuery, idAccount)
	if err != nil {
		return false, err
	}

	var balance int
	if rows.Next() {
		err := rows.Scan(&balance)
		if err != nil {
			return false, err
		}
	}

	if balance < amount {
		return false, nil
	}

	return true, nil
}
