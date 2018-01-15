package model

import (
	"database/sql"
	. "website/types"
)

type Mdl_masuk struct{}

func (mdl_masuk Mdl_masuk) Get_pengguna(ty_masuk Ty_masuk) bool {

	db, _ := connect()
	defer db.Close()

	var username string
	var result error

	var row = db.QueryRow("select nama_pengguna from pengguna where nama_pengguna = ? and kata_sandi = ?", ty_masuk.Nama_pengguna, ty_masuk.Kata_sandi)

	result = row.Scan(&username)

	if result == sql.ErrNoRows {
		return false
	} else {
		return true
	}

}
