package models

type Rating struct {
	Id                 int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	UserId 			   int 		 `gorm:"column:user_id" json:"user_id" redis:"user_id"`
	ArticleId 		   int 		 `gorm:"column:article_id" json:"article_id" redis:"article_id"`
	Rating			   int 		 `gorm:"column:rating" json:"rating" redis:"rating"`
}
