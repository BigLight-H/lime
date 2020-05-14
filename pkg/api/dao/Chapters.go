package dao

import (
	"github.com/jinzhu/gorm"
	"lime/pkg/api/dto"
	"lime/pkg/api/model"
	"lime/pkg/common/db"
)

type ChaptersDao struct {
}

func (c ChaptersDao) Get(id int) model.Chapters {
	var Chapters model.Chapters
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Chapters)
	return Chapters
}

func (c ChaptersDao) GetAll() []model.Chapters {
	var Chapters []model.Chapters
	db := db.GetGormDB()
	db.Model(&model.Chapters{}).Find(&Chapters)
	return Chapters
}

func (c ChaptersDao) List(listDto dto.ChaptersListDto) ([]model.Chapters, int64) {
	var Chapters []model.Chapters
	var total int64
	//ChannelId, err := strconv.Atoi(listDto.ChannelId)
	db := db.GetGormDB()
	//if err == nil && ChannelId == 1 {
	//	db = db.Where("created_at >=?", listDto.Start_time)
	//	db = db.Where("created_at <=?", listDto.End_time)
	//}
	db.Offset(listDto.Skip).Limit(listDto.Limit).Find(&Chapters)
	db.Model(&model.Chapters{}).Count(&total)
	return Chapters, total
}

func (c ChaptersDao) Create(Chapters *model.Chapters) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Chapters)
}

// Update
func (c ChaptersDao) Update(Chapters *model.Chapters, ups map[string]interface{}) *gorm.DB {
	db := db.GetGormDB()
	return db.Model(Chapters).Update(ups)
}

func (c ChaptersDao) Delete(Chapters *model.Chapters) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(Chapters)
}