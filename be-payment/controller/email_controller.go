package controller

import (
	"azrielrisywan/be-assignment-payment/config"
)

func GetEmailByUser(idUser int) (email string) {
	db := config.SetupDatabase()

	sqlQuery := `SELECT 
					n_email 
				FROM be_assignment.users 
				WHERE i_id = $1`
	rows, err := db.Query(sqlQuery, idUser)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&email)
		if err != nil {
			panic(err)
		}
	}

	return email
}