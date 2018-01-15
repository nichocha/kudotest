package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"strconv"
	. "website/controller"
	. "website/types"
)

var (
	content string
	page    Page

	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key           = []byte("super-secret-key")
	session_store = sessions.NewCookieStore(key)
)

func handle_selamat_datang(res http.ResponseWriter, req *http.Request) {
	session, _ := session_store.Get(req, "cookie-name")

	content = "view/selamat_datang.html"
	page.Page_title = "Selamat Datang"
	page.Is_login = session.Values["is_login"].(bool)

	if page.Is_login == true {
		http.Redirect(res, req, "/beranda", http.StatusSeeOther)
	}

	tmpl := template.Must(template.ParseFiles(content, "view/header.html", "view/footer.html"))

	tmpl.Execute(res, page)
}

func handle_daftar(res http.ResponseWriter, req *http.Request) {

	var ty_daftar Ty_daftar
	var ctr_daftar Ctr_daftar

	session, _ := session_store.Get(req, "cookie-name")

	content = "view/daftar.html"
	ty_daftar.Page_title = "Daftar"
	ty_daftar.Is_login = session.Values["is_login"].(bool)

	if ty_daftar.Is_login == true {
		http.Redirect(res, req, "/beranda", http.StatusSeeOther)
	}

	if req.Method == http.MethodPost {

		ty_daftar = ctr_daftar.Validasi_daftar(req)

		if ty_daftar.Error_message == "" {
			ctr_daftar.Insert_pengguna(ty_daftar)
			http.Redirect(res, req, "/masuk", http.StatusSeeOther)
		}

	}

	tmpl := template.Must(template.ParseFiles(content, "view/header.html", "view/footer.html"))

	tmpl.Execute(res, ty_daftar)
}

func handle_masuk(res http.ResponseWriter, req *http.Request) {

	var ty_masuk Ty_masuk
	var ctr_masuk Ctr_masuk

	session, _ := session_store.Get(req, "cookie-name")

	content = "view/masuk.html"
	ty_masuk.Page_title = "Masuk"
	ty_masuk.Is_login = session.Values["is_login"].(bool)

	if ty_masuk.Is_login == true {
		http.Redirect(res, req, "/beranda", http.StatusSeeOther)
	}

	if req.Method == http.MethodPost {

		ty_masuk = ctr_masuk.Validasi_masuk(req)

		if ty_masuk.Error_message == "" {

			session, _ := session_store.Get(req, "cookie-name")

			// Set user as authenticated
			session.Values["is_login"] = true
			session.Values["nama_pengguna"] = ty_masuk.Nama_pengguna
			session.Save(req, res)

			http.Redirect(res, req, "/beranda", http.StatusSeeOther)
		}

	}

	tmpl := template.Must(template.ParseFiles(content, "view/header.html", "view/footer.html"))

	tmpl.Execute(res, ty_masuk)
}

func handle_beranda(res http.ResponseWriter, req *http.Request) {

	var ty_beranda Ty_beranda

	session, _ := session_store.Get(req, "cookie-name")

	content = "view/beranda.html"
	ty_beranda.Page_title = "Beranda"
	ty_beranda.Is_login = session.Values["is_login"].(bool)
	ty_beranda.Nama_pengguna = session.Values["nama_pengguna"].(string)

	if ty_beranda.Is_login == false {
		http.Redirect(res, req, "/masuk", http.StatusSeeOther)
	}

	tmpl := template.Must(template.ParseFiles(content, "view/header.html", "view/footer.html"))

	tmpl.Execute(res, ty_beranda)
}

