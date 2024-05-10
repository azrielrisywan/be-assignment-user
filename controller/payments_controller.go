package controller

import(
	"azrielrisywan/be-assignment-user/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getPaymentsListByUserRequest struct {
	IdUser int `json:"idUser"`
}

type Payment struct {
	ID          int    `json:"idPayments"`
	Email       string `json:"email"`
	Amount      int    `json:"amount"`
	IDAccountFrom int `json:"idAccountFrom"`
	IDAccountTo *int `json:"idAccountTo"`
	PaymentType string `json:"paymentType"`
}

func GetPaymentsListByUser(ctx *gin.Context) {
	var req getPaymentsListByUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.SetupDatabase()

	sqlQuery := `SELECT 
		t1.i_id AS idPayments,
		t2.n_email AS email,
		t1.v_amount AS amount,
		t1.i_id_accounts_from as idAccountFrom,
		t1.i_id_accounts_to as idAccountTo,
		CASE
			WHEN t1.i_payments = '1' THEN 'SEND'
			WHEN t1.i_payments = '2' THEN 'WITHDRAW'
		END AS paymentType
	FROM 
		be_assignment.payments t1
	JOIN 
		be_assignment.users t2 ON t2.i_id = t1.i_id_users 
	JOIN 
		be_assignment.accounts t3 ON t3.i_id = t1.i_id_accounts_from
	WHERE 
		t1.i_id_users = $1
	union all 
	SELECT 
		t1.i_id AS idPayments,
		t2.n_email AS email,
		t1.v_amount AS amount,
		t1.i_id_accounts_from as idAccountFrom,
		t1.i_id_accounts_to as idAccountTo,
		CASE
			WHEN t1.i_payments = '1' THEN 'RECEIVE'
		END AS paymentType
	FROM 
		be_assignment.payments t1 
	JOIN 
		be_assignment.accounts t3 ON t3.i_id = t1.i_id_accounts_to
	join 
		be_assignment.users t2 on t2.i_id = t3.i_id_users 
	WHERE 
		t2.i_id = $1`

	rows, err := db.Query(sqlQuery, req.IdUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var payments []Payment
	for rows.Next() {
		var payment Payment
		var idAccountTo *int
		err := rows.Scan(&payment.ID, &payment.Email, &payment.Amount, &payment.IDAccountFrom, &payment.IDAccountTo, &payment.PaymentType)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if idAccountTo != nil {
			payment.IDAccountTo = idAccountTo
		}
		payments = append(payments, payment)
	}

	ctx.JSON(http.StatusOK, gin.H{"payments": payments})
}