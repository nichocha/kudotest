package types

type Page struct {
	Page_title    string
	Error_message string
	Is_login      bool
	Nama_pengguna string
}

type Ty_daftar struct {
	Page
	Nama                  string
	Surel                 string
	Telepon               string
	Kata_sandi            string
	Konfirmasi_kata_sandi string
}

type Ty_masuk struct {
	Page
	Kata_sandi string
}

type Ty_beranda struct {
	Page
}

type Ty_profil struct {
	Page
	Nama            string
	Surel           string
	Telepon         string
	Success_message string
}

type Ty_grup_data struct {
	Id_grup   int
	Nama_grup string
	Deskripsi string
}

type Ty_grup struct {
	Page
	Ty_grup_data
	Grups []Ty_grup_data
}

type Ty_akses_data struct {
	Id_akses   int
	Nama_akses string
	Deskripsi  string
}

type Ty_akses struct {
	Page
	Ty_akses_data
	Aksess []Ty_akses_data
}
