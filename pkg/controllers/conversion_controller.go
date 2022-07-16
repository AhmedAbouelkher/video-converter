package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/AhmedAbouelkher/video-converter/pkg/ffmpeg"
)

const (
	MAX_UPLOAD_SIZE = 1024 * 1024 * 1024 // 1 GB
)

func HandleDisplayingIndexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "./public/index.html")
}

func HandleUploadingFiled(w http.ResponseWriter, r *http.Request) {

	// ##### Upload file #####
	name, disPath, shouldReturn := extractAndSaveUploadedVideo(r, w)

	if shouldReturn {
		return
	}

	// ##### Process File #####
	cPath := fmt.Sprintf("./storage/compressions/%d", name)
	path, cErr := ffmpeg.GenerateHLSFromVideoSource(disPath, cPath)

	if cErr != nil {
		fmt.Print(cErr.Error())
	}

	http.Redirect(w, r, "/?path="+path, http.StatusSeeOther)
}

func extractAndSaveUploadedVideo(r *http.Request, w http.ResponseWriter) (int64, string, bool) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
		return 0, "", true
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0, "", true
	}

	defer file.Close()

	err = os.MkdirAll("./storage", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 0, "", true
	}

	name := time.Now().UnixMilli()
	disPath := fmt.Sprintf("./storage/%d%s", name, filepath.Ext(fileHeader.Filename))

	dst, err := os.Create(disPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 0, "", true
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 0, "", true
	}
	return name, disPath, false
}
