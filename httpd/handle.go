package httpd

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/miiniper/loges"
	"go.uber.org/zap"

	"github.com/julienschmidt/httprouter"
)

func Upload(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "some error")
			loges.Loges.Error("read static file error", zap.Error(err))
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		file, head, err := r.FormFile("file")
		if err != nil {
			loges.Loges.Error("read post file err", zap.Error(err))
			return
		}
		defer file.Close()
		tmpfile, err := os.Create("/tmp/" + head.Filename)
		if err != nil {
			loges.Loges.Error("create file err", zap.Error(err))
			return
		}
		defer tmpfile.Close()
		_, err = io.Copy(tmpfile, file)
		if err != nil {
			loges.Loges.Error("copy file error", zap.Error(err))
			return
		}

		http.Redirect(w, r, "/upload/ok", http.StatusFound)

	}

}
func UploadOk(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	io.WriteString(w, "upload file successful!")
}
