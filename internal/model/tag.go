package model

import (
	"github.com/i0Ek3/blogie/pkg/app"
	"github.com/jinzhu/gorm"
)

// Tag indicates the field of tag table
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// TagSwagger defines a Swagger object
// used for Swagger interface document display
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// TableName returns the name of tag table
func (t Tag) TableName() string {
	return "blogie_tag"
}

func (t Tag) Get(db *gorm.DB) (Tag, error) {
	var tag Tag
	err := db.Where("id = ? AND is_del = ? AND state = ?", t.ID, 0, t.State).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}

	return tag, nil
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag) ListByIDs(db *gorm.DB, ids []uint32) ([]*Tag, error) {
	var tags []*Tag
	db = db.Where("state = ? AND is_del = ?", t.State, 0)
	err := db.Where("id IN (?)", ids).Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values any) error {
	if err := db.Model(t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	err := db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error

	return err
}

func CleanAllTag(db *gorm.DB) bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})

	return true
}
