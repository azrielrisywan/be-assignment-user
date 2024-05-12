package controller

import (
	"azrielrisywan/be-assignment-user/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAccountsByUserRequest struct {
	IdUser    int `json:"idUser"`
}

type Account struct {
	ID          int    `json:"idAccounts"`
	Email       string `json:"email"`
	AccountType string    `json:"accountType"`
	Balance     int    `json:"balance"`
}

// Get Accounts By User godoc
// @Summary Get Accounts By User
// @Schemes
// @Description Get Accounts By User, with body request idUser Integer
// @Tags BE-USER
// @Accept json
// @Produce json
// @Param idUser body int true "idUser"
// @Success 200 {string} Success 
// @Router /getAccountsByUser [post]
func GetAccountsByUser(ctx *gin.Context) {
	var req getAccountsByUserRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	db := config.SetupDatabase()

	sqlQuery := `SELECT 
	t1.i_id as idAccounts,
	t2.n_email as email,
	t1.i_acc_type as accountType,
	t1.v_balance as balance
	FROM be_assignment.accounts t1
	join be_assignment.users t2 on t2.i_id = t1.i_id_users 
	where t1.i_id_users = $1`

	rows, err := db.Query(sqlQuery, req.IdUser)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	var accounts []Account
    for rows.Next() {
        var account Account
        rows.Scan(&account.ID, &account.Email, &account.AccountType, &account.Balance)
        accounts = append(accounts, account)
    }

    ctx.JSON(http.StatusOK, gin.H{"accounts": accounts})
}