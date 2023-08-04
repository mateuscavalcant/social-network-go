package controller

import (
	CON "social-network-go/pkg/config"
)

func Get(id interface{}, what string) string {
	db := CON.DB()
	var RET string
	db.QueryRow("SELECT "+what+" AS RET FROM user1 WHERE id=?", id).Scan(&RET)
	return RET
}
