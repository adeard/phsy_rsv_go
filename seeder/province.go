package seeder

import (
	"phsy_rsv_go/domain"

	"gorm.io/gorm"
)

func InsertProvince(Db *gorm.DB) {
	var total int64
	var ut domain.Province
	Db.Model(&ut).Count(&total)
	if total == 0 {
		var data = []domain.Province{
			{
				Name: "Nanggroe Aceh Darussalam", IsActive: true,
			},
			{
				Name: "Sumatera Utara", IsActive: true,
			},
			{
				Name: "Sumatera Selatan", IsActive: true,
			},
			{
				Name: "Sumatera Barat", IsActive: true,
			},
			{
				Name: "Bengkulu", IsActive: true,
			},
			{
				Name: "Riau", IsActive: true,
			},
			{
				Name: "Kepulauan Riau", IsActive: true,
			},
			{
				Name: "Jambi", IsActive: true,
			},
			{
				Name: "Lampung", IsActive: true,
			},
			{
				Name: "Bangka Belitung", IsActive: true,
			},
			{
				Name: "Kalimantan Barat", IsActive: true,
			},
			{
				Name: "Kalimantan Timur", IsActive: true,
			},
			{
				Name: "Kalimantan Selatan", IsActive: true,
			},
			{
				Name: "Kalimantan Tengah", IsActive: true,
			},
			{
				Name: "Kalimantan Utara", IsActive: true,
			},
			{
				Name: "Banten", IsActive: true,
			},
			{
				Name: "DKI Jakarta", IsActive: true,
			},
			{
				Name: "Jawa Barat", IsActive: true,
			},
			{
				Name: "Jawa Timur", IsActive: true,
			},
			{
				Name: "Jawa Tengah", IsActive: true,
			},
			{
				Name: "Daerah Istimewa Yogyakarta", IsActive: true,
			},
			{
				Name: "Bali", IsActive: true,
			},
			{
				Name: "Nusa Tenggara Timur", IsActive: true,
			},
			{
				Name: "Nusa Tenggara Barat", IsActive: true,
			},
			{
				Name: "Gorontalo", IsActive: true,
			},
			{
				Name: "Sulawesi Barat", IsActive: true,
			},
			{
				Name: "Sulawesi Tengah", IsActive: true,
			},
			{
				Name: "Sulawesi Utara", IsActive: true,
			},
			{
				Name: "Sulawesi Tenggara", IsActive: true,
			},
			{
				Name: "Sulawesi Selatan", IsActive: true,
			},
			{
				Name: "Maluku Utara", IsActive: true,
			},
			{
				Name: "Maluku", IsActive: true,
			},
			{
				Name: "Papua Barat", IsActive: true,
			},
			{
				Name: "Papua", IsActive: true,
			},
			{
				Name: "Papua Tengah", IsActive: true,
			},
			{
				Name: "Papua Pegunungan", IsActive: true,
			},
			{
				Name: "Papua Selatan", IsActive: true,
			},
			{
				Name: "Papua Barat Daya", IsActive: true,
			},
		}

		Db.Create(&data)
	}
}
