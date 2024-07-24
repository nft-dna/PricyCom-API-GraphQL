package types

import (
	"fmt"
	"net/http"
	"strings"

	svg "github.com/h2non/go-is-svg"

	"crypto/sha256"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Image represents image of NFT downloaded from specified URI
type Image struct {
	Data []byte    `bson:"data"`
	Type ImageType `bson:"type"`
}

type ImageType int8

const (
	ImageTypeUnknown ImageType = iota
	ImageTypeSvg
	ImageTypeGif
	ImageTypeJpeg
	ImageTypePng
	ImageTypeWebp
	ImageTypeMp4
)

func (i ImageType) Mimetype() string {
	switch i {
	case ImageTypeSvg:
		return "image/svg+xml"
	case ImageTypeGif:
		return "image/gif"
	case ImageTypeJpeg:
		return "image/jpeg"
	case ImageTypePng:
		return "image/png"
	case ImageTypeWebp:
		return "image/webp"
	case ImageTypeMp4:
		return "video/mp4"
	}
	return ""
}

func (i ImageType) Extension() string {
	switch i {
	case ImageTypeSvg:
		return ".svg"
	case ImageTypeGif:
		return ".gif"
	case ImageTypeJpeg:
		return ".jpg"
	case ImageTypePng:
		return ".png"
	case ImageTypeWebp:
		return ".webp"
	case ImageTypeMp4:
		return ".mp4"
	}
	return ""
}

func ImageTypeFromMimetype(data []byte) (ImageType, error) {
	mimetype := http.DetectContentType(data)
	switch mimetype {
	case "image/svg+xml":
		return ImageTypeSvg, nil
	case "image/gif":
		return ImageTypeGif, nil
	case "image/jpeg":
		return ImageTypeJpeg, nil
	case "image/png":
		return ImageTypePng, nil
	case "image/webp":
		return ImageTypeWebp, nil
	case "video/mp4":
		return ImageTypeMp4, nil
	}
	if strings.HasPrefix(mimetype, "text/xml") || strings.HasPrefix(mimetype, "text/plain") {
		if svg.Is(data) {
			return ImageTypeSvg, nil
		}
	}
	return ImageTypeUnknown, fmt.Errorf("unrecognized image type %s", mimetype)
}

func ImageTypeFromExtension(uri string) (mimetype ImageType) {
	uri = strings.ToLower(uri)
	if strings.HasSuffix(uri, ".svg") {
		return ImageTypeSvg
	}
	if strings.HasSuffix(uri, ".gif") {
		return ImageTypeGif
	}
	if strings.HasSuffix(uri, ".jpg") || strings.HasSuffix(uri, ".jpeg") {
		return ImageTypeJpeg
	}
	if strings.HasSuffix(uri, ".png") {
		return ImageTypePng
	}
	if strings.HasSuffix(uri, ".webp") {
		return ImageTypeWebp
	}
	if strings.HasSuffix(uri, ".mp4") {
		return ImageTypeMp4
	}
	return ImageTypeUnknown
}

// MM
// ---------------------------------------------------------------
func NewMimeImage(content []byte) *Image {
	t, _ := ImageTypeFromMimetype(content)
	return &Image{
		Data: content,
		Type: t,
	}
}

func NewTypedImage(content []byte, imgtype ImageType) *Image {
	return &Image{
		Data: content,
		Type: imgtype,
	}
}

func NewNamedImage(content []byte, name string) *Image {
	return &Image{
		Data: content,
		Type: ImageTypeFromExtension(name),
	}
}

func ImageIDFromContent(content []byte) primitive.ObjectID {
	hash := sha256.New()
	hash.Write(content)

	var id [12]byte
	copy(id[:], hash.Sum(nil))
	return id
}

func ImageIDFromHexString(hexid string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(hexid)
}

func (t *Image) ID() primitive.ObjectID {
	return ImageIDFromContent(t.Data)
}

// ---------------------------------------------------------------
