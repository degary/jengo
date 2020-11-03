package initpro

import (
	"fmt"
	"os"
)

func CopyFile(dest, content string) error {
	file, err := os.OpenFile(dest, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
