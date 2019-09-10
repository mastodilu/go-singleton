package altro

import "../mydb"

func Store(k, v string) {
	mydb.Store(k, v)
}
