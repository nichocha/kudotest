package controller

import (
	"net/http"
	. "website/model"
	. "website/types"
)

type Ctr_profil struct{}

func (ctr_profil Ctr_profil) Validasi_profil(nama_pengguna string, req *http.Request) Ty_profil {

	var ty_profil Ty_profil

	ty_profil.Nama_pengguna = nama_pengguna
	ty_profil.Nama = req.FormValue("nama")
	ty_profil.Surel = req.FormValue("surel")
	ty_profil.Telepon = req.FormValue("telepon")

	if ty_profil.Nama == "" {
		ty_profil.Error_message = "Mohon mengisi Nama"
	} else if ty_profil.Surel == "" {
		ty_profil.Error_message = "Mohon mengisi Surel"
	} else if ty_profil.Telepon == "" {
		ty_profil.Error_message = "Mohon mengisi Telepon"
	}

	return ty_profil
}

func (ctr_profil Ctr_profil) Get_profil(nama_pengguna string) Ty_profil {

	var mdl_profil Mdl_profil

	return mdl_profil.Get_pengguna(nama_pengguna)

}

func (ctr_profil Ctr_profil) Edit_profil(ty_profil Ty_profil) {

	var mdl_profil Mdl_profil

	mdl_profil.Edit_profil(ty_profil)
}
