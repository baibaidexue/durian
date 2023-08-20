package dao

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ComicMeta struct {
	gorm.Model

	OriginURL string `gorm:"column:origin_url"`
	JMID      string `gorm:"column:jm_id"`

	Name         string    `gorm:"unique_index:unique_idx_comic_name_url;column:name"`
	Overview     string    `gorm:"column:overview"`
	PageCount    int       `gorm:"column:page_total"`
	JMViewCount  int       `gorm:"column:jm_viewcount"`
	JMLiked      int       `gorm:"column:jm_liked"`
	JMUploadTime time.Time `gorm:"column:jm_uploadtime"`
	JMUpdateTime time.Time `gorm:"column:jm_updatetime"`
	JMComments   string    `gorm:"column:jm_comments"`

	PosterURL      string `gorm:"column:poster_url"`
	PosterFilePath string `gorm:"column:poster_filepath"`

	UserLiked int `gorm:"column:user_liked"`
	UserScore int `gorm:"column:user_score"`

	Ep     []ComicEPInfo `gorm:"many2many:comic_ep_link;"`
	Artist []ComicAuthor `gorm:"many2many:comic_author_link;"`
	Tag    []ComicTag    `gorm:"many2many:comic_tag_link;"`
	Actor  []ComicActor  `gorm:"many2many:comic_actor_link;"`
}

type ComicEPInfo struct {
	DownloadUrl string `gorm:"primary_key"`
	FilePath    string
	FileName    string
	Size        uint64

	Healthy int
}

type ComicTag struct {
	Name       string `gorm:"primary_key"`
	SearchName string
	Link       string
}

type ComicAuthor struct {
	Name       string `gorm:"primary_key"`
	SearchName string
	Link       string
}
type ComicActor struct {
	Name       string `gorm:"primary_key"`
	SearchName string
	Link       string
}
