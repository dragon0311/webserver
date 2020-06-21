package wallpaper

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func EncodePic(picBuffer string, picName string) error {
	picfile, _ := base64.StdEncoding.DecodeString(picBuffer)
	picUrl := "D:/Code/GoPath/src/webserver/data/" + picName + ".jpg"
	if err := ioutil.WriteFile(picUrl, picfile, 0666); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
