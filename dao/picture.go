package dao

import "github.com/jinzhu/gorm"

type Picture struct {
	gorm.Model

	SamplePath string
	OriginURL  string
	Name       string
	ImageURL   string `gorm:"unique_index:unique_idx_picture_url"`
	FilePath   string
	FileName   string
	Size       uint64

	CopyRight []PictureCopyRight `gorm:"many2many:picture_copyright_link;"`
	Artist    []PictureArtist    `gorm:"many2many:picture_artist_link;"`
	Tags      []PictureTag       `gorm:"many2many:picture_tag_link;"`
	Character []PictureCharacter `gorm:"many2many:picture_character_link;"`
	Metadata  []PictureMetaData  `gorm:"many2many:picture_metadata_link;"`
}

// gorm:"many2many:picture_artist_link;uniqueIndex:unique_idx_picture_artist"

type PictureCopyRight struct {
	// gorm.Model
	CopyRightName string `gorm:"primary_key"`
	SearchName    string
	Link          string
}

type PictureArtist struct {
	// gorm.Model
	ArtistName string `gorm:"primary_key"`
	SearchName string
	Link       string
}

type PictureMetaData struct {
	// gorm.Model
	MetaValue  string `gorm:"primary_key"`
	SearchName string
	Link       string
}

type PictureTag struct {
	// gorm.Model
	TagName    string `gorm:"primary_key"`
	SearchName string
	Link       string
}

type PictureCharacter struct {
	// gorm.Model
	CharacterName string `gorm:"primary_key"`
	SearchName    string
	Link          string
}
