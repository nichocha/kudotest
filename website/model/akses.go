package model

import (
	. "website/types"
)

type Mdl_akses struct{}

func (mdl_akses Mdl_akses) Get_all_akses() []Ty_akses_data {

	db, _ := connect()
	defer db.Close()

	rows, _ := db.Query("select * from akses")
	defer rows.Close()

	var tt_akses []Ty_akses_data

	for rows.Next() {
		var each = Ty_akses_data{}
		var _ = rows.Scan(
			&each.Id_akses,
			&each.Nama_akses,
			&each.Deskripsi,
		)

		tt_akses = append(tt_akses, each)
	}

	return tt_akses
}

func (mdl_akses Mdl_akses) Get_akses(id_akses string) Ty_akses_data {

	db, _ := connect()
	defer db.Close()

	var ty_akses_data Ty_akses_data

	var row = db.QueryRow("select * from akses where id_akses = ?", id_akses)

	row.Scan(
		&ty_akses_data.Id_akses,
		&ty_akses_data.Nama_akses,
		&ty_akses_data.Deskripsi,
	)

	return ty_akses_data

}

func (mdl_akses Mdl_akses) Insert_akses(ty_akses Ty_akses) {

	db, _ := connect()
	defer db.Close()

	db.Exec("insert into akses (nama_akses, deskripsi) values (?, ?)", ty_akses.Nama_akses, ty_akses.Deskripsi)

}

func (mdl_akses Mdl_akses) Update_akses(ty_akses Ty_akses) {

	db, _ := connect()
	defer db.Close()

	db.Exec("update akses set nama_akses = ?, deskripsi = ? where id_akses = ?", ty_akses.Nama_akses, ty_akses.Deskripsi, ty_akses.Id_akses)

}

func (mdl_akses Mdl_akses) Delete_akses(id_akses string) {

	db, _ := connect()
	defer db.Close()

	db.Exec("delete from akses where id_akses = ?", id_akses)

}
