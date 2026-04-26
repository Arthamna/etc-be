package dtos

import (
	"etc-backend/internal/models"
	"time"
)

type CreateRekrutmenRequest struct {
	Kegiatan       string    `json:"kegiatan" binding:"required,oneof=projek riset lomba"`
	TanggalMulai   time.Time `json:"tanggal_mulai" binding:"required"`
	TanggalSelesai time.Time `json:"tanggal_selesai" binding:"required"`
	Fee            float64   `json:"fee"`
	Role           string    `json:"role" binding:"required"`
	ContactPerson  string    `json:"contact_person" binding:"required"`
}

type RekrutmenResponse struct {
	RekrutmenID    string    `json:"rekrutmen_id"`
	UserID         string    `json:"user_id"`
	KegiatanID     string    `json:"kegiatan_id"`
	Kegiatan       string    `json:"kegiatan"`
	TanggalMulai   time.Time `json:"tanggal_mulai"`
	TanggalSelesai time.Time `json:"tanggal_selesai"`
	Fee            float64   `json:"fee"`
	Role           string    `json:"role"`
	ContactPerson  string    `json:"contact_person"`
	PembuatNama    string    `json:"pembuat_nama"`
}

type RekrutmenListResponse struct {
	Data       []RekrutmenResponse `json:"data"`
	Page       int                 `json:"page"`
	Limit      int                 `json:"limit"`
	TotalItems int64               `json:"total_items"`
	TotalPages int                 `json:"total_pages"`
}

func ToRekrutmenResponse(r *models.Rekrutmen) RekrutmenResponse {
	return RekrutmenResponse{
		RekrutmenID:    r.RekrutmenID,
		UserID:         r.UserID,
		KegiatanID:     r.KegiatanID,
		Kegiatan:       r.Kegiatan,
		TanggalMulai:   r.TanggalMulai,
		TanggalSelesai: r.TanggalSelesai,
		Fee:            r.Fee,
		Role:           r.Role,
		ContactPerson:  r.ContactPerson,
		PembuatNama:    r.User.Nama,
	}
}
