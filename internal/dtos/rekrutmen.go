package dtos

import (
	"time"
)

type CreateRekrutmenRequest struct {
	Kegiatan       string    `json:"kegiatan" binding:"required,oneof=projek riset lomba"`
	Kriteria       string    `json:"Kriteria" binding:"required"`
	TanggalMulai   time.Time `json:"tanggal_mulai" binding:"required"`
	TanggalSelesai time.Time `json:"tanggal_selesai" binding:"required"`
	Fee            float64   `json:"fee"`
	Role           string    `json:"role" binding:"required"`
	ContactPerson  string    `json:"contact_person" binding:"required"`
}

type RekrutmenResponse struct {
	RekrutmenID    string    `json:"rekrutmen_id"`
	UserID         string    `json:"user_id"`
	Kegiatan       string    `json:"kegiatan"`
	Kriteria       string    `json:"Kriteria" binding:"required"`
	TanggalMulai   time.Time `json:"tanggal_mulai"`
	TanggalSelesai time.Time `json:"tanggal_selesai"`
	Fee            float64   `json:"fee"`
	Role           string    `json:"role"`
	ContactPerson  string    `json:"contact_person"`
}

type ApplierRekrutmenResponse struct {
	RekrutmenID    string              `json:"rekrutmen_id"`
	UserID         string              `json:"user_id"`
	Kegiatan       string              `json:"kegiatan"`
	Kriteria       string              `json:"Kriteria" binding:"required"`
	TanggalMulai   time.Time           `json:"tanggal_mulai"`
	TanggalSelesai time.Time           `json:"tanggal_selesai"`
	Fee            float64             `json:"fee"`
	Role           string              `json:"role"`
	ContactPerson  string              `json:"contact_person"`
	Pendaftar      []PendaftarResponse `json:"pendaftar"`
}

type RekrutmenListResponse struct {
	Data       []RekrutmenResponse `json:"data"`
	Page       int                 `json:"page"`
	Limit      int                 `json:"limit"`
	TotalItems int64               `json:"total_items"`
	TotalPages int                 `json:"total_pages"`
}

type ApplyRequest struct {
	AlasanMendaftar string `json:"alasan_mendaftar"`
	CVURL           string `json:"cv_url"`
	PortofolioURL   string `json:"portofolio_url"`
}

type PendaftarResponse struct {
	PendaftarID     string  `json:"pendaftar_id"`
	RekrutmenID     string  `json:"rekrutmen_id"`
	UserID          string  `json:"user_id"`
	AlasanMendaftar *string `json:"alasan_mendaftar"`
	CVURL           string  `json:"cv_url"`
	PortofolioURL   string  `json:"portofolio_url"`
	Status          string  `json:"status"`
	NamaPendaftar   string  `json:"nama_pendaftar"`
}

type TimMemberResponse struct {
	UserID       string `json:"user_id"`
	Nama         string `json:"nama"`
	Jurusan      string `json:"jurusan"`
	MemberKe     int64  `json:"member_ke"`
	NoPengenal   string `json:"no_pengenal"`
	NoTelp       string `json:"no_telp"`
	Spesialisasi []string `json:"spesialisasi"`
}

type GiveMemberRatingRequest struct {
	Rating    int64  `json:"rating" binding:"required,min=1,max=5"`
	Deskripsi string `json:"deskripsi" binding:"required"`
}

type UpdateApplyStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=accepted rejected"`
}

type UploadFileResponse struct {
	URL string `json:"url"`
}
