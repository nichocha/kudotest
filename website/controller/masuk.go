package controller

import (
	"net/http"
	. "website/model"
	. "website/types"
)

type Ctr_masuk struct{}

func (ctr_masuk Ctr_masuk) Validasi_masuk(req *http.Request) Ty_masuk {

	var ty_masuk Ty_masuk
	var mdl_masuk Mdl_masuk

	ty_masuk.Nama_pengguna = req.FormValue("nama_pengguna")
	ty_masuk.Kata_sandi = req.FormValue("kata_sandi")

	if ty_masuk.Nama_pengguna == "" {
		ty_masuk.Error_message = "Mohon mengisi Nama Pengguna"
	} else if ty_masuk.Kata_sandi == "" {
		ty_masuk.Error_message = "Mohon mengisi Kata Sandi"
	} else if mdl_masuk.Get_pengguna(ty_masuk) == false {
		ty_masuk.Error_message = "Kombinasi Nama Pengguna dan Kata Sandi salah"
	}

	return ty_masuk
}
