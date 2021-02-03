package model

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"createdon"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"delete_on"`
	IsDel      uint8  `json:"is_del"`
}