func handle_keluar(res http.ResponseWriter, req *http.Request) {
	session, _ := session_store.Get(req, "cookie-name")

	// Revoke users authentication
	session.Values["is_login"] = false
	session.Values["nama_pengguna"] = ""

	session.Save(req, res)

	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func handle_profil(res http.ResponseWriter, req *http.Request) {
	var ty_profil Ty_profil
	var ctr_profil Ctr_profil

	session, _ := session_store.Get(req, "cookie-name")

	content = "view/profil.html"
	ty_profil.Page_title = "Profil"
	ty_profil.Is_login = session.Values["is_login"].(bool)
	ty_profil.Nama_pengguna = session.Values["nama_pengguna"].(string)

	if ty_profil.Is_login == false {
		http.Redirect(res, req, "/masuk", http.StatusSeeOther)
	}

	if req.Method == http.MethodPost {

		ty_profil = ctr_profil.Validasi_profil(ty_profil.Nama_pengguna, req)

		if ty_profil.Error_message == "" {
			ctr_profil.Edit_profil(ty_profil)

			ty_profil.Success_message = "Data updated"
		}
	} else {
		ty_profil = ctr_profil.Get_profil(ty_profil.Nama_pengguna)
	}

	ty_profil.Is_login = true

	tmpl := template.Must(template.ParseFiles(content, "view/header.html", "view/footer.html"))

	tmpl.Execute(res, ty_profil)
}

func handle_grup(res http.ResponseWriter, req *http.Request) {
	var ty_grup Ty_grup
	var ctr_grup Ctr_grup

	session, _ := session_store.Get(req, "cookie-name")

	content = "view/grup.html"
	ty_grup.Page_title = "Grup"
	ty_grup.Is_login = session.Values["is_login"].(bool)
	ty_grup.Nama_pengguna = session.Values["nama_pengguna"].(string)

	if ty_grup.Is_login == false {
		http.Redirect(res, req, "/masuk", http.StatusSeeOther)
	}

	ty_grup.Grups = ctr_grup.Get_all_grup()

	tmpl := template.Must(template.ParseFiles(content, "view/header.html", "view/footer.html"))

	tmpl.Execute(res, ty_grup)
}

func handle_grup_insert(res http.ResponseWriter, req *http.Request) {
	var ty_grup Ty_grup
	var ctr_grup Ctr_grup

	session, _ := session_store.Get(req, "cookie-name")

	content = "view/grup_form.html"
	ty_grup.Page_title = "Tambah Grup"
	ty_grup.Is_login = session.Values["is_login"].(bool)
	ty_grup.Nama_pengguna = session.Values["nama_pengguna"].(string)

	if ty_grup.Is_login == false {
		http.Redirect(res, req, "/masuk", http.StatusSeeOther)
	}

	if req.Method == http.MethodPost {

		ty_grup = ctr_grup.Validasi_grup(0, req)

		if ty_grup.Error_message == "" {

			ctr_grup.Insert_grup(ty_grup)
			http.Redirect(res, req, "/grup", http.StatusSeeOther)
		}

	}
	ty_grup.Is_login = true
	tmpl := template.Must(template.ParseFiles(content, "view/header.html", "view/footer.html"))

	tmpl.Execute(res, ty_grup)
}

func handle_grup_ubah(res http.ResponseWriter, req *http.Request) {
	var ty_grup Ty_grup
	var ty_grup_data Ty_grup_data
	var ctr_grup Ctr_grup

	session, _ := session_store.Get(req, "cookie-name")

	content = "view/grup_form.html"
	ty_grup.Page_title = "Ubah Grup"
	ty_grup.Is_login = session.Values["is_login"].(bool)
	ty_grup.Nama_pengguna = session.Values["nama_pengguna"].(string)

	if ty_grup.Is_login == false {
		http.Redirect(res, req, "/masuk", http.StatusSeeOther)
	}
	vars := mux.Vars(req)
	id_grup := vars["id_grup"]

	var id_grup_int, _ = strconv.Atoi(id_grup)

	if req.Method == http.MethodPost {

		ty_grup = ctr_grup.Validasi_grup(id_grup_int, req)

		if ty_grup.Error_message == "" {

			ctr_grup.Update_grup(ty_grup)
			http.Redirect(res, req, "/grup", http.StatusSeeOther)
		}

	} else {

		ty_grup_data = ctr_grup.Get_grup(id_grup)

		ty_grup.Id_grup = ty_grup_data.Id_grup
		ty_grup.Nama_grup = ty_grup_data.Nama_grup
		ty_grup.Deskripsi = ty_grup_data.Deskripsi
	}

	ty_grup.Is_login = true
	tmpl := template.Must(template.ParseFiles(content, "view/header.html", "view/footer.html"))

	tmpl.Execute(res, ty_grup)
}

func handle_grup_hapus(res http.ResponseWriter, req *http.Request) {
	var ty_grup Ty_grup
	var ctr_grup Ctr_grup

	session, _ := session_store.Get(req, "cookie-name")

	content = "view/grup_form.html"
	ty_grup.Page_title = "Tambah Grup"
	ty_grup.Is_login = session.Values["is_login"].(bool)
	ty_grup.Nama_pengguna = session.Values["nama_pengguna"].(string)

	if ty_grup.Is_login == false {
		http.Redirect(res, req, "/masuk", http.StatusSeeOther)
	}

	vars := mux.Vars(req)
	id_grup := vars["id_grup"]

	ctr_grup.Delete_grup(id_grup)
	http.Redirect(res, req, "/grup", http.StatusSeeOther)
}

func handle_akses(res http.ResponseWriter, req *http.Request) {
	var ty_akses Ty_akses
	var ctr_akses Ctr_akses

	session, _ := session_store.Get(req, "cookie-name")

	content = "view/akses.html"
	ty_akses.Page_title = "Akses"
	ty_akses.Is_login = session.Values["is_login"].(bool)
	ty_akses.Nama_pengguna = session.Values["nama_pengguna"].(string)

	if ty_akses.Is_login == false {
		http.Redirect(res, req, "/masuk", http.StatusSeeOther)
	}

	ty_akses.Aksess = ctr_akses.Get_all_akses()

	tmpl := template.Must(template.ParseFiles(content, "view/header.html", "view/footer.html"))

	tmpl.Execute(res, ty_akses)
}

func handle_akses_insert(res http.ResponseWriter, req *http.Request) {
	var ty_akses Ty_akses
	var ctr_akses Ctr_akses

	session, _ := session_store.Get(req, "cookie-name")

	content = "view/akses_form.html"
	ty_akses.Page_title = "Tambah Akses"
	ty_akses.Is_login = session.Values["is_login"].(bool)
	ty_akses.Nama_pengguna = session.Values["nama_pengguna"].(string)

	if ty_akses.Is_login == false {
		http.Redirect(res, req, "/masuk", http.StatusSeeOther)
	}

	if req.Method == http.MethodPost {

		ty_akses = ctr_akses.Validasi_akses(0, req)

		if ty_akses.Error_message == "" {

			ctr_akses.Insert_akses(ty_akses)
			http.Redirect(res, req, "/akses", http.StatusSeeOther)
		}

	}
	ty_akses.Is_login = true
	tmpl := template.Must(template.ParseFiles(content, "view/header.html", "view/footer.html"))

	tmpl.Execute(res, ty_akses)
}

func handle_akses_ubah(res http.ResponseWriter, req *http.Request) {
	var ty_akses Ty_akses
	var ty_akses_data Ty_akses_data
	var ctr_akses Ctr_akses

	session, _ := session_store.Get(req, "cookie-name")

	content = "view/akses_form.html"
	ty_akses.Page_title = "Ubah Akses"
	ty_akses.Is_login = session.Values["is_login"].(bool)
	ty_akses.Nama_pengguna = session.Values["nama_pengguna"].(string)

	if ty_akses.Is_login == false {
		http.Redirect(res, req, "/masuk", http.StatusSeeOther)
	}
	vars := mux.Vars(req)
	id_akses := vars["id_akses"]

	var id_akses_int, _ = strconv.Atoi(id_akses)

	if req.Method == http.MethodPost {

		ty_akses = ctr_akses.Validasi_akses(id_akses_int, req)

		if ty_akses.Error_message == "" {

			ctr_akses.Update_akses(ty_akses)
			http.Redirect(res, req, "/akses", http.StatusSeeOther)
		}

	} else {

		ty_akses_data = ctr_akses.Get_akses(id_akses)

		ty_akses.Id_akses = ty_akses_data.Id_akses
		ty_akses.Nama_akses = ty_akses_data.Nama_akses
		ty_akses.Deskripsi = ty_akses_data.Deskripsi
	}

	ty_akses.Is_login = true
	tmpl := template.Must(template.ParseFiles(content, "view/header.html", "view/footer.html"))

	tmpl.Execute(res, ty_akses)
}

func handle_akses_hapus(res http.ResponseWriter, req *http.Request) {
	var ty_akses Ty_akses
	var ctr_akses Ctr_akses

	session, _ := session_store.Get(req, "cookie-name")

	content = "view/akses_form.html"
	ty_akses.Page_title = "Tambah Grup"
	ty_akses.Is_login = session.Values["is_login"].(bool)
	ty_akses.Nama_pengguna = session.Values["nama_pengguna"].(string)

	if ty_akses.Is_login == false {
		http.Redirect(res, req, "/masuk", http.StatusSeeOther)
	}

	vars := mux.Vars(req)
	id_akses := vars["id_akses"]

	ctr_akses.Delete_akses(id_akses)
	http.Redirect(res, req, "/akses", http.StatusSeeOther)
}

func main() {

	route := mux.NewRouter()

	//Asset Prefix Define
	fs := http.FileServer(http.Dir("asset"))
	route.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", fs))

	route.HandleFunc("/", handle_selamat_datang)
	route.HandleFunc("/daftar", handle_daftar)
	route.HandleFunc("/masuk", handle_masuk)
	route.HandleFunc("/beranda", handle_beranda)
	route.HandleFunc("/profil", handle_profil)
	route.HandleFunc("/grup", handle_grup)
	route.HandleFunc("/grup/tambah", handle_grup_insert)
	route.HandleFunc("/grup/ubah/{id_grup}", handle_grup_ubah)
	route.HandleFunc("/grup/hapus/{id_grup}", handle_grup_hapus)
	route.HandleFunc("/akses", handle_akses)
	route.HandleFunc("/akses/tambah", handle_akses_insert)
	route.HandleFunc("/akses/ubah/{id_akses}", handle_akses_ubah)
	route.HandleFunc("/akses/hapus/{id_akses}", handle_akses_hapus)
	route.HandleFunc("/keluar", handle_keluar)

	fmt.Println("starting web server at http://localhost:7000/")

	http.ListenAndServe(":7000", route)

}
