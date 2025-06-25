package gameengine

import "github.com/mnmonherdene1234/uns-game/gameengine/render"

type Image struct {
	ID   uint32 // OpenGL texture ID
	Path string // Path to the image file
	Name string // Name of the image
}

type AssetsManager struct {
	Images []*Image
}

func NewAssetsManager() *AssetsManager {
	return &AssetsManager{
		Images: make([]*Image, 0),
	}
}

func (am *AssetsManager) LoadImages() error {
	for _, img := range am.Images {
		textureID, err := render.LoadTexture(img.Path)
		if err != nil {
			return err // Return error if texture loading fails
		}
		img.ID = textureID // Assign the loaded texture ID to the image
	}

	return nil // Return nil if all images are loaded successfully
}

func (am *AssetsManager) AddImage(path, name string) {
	am.Images = append(am.Images, &Image{
		Path: path,
		Name: name,
	})
}

func (am *AssetsManager) GetImageByName(name string) (*Image, bool) {
	for _, img := range am.Images {
		if img.Name == name {
			return img, true // Return the image if found
		}
	}
	return nil, false // Return nil if not found
}
