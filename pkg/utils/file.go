package utils

import (
	"fmt"
	"io"
	"log"
	"ncd_operators/pkg/raven"
	s "ncd_operators/pkg/settings"
	"net/http"
	"os"
	"strings"
)

func FileSave(r *http.Request, fileType, passportSerial string) []string {
	mediaPath := fmt.Sprintf("%s/%s", s.Conf["MEDIA_PATH"], passportSerial)
	log.Println(mediaPath)
	if _, err := os.Stat(mediaPath); os.IsNotExist(err) {
		err = os.Mkdir(mediaPath, os.ModePerm)
		if err != nil {
			raven.ReportIfError(err)
			log.Println(err.Error())
		}
	}
	switch fileType {
	case "passport_image":
		f, fh, err := r.FormFile("passport_image")
		if err == nil {
			file, err := os.Create(fmt.Sprintf("%s/passport_image_%s", mediaPath, fh.Filename))
			raven.ReportIfError(err)
			if err == nil {
				_, err = io.Copy(file, f)
				raven.ReportIfError(err)
				if err == nil {
					fName := strings.Split(file.Name(), "media/")
					return []string{fName[len(fName)-1]}
				}
			}
		}
		return []string{""}
	case "photo_1":
		f, fh, err := r.FormFile("photo_1")
		if err == nil {
			file, err := os.Create(fmt.Sprintf("%s/photo_1_%s", mediaPath, fh.Filename))
			raven.ReportIfError(err)
			if err == nil {
				_, err = io.Copy(file, f)
				raven.ReportIfError(err)
				if err == nil {
					fName := strings.Split(file.Name(), "media/")
					return []string{fName[len(fName)-1]}
				}
			}
		}
		return []string{""}
	case "photo_2":
		f, fh, err := r.FormFile("photo_2")
		if err == nil {
			file, err := os.Create(fmt.Sprintf("%s/photo_2_%s", mediaPath, fh.Filename))
			raven.ReportIfError(err)
			if err == nil {
				_, err = io.Copy(file, f)
				raven.ReportIfError(err)
				if err == nil {
					fName := strings.Split(file.Name(), "media/")
					return []string{fName[len(fName)-1]}
				}
			}
		}
		return []string{""}
	case "photo_3":
		f, fh, err := r.FormFile("photo_3")
		if err == nil {
			file, err := os.Create(fmt.Sprintf("%s/photo_3_%s", mediaPath, fh.Filename))
			raven.ReportIfError(err)
			if err == nil {
				_, err = io.Copy(file, f)
				raven.ReportIfError(err)
				if err == nil {
					fName := strings.Split(file.Name(), "media/")
					return []string{fName[len(fName)-1]}
				}
			}
		}
	case "photo_4":
		f, fh, err := r.FormFile("photo_4")
		if err == nil {
			file, err := os.Create(fmt.Sprintf("%s/photo_4_%s", mediaPath, fh.Filename))
			raven.ReportIfError(err)
			if err == nil {
				_, err = io.Copy(file, f)
				raven.ReportIfError(err)
				if err == nil {
					fName := strings.Split(file.Name(), "media/")
					return []string{fName[len(fName)-1]}
				}
			}
		}
	case "education":
		var names []string
		fhs := r.MultipartForm.File["edu_files"]
		for _, fh := range fhs {
			f, _ := fh.Open()
			file, err := os.Create(fmt.Sprintf("%s/education_%s", mediaPath, fh.Filename))
			raven.ReportIfError(err)
			_, err = io.Copy(file, f)
			raven.ReportIfError(err)
			if err == nil {
				fName := strings.Split(file.Name(), "media/")
				names = append(names, fName[len(fName)-1])
			}
		}
		return names
	case "army":
		var names []string
		fhs := r.MultipartForm.File["army_files"]
		for _, fh := range fhs {
			f, _ := fh.Open()
			file, err := os.Create(fmt.Sprintf("%s/army_%s", mediaPath, fh.Filename))
			raven.ReportIfError(err)
			_, err = io.Copy(file, f)
			raven.ReportIfError(err)
			if err == nil {
				fName := strings.Split(file.Name(), "media/")
				names = append(names, fName[len(fName)-1])
			}
		}
		return names
	case "family":
		var names []string
		fhs := r.MultipartForm.File["family_files"]
		for _, fh := range fhs {
			f, _ := fh.Open()
			file, err := os.Create(fmt.Sprintf("%s/army_%s", mediaPath, fh.Filename))
			raven.ReportIfError(err)
			_, err = io.Copy(file, f)
			raven.ReportIfError(err)
			if err == nil {
				fName := strings.Split(file.Name(), "media/")
				names = append(names, fName[len(fName)-1])
			}
		}
		return names
	case "reward":
		var names []string
		fhs := r.MultipartForm.File["reward_files"]
		for _, fh := range fhs {
			f, _ := fh.Open()
			file, err := os.Create(fmt.Sprintf("%s/reward_%s", mediaPath, fh.Filename))
			raven.ReportIfError(err)
			_, err = io.Copy(file, f)
			raven.ReportIfError(err)
			if err == nil {
				fName := strings.Split(file.Name(), "media/")
				names = append(names, fName[len(fName)-1])
			}
		}
		return names
	case "relative":
		var names []string
		fhs := r.MultipartForm.File["relative_files"]
		for _, fh := range fhs {
			f, _ := fh.Open()
			file, err := os.Create(fmt.Sprintf("%s/relative_%s", mediaPath, fh.Filename))
			raven.ReportIfError(err)
			_, err = io.Copy(file, f)
			raven.ReportIfError(err)
			if err == nil {
				fName := strings.Split(file.Name(), "media/")
				names = append(names, fName[len(fName)-1])
			}
		}
		return names
	case "experience":
		var names []string
		fhs := r.MultipartForm.File["exp_files"]
		for _, fh := range fhs {
			f, _ := fh.Open()
			file, err := os.Create(fmt.Sprintf("%s/experience_%s", mediaPath, fh.Filename))
			raven.ReportIfError(err)
			_, err = io.Copy(file, f)
			raven.ReportIfError(err)
			if err == nil {
				fName := strings.Split(file.Name(), "media/")
				names = append(names, fName[len(fName)-1])
			}
		}
		return names
	case "language":
		var names []string
		fhs := r.MultipartForm.File["lang_files"]
		for _, fh := range fhs {
			f, _ := fh.Open()
			file, err := os.Create(fmt.Sprintf("%s/language_%s", mediaPath, fh.Filename))
			raven.ReportIfError(err)
			_, err = io.Copy(file, f)
			raven.ReportIfError(err)
			if err == nil {
				fName := strings.Split(file.Name(), "media/")
				names = append(names, fName[len(fName)-1])
			}
		}
		return names
	}
	return []string{""}
}
