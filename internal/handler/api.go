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

type apiHandler struct{}

func RegisterApiHandler(router *chi.Mux) {
	handler := &apiHandler{}

	router.Route("/demo", func(apiRouter chi.Router) {
		apiRouter.Post("/postRecognize", handler.postRecognize)
		apiRouter.Post("/uploadRecognize", handler.uploadRecognize)
		apiRouter.Post("/translate", handler.translate)
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

	audioPath := fmt.Sprintf("/mnt/md0/user_dodohow1011/kaldi-gstreamer-server/tmp/%s.raw", segment.Id)

	command := exec.Command("bash", "-c", fmt.Sprintf("source ~/miniconda3/etc/profile.d/conda.sh;conda activate pkasr;sh /mnt/md0/PKASR/formospeech/gstreamer/run_rec_post.sh %s %s %s %f %f", segment.LangKind, segment.AsrKind, audioPath, segment.Start, segment.Length))

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

	command := exec.Command("bash", "-c", fmt.Sprintf("source ~/miniconda3/etc/profile.d/conda.sh;conda activate pkasr;sh /mnt/md0/PKASR/formospeech/gstreamer/run_rec_upload.sh %s %s %s", langKind, asrKind, fmt.Sprintf("/mnt/md0/user_dodohow1011/asr-demo/files/%s", filename)))

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

// Translate godoc
// @Summary Do translate after post recognize
// @Description get translate result
// @Accept  plain
// @Produce  plain
// @Param string body string true "string that need to translate"
// @Success 200 {string} string "return translated string"
// @Failure 500
// @Router /translate [post]
func (handler apiHandler) translate(w http.ResponseWriter, r *http.Request) {

	res, err := http.Post("http://140.109.19.147:5566/api/v1/translate", "application/json", r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(res.StatusCode)
	w.Write(body)
}
