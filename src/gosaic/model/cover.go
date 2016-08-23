package model

type Cover struct {
	Id       int64  `db:"id"`
	AspectId int64  `db:"aspect_id"`
	Type     string `db:"type"`
	Width    int    `db:"width"`
	Height   int    `db:"height"`
	Num      int    `db:"num"`
}
