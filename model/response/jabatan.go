package response

import (
	"materi/helper"
)

type JabatanResponse struct {
	ID          int               `json:"id"`
	NamaJabatan string            `json:"nama_jabatan"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
	DeletedAt   helper.NullString `json:"deleted_at"`
}
