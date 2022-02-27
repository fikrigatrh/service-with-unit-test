package models

import "gorm.io/gorm"

type SocialMediaDetail struct {
	Instagram string `json:"instagram"`
	Facebook  string `json:"facebook"`
	Twitter   string `json:"twitter"`
}

type AboutUsRequest struct {
	Profil            string                    `json:"profil" validate:"required"`
	Visi              string                    `json:"visi" validate:"required"`
	Misi              []MisiDetail              `json:"misi" validate:"required"`
	Motto             string                    `json:"motto" validate:"required"`
	PerusahaanRekanan []PerusahaanRekananDetail `json:"perusahaan_rekanan"`
	Office            string                    `json:"office" validate:"required"`
	Warehouse         string                    `json:"warehouse"`
	Email             string                    `json:"email" validate:"required"`
	NoTelp            string                    `json:"no_telp" validate:"required"`
	SocialMedia       []SocialMediaDetail       `json:"social_media" validate:"required"`
}

type Footer struct {
}

type MisiDetail struct {
	Item string `json:"item"`
}

type Header struct {
}

type Booking struct {
	gorm.Model
	NamaPengirim             string `json:"nama_pengirim"`
	Email                    string `json:"email" validate:"email"`
	NoTelpPengirim           string `json:"no_telp_pengirim" validate:"numeric"`
	KotaPenjemputan          string `json:"kota_penjemputan"`
	AlamatLengkapPenjemputan string `json:"alamat_lengkap_penjemputan"`
	NamaPenerima             string `json:"nama_penerima"`
	NoTelpPenerima           string `json:"no_telp_penerima" validate:"numeric"`
	KotaPenerima             string `json:"kota_penerima"`
	AlamatPenerima           string `json:"alamat_penerima"`
	JenisBarang              string `json:"jenis_barang"`
	PerkiraanBerat           string `json:"perkiraan_berat" validate:"numeric"`
	JumlahKoli               string `json:"jumlah_koli" validate:"numeric"`
	Kubikasi                 string `json:"kubikasi"`
	Keterangan               string `json:"keterangan"`
}

type AboutUsDb struct {
	gorm.Model
	Profil            string `json:"profil"`
	Visi              string `json:"visi"`
	Misi              string `json:"misi" gorm:"type:JSONB NULL DEFAULT '{}'::JSONB"`
	Motto             string `json:"motto"`
	PerusahaanRekanan string `json:"perusahaan_rekanan" gorm:"type:JSONB NULL DEFAULT '{}'::JSONB"`
	SocialMedia       string `json:"social_media" gorm:"type:JSONB NULL DEFAULT '{}'::JSONB"`
	Email             string `json:"email"`
	NoTelp            string `json:"no_telp"`
}

type PerusahaanRekananDetail struct {
	NamaPerusahaan string `json:"nama_perusahaan"`
}
