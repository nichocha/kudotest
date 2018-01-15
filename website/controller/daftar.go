package controller

import (
	"net/http"
	. "website/model"
	. "website/types"
)

type Ctr_daftar struct{}

func (ctr_daftar Ctr_daftar) Validasi_daftar(req *http.Request) Ty_daftar {

	var ty_daftar Ty_daftar
	var mdl_daftar Mdl_daftar

	ty_daftar.Nama_pengguna = req.FormValue("nama_pengguna")
	ty_daftar.Nama = req.FormValue("nama")
	ty_daftar.Surel = req.FormValue("surel")
	ty_daftar.Telepon = req.FormValue("telepon")
	ty_daftar.Kata_sandi = req.FormValue("kata_sandi")
	ty_daftar.Konfirmasi_kata_sandi = req.FormValue("konfirmasi_kata_sandi")

	if ty_daftar.Nama_pengguna == "" {
		ty_daftar.Error_message = "Mohon mengisi Nama Pengguna"
	} else if ty_daftar.Nama == "" {
		ty_daftar.Error_message = "Mohon mengisi Nama"
	} else if ty_daftar.Surel == "" {
		ty_daftar.Error_message = "Mohon mengisi Surel"
	} else if ty_daftar.Telepon == "" {
		ty_daftar.Error_message = "Mohon mengisi Telepon"
	} else if ty_daftar.Kata_sandi == "" {
		ty_daftar.Error_message = "Mohon mengisi Kata Sandi"
	} else if ty_daftar.Konfirmasi_kata_sandi == "" {
		ty_daftar.Error_message = "Mohon mengisi Konfirmasi Kata Sandi"
	} else if ty_daftar.Kata_sandi != ty_daftar.Konfirmasi_kata_sandi {
		ty_daftar.Error_message = "Kata Sandi dan Konfirmasi Kata Sandi berbeda"
	} else if mdl_daftar.Get_pengguna(ty_daftar.Nama_pengguna) == true {
		ty_daftar.Error_message = "Nama Pengguna sudah terdaftar"
	}

	return ty_daftar
}

func (ctr_daftar Ctr_daftar) Insert_pengguna(ty_daftar Ty_daftar) {

	var mdl_daftar Mdl_daftar

	mdl_daftar.Insert_pengguna(ty_daftar)
}
