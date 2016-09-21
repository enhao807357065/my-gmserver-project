package dao

import (
	"database/sql"
	_ "mysql"
)

var Db *sql.DB

func init() {

	//var err error
	//Db, err = sql.Open("mysql", "root:liang123@tcp(192.168.2.126:3306)/fashion_boutique?charset=utf8")
	//fmt.Println("err: ", err)
	//if err != nil {
	//	panic("db can't connect!")
	//}
	//
	//// set MaxOpenConns
	//Db.SetMaxOpenConns(20)
	//
	//// set MaxIdleConns
	//Db.SetMaxIdleConns(5)
	//
	//fmt.Println("orm1--------------------------------------------")
	//arr, err := testQuery()
	//fmt.Println("arr: ", arr)
}

// 帮助请求结构体
type HelpRequest struct {
	HelpRequestId int64  `json:"help_request_id"` // 求助记录id
	AppAccountId  int64  `json:"app_account_id"`  // 求助者id
	Type          int    `json:"type"`            // 求助类型
	ItemId        string `json:"item_id"`         // 求助物品id
	Number        string `json:"number"`          // 请求数量
	RequestTime   int64  `json:"request_time"`    // 求助时间
}

func testQuery() (items []*HelpRequest, err error) {

	var stmt *sql.Stmt
	var rows *sql.Rows

	stmt, err = Db.Prepare("select id, app_account_id, type, item_id, number, request_time from help_request")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err = stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	items = make([]*HelpRequest, 0)

	for rows.Next() {

		item := &HelpRequest{}

		err = rows.Scan(&item.HelpRequestId, &item.AppAccountId, &item.Type, &item.ItemId, &item.Number, &item.RequestTime)
		if err != nil {
			return
		}

		items = append(items, item)
	}

	return
}