package response

import "time"

type JabatanResponse struct {
	ID          string    `json:"id"`
	NamaJabatan string    `json:"nama_jabatan"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
	DeleteAt    time.Time `json:"delete_at"`
}
