package model

import (
	. "website/types"
)

type Mdl_profil struct{}

func (mdl_profil Mdl_profil) Edit_profil(ty_profil Ty_profil) {

	db, _ := connect()
	defer db.Close()

	db.Exec("update pengguna set nama = ?, surel = ?, telepon = ? where nama_pengguna = ?", ty_profil.Nama, ty_profil.Surel, ty_profil.Telepon, ty_profil.Nama_pengguna)

}

func (mdl_profil Mdl_profil) Get_pengguna(nama_pengguna string) Ty_profil {

	db, _ := connect()
	defer db.Close()

	var ty_profil Ty_profil

	var row = db.QueryRow("select nama, surel, telepon from pengguna where nama_pengguna = ?", nama_pengguna)

	row.Scan(
		&ty_profil.Nama,
		&ty_profil.Surel,
		&ty_profil.Telepon,
	)

	return ty_profil
}
