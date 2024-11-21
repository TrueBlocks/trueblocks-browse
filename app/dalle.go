package app

import (
	"path/filepath"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

var ImageStoragePath = filepath.Join("./frontend/src/assets/dalle_images")

func (a *App) LoadDalleImage(address base.Address) (bool, error) {
	return false, nil
	// imageName := address.Hex() + ".png"
	// imagePath := filepath.Join(ImageStoragePath, imageName)
	// imageFolder := filepath.Dir(imagePath)
	// if err := file.EstablishFolder(imageFolder); err != nil {
	// 	return false, fmt.Errorf("failed to establish folder: %w", err)
	// }

	// if _, err := os.Stat(imagePath); err == nil {
	// 	return true, nil
	// } else if !os.IsNotExist(err) {
	// 	return false, err
	// }
	// err := downloadImage(imageName, imagePath)
	// return false, err
}

// func downloadImage(imageName, imagePath string) error {
// 	url := fmt.Sprintf("https://your-remote-url.com/images/%s", imageName)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	out, err := os.Create(imagePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer out.Close()
// 	_, err = io.Copy(out, resp.Body)
// 	return err
// }
