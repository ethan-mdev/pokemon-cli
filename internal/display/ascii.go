package display

import (
	"fmt"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

// DisplayImage displays an image as ASCII art
func DisplayImage(url string) error {
	if url == "" {
		return fmt.Errorf("no image URL provided")
	}

	// Convert to ASCII using the library
	flags := aic_package.DefaultFlags()
	flags.Dimensions = []int{60, 30} // Set desired size
	flags.Colored = true             // Enable colors

	// Pass the url
	asciiArt, err := aic_package.Convert(url, flags)
	if err != nil {
		return fmt.Errorf("failed to convert image: %v", err)
	}

	fmt.Printf("%v\n", asciiArt)
	return nil
}
