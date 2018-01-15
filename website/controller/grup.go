package controller

import (
	"net/http"
	. "website/model"
	. "website/types"
)

type Ctr_grup struct{}

func (ctr_grup Ctr_grup) Get_all_grup() []Ty_grup_data {

	var ty_grup_data []Ty_grup_data
	var mdl_grup Mdl_grup

	ty_grup_data = mdl_grup.Get_all_grup()

	return ty_grup_data
}

func (ctr_grup Ctr_grup) Get_grup(id_grup string) Ty_grup_data {

	var ty_grup_data Ty_grup_data
	var mdl_grup Mdl_grup

	ty_grup_data = mdl_grup.Get_grup(id_grup)

	return ty_grup_data
}

func (ctr_grup Ctr_grup) Validasi_grup(id_grup int, req *http.Request) Ty_grup {

	var ty_grup Ty_grup

	ty_grup.Id_grup = id_grup
	ty_grup.Nama_grup = req.FormValue("nama_grup")
	ty_grup.Deskripsi = req.FormValue("deskripsi")

	if ty_grup.Nama_grup == "" {
		ty_grup.Error_message = "Mohon mengisi Nama Grup"
	} else if ty_grup.Deskripsi == "" {
		ty_grup.Error_message = "Mohon mengisi Deskripsi"
	}

	return ty_grup
}

func (ctr_grup Ctr_grup) Insert_grup(ty_grup Ty_grup) {

	var mdl_grup Mdl_grup

	mdl_grup.Insert_grup(ty_grup)

}

func (ctr_grup Ctr_grup) Update_grup(ty_grup Ty_grup) {

	var mdl_grup Mdl_grup

	mdl_grup.Update_grup(ty_grup)

}

func (ctr_grup Ctr_grup) Delete_grup(id_grup string) {

	var mdl_grup Mdl_grup

	mdl_grup.Delete_grup(id_grup)

}
