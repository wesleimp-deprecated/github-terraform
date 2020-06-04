package output

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

// Save the file in the specified name and path
func Save(dest, name, content string) error {
	var filename = fmt.Sprintf("%s/%s.tf", dest, name)

	if _, err := os.Stat(dest); os.IsNotExist(err) {
		log.Printf("Creating %s folfer", dest)
		os.MkdirAll(dest, 0700)
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	color.Green("%s saved successfully.\n", filename)

	return nil
}
