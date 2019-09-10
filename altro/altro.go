package altro

import "../mydb"

func Store(k, v string) {
	mydb.New()
	mydb.Store(k, v)
}
