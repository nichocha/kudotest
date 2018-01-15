package controller

import (
	"net/http"
	. "website/model"
	. "website/types"
)

type Ctr_akses struct{}

func (ctr_akses Ctr_akses) Get_all_akses() []Ty_akses_data {

	var ty_akses_data []Ty_akses_data
	var mdl_akses Mdl_akses

	ty_akses_data = mdl_akses.Get_all_akses()

	return ty_akses_data
}

func (ctr_akses Ctr_akses) Get_akses(id_akses string) Ty_akses_data {

	var ty_akses_data Ty_akses_data
	var mdl_akses Mdl_akses

	ty_akses_data = mdl_akses.Get_akses(id_akses)

	return ty_akses_data
}

func (ctr_akses Ctr_akses) Validasi_akses(id_akses int, req *http.Request) Ty_akses {

	var ty_akses Ty_akses

	ty_akses.Id_akses = id_akses
	ty_akses.Nama_akses = req.FormValue("nama_akses")
	ty_akses.Deskripsi = req.FormValue("deskripsi")

	if ty_akses.Nama_akses == "" {
		ty_akses.Error_message = "Mohon mengisi Nama Grup"
	} else if ty_akses.Deskripsi == "" {
		ty_akses.Error_message = "Mohon mengisi Deskripsi"
	}

	return ty_akses
}

func (ctr_akses Ctr_akses) Insert_akses(ty_akses Ty_akses) {

	var mdl_akses Mdl_akses

	mdl_akses.Insert_akses(ty_akses)

}

func (ctr_akses Ctr_akses) Update_akses(ty_akses Ty_akses) {

	var mdl_akses Mdl_akses

	mdl_akses.Update_akses(ty_akses)

}

func (ctr_akses Ctr_akses) Delete_akses(id_akses string) {

	var mdl_akses Mdl_akses

	mdl_akses.Delete_akses(id_akses)

}
