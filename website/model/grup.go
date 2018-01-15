package model

import (
	. "website/types"
)

type Mdl_grup struct{}

func (mdl_grup Mdl_grup) Get_all_grup() []Ty_grup_data {

	db, _ := connect()
	defer db.Close()

	rows, _ := db.Query("select * from grup")
	defer rows.Close()

	var tt_grup []Ty_grup_data

	for rows.Next() {
		var each = Ty_grup_data{}
		var _ = rows.Scan(
			&each.Id_grup,
			&each.Nama_grup,
			&each.Deskripsi,
		)

		tt_grup = append(tt_grup, each)
	}

	return tt_grup
}

func (mdl_grup Mdl_grup) Get_grup(id_grup string) Ty_grup_data {

	db, _ := connect()
	defer db.Close()

	var ty_grup_data Ty_grup_data

	var row = db.QueryRow("select * from grup where id_grup = ?", id_grup)

	row.Scan(
		&ty_grup_data.Id_grup,
		&ty_grup_data.Nama_grup,
		&ty_grup_data.Deskripsi,
	)

	return ty_grup_data

}

func (mdl_grup Mdl_grup) Insert_grup(ty_grup Ty_grup) {

	db, _ := connect()
	defer db.Close()

	db.Exec("insert into grup (nama_grup, deskripsi) values (?, ?)", ty_grup.Nama_grup, ty_grup.Deskripsi)

}

func (mdl_grup Mdl_grup) Update_grup(ty_grup Ty_grup) {

	db, _ := connect()
	defer db.Close()

	db.Exec("update grup set nama_grup = ?, deskripsi = ? where id_grup = ?", ty_grup.Nama_grup, ty_grup.Deskripsi, ty_grup.Id_grup)

}

func (mdl_grup Mdl_grup) Delete_grup(id_grup string) {

	db, _ := connect()
	defer db.Close()

	db.Exec("delete from grup where id_grup = ?", id_grup)

}
