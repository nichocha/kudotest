package model

import (
	"database/sql"
	. "website/types"
)

type Mdl_daftar struct{}

func (mdl_daftar Mdl_daftar) Get_pengguna(nama_pengguna string) bool {

	db, _ := connect()
	defer db.Close()

	var username string
	var result error

	var row = db.QueryRow("select nama_pengguna from pengguna where nama_pengguna = ?", nama_pengguna)

	result = row.Scan(&username)

	if result == sql.ErrNoRows {
		return false
	} else {
		return true
	}

}

func (mdl_daftar Mdl_daftar) Insert_pengguna(ty_daftar Ty_daftar) {

	db, _ := connect()
	defer db.Close()

	db.Exec("insert into pengguna values (?, ?, ?, ?, ?)", ty_daftar.Nama_pengguna, ty_daftar.Nama, ty_daftar.Kata_sandi, ty_daftar.Surel, ty_daftar.Telepon)
}
