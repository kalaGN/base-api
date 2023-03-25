package clickhouse

import (
	"agn/pkg/config"
	"fmt"
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

type UserRequestLog struct {
	Project string `db:"project"`
	Apikey  string `db:"apikey"`
	Pid     int64  `db:"pid"`
	Input   string `db:"input"`
	Data    string `db:"data"`
	Extend  string `db:"extend"`
	Sid     string `db:"sid"`
}

func CkConnect() (ur []UserRequestLog) {
	source := "tcp://" + config.GetString("clickhouse.ip") + ":" + config.GetString("clickhouse.port") + "?debug=true&database=" + config.GetString("clickhouse.database") + ""
	connect, err := sqlx.Connect("clickhouse", source)
	if err != nil {
		fmt.Printf("clickhouse open err %s", err.Error())
	}
	defer func() {
		_ = connect.Close()
	}()

	var items []UserRequestLog
	if err := connect.Select(&items, "select project,apikey,pid,input,data,extend,sid from ods_financial_log.user_request_log limit 10"); err != nil {
		log.Fatal(err)
	}
	return items
	// for _, item := range items {
	// 	log.Fatalf("Apikey: %s, Input: %s, Sid: %s\n", item.Apikey, item.Input, item.Sid)
	// 	fmt.Printf("Apikey: %s, Input: %s, Sid: %s\n", item.Apikey, item.Input, item.Sid)
	// }
}
