package handler

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/go-chi/chi/v5"
)

const SCRIPT_PREFIX = `source ~/miniconda3/etc/profile.d/conda.sh;conda activate pkasr;bash /mnt/md0/nfs_share/PKASR/sinica_asr/gstreamer`

type apiHandler struct{}

func RegisterApiHandler(router *chi.Mux) {
	handler := &apiHandler{}

	router.Route("/demo", func(apiRouter chi.Router) {
		apiRouter.Post("/postRecognize", handler.postRecognize)
		apiRouter.Post("/uploadRecognize", handler.uploadRecognize)
	})
}

// PostRecognize godoc
// @Summary Do post recognize
// @Description get post recognize result
// @Accept  json
// @Produce  json
// @Param Body body segmentInfo true "Set AsrKind to model name you want to use in post recognize"
// @Success 200 {array} wordalignment
// @Failure 400
// @Failure 500
// @Router /postRecognize [post]
func (handler apiHandler) postRecognize(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var segment segmentInfo

	err := decoder.Decode(&segment)
	if err != nil {
		fmt.Printf("%s", err.Error())
		w.WriteHeader(400)
		return
	}

	if segment.AsrKind == "formospeech_me_1" {
		segment.AsrKind = "sa_me_old"
	}

	audioPath := fmt.Sprintf("/mnt/md0/kaldi-gstreamer-server/tmp/%s.raw", segment.Id)

	command := exec.Command("bash", "-c", fmt.Sprintf("%s/run_rec_post.sh %s %s %s %f %f", SCRIPT_PREFIX, segment.LangKind, segment.AsrKind, audioPath, segment.Start, segment.Length))

	out, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err.Error())
		w.WriteHeader(500)
		return
	}
	fmt.Println(string(out))

	w.WriteHeader(200)
	w.Write(out)
}

// UploadRecognize godoc
// @Summary Do upload recognize
// @Description get upload recognize result
// @Accept multipart/form-data
// @Produce  json
// @Param Form formData uploadInfo true "Upload file via file field"
// @Success 200 {array} wordalignment
// @Failure 400
// @Failure 500
// @Router /uploadRecognize [post]
func (handler apiHandler) uploadRecognize(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024 * 1024 * 100)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	uploadFile, header, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(400)
		return
	}
	defer uploadFile.Close()

	asrKind := r.FormValue("asrKind")
	langKind := r.FormValue("langKind")

	if _, err := os.Stat("./files"); os.IsNotExist(err) {
		err := os.Mkdir("./files", 0755)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	filename := fmt.Sprintf("%x.%s", md5.Sum([]byte(strings.Split(header.Filename, ".")[0])), strings.Split(header.Filename, ".")[1])

	file, err := os.Create(fmt.Sprintf("files/%s", filename))
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(400)
		return
	}
	defer file.Close()

	io.Copy(file, uploadFile)

	command := exec.Command("bash", "-c", fmt.Sprintf("%s/run_rec_upload.sh %s %s %s", SCRIPT_PREFIX, langKind, asrKind, fmt.Sprintf("/mnt/md0/asr-demo/files/%s", filename)))

	out, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err.Error())
		w.WriteHeader(500)
		return
	}

	err = os.Remove(fmt.Sprintf("files/%s", filename))
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(500)
		return
	}

	fmt.Println(string(out))

	w.WriteHeader(200)
	w.Write(out)
}

// func (handler apiHandler) postRecognizeHandler(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	r.Form.Get()

// 	audioPath := fmt.Sprintf("/mnt/md0/user_dodohow1011/kaldi-gstreamer-server/tmp/%s.raw", segment.Id)

// 	command := exec.Command("bash", "-c", fmt.Sprintf("source ~/miniconda3/etc/profile.d/conda.sh;conda activate pkasr;bash /mnt/md0/PKASR/formospeech/gstreamer/run_rec_post.sh %s %s %s %f %f", segment.LangKind, segment.AsrKind, audioPath, segment.Start, segment.Length))

// 	out, err := command.CombinedOutput()
// 	if err != nil {
// 		fmt.Printf("%s", err.Error())
// 		w.WriteHeader(500)
// 		return
// 	}
// 	fmt.Println(string(out))

// 	w.WriteHeader(200)
// 	w.Write(out)
// }
