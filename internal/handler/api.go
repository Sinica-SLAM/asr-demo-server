package handler

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	httpmiddleware "github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"

	youtube "asr-demo-recognize/pkg/youtube"
)

const SCRIPT_PREFIX = `source ~/miniconda3/etc/profile.d/conda.sh;conda activate pkasr;bash /mnt/md0/nfs_share/PKASR/sinica_asr/demo`
const API_SCRIPT_PREFIX = `source ~/miniconda3/etc/profile.d/conda.sh;conda activate pkasr;bash /mnt/md0/nfs_share/PKASR/sinica_asr/api`

type apiHandler struct {
	youtubeService *youtube.Service
}

var results = make(map[string]result)

func RegisterApiHandler(router *chi.Mux, youtubeService *youtube.Service, mdlw httpmiddleware.Middleware) {
	handler := &apiHandler{
		youtubeService: youtubeService,
	}

	router.Route("/demo", func(apiRouter chi.Router) {
		apiRouter.Use(std.HandlerProvider("", mdlw))
		apiRouter.Post("/postRecognize", handler.postRecognize)
		apiRouter.Post("/uploadRecognize", handler.uploadRecognize)
		apiRouter.Post("/youtubeSrt", handler.youtubeSrt)
		apiRouter.Get("/result/{filename}", handler.getResult)
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

	audioPath := fmt.Sprintf("/mnt/md0/kaldi-gstreamer-server/tmp/%s.raw", segment.Id)

	command := exec.Command("bash", "-c", fmt.Sprintf("%s/run_rec_post.sh %s %s %s %f %f", SCRIPT_PREFIX, segment.LangKind, segment.AsrKind, audioPath, segment.Start, segment.Length))

	out, err := command.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(string(out))
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
// @Param Form formData uploadInfo true "Upload file via file field"
// @Success 200
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
		w.Write([]byte(err.Error()))
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

	filename := fmt.Sprintf("%x%s", md5.Sum([]byte(header.Filename)), strings.ToLower(filepath.Ext(header.Filename)))
	fmt.Println(filename)

	file, err := os.Create(fmt.Sprintf("files/%s", filename))
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	defer file.Close()

	_, err = io.Copy(file, uploadFile)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	results[filename] = result{Done: false, Data: "辨識中...(2/3)"}
	go handleUploadRecognize(langKind, asrKind, filename)

	w.WriteHeader(200)
	w.Write([]byte(filename))
}

func handleUploadRecognize(langKind string, asrKind string, filename string) {
	command := exec.Command("bash", "-c", fmt.Sprintf("%s/run_rec_upload.sh %s %s %s", SCRIPT_PREFIX, langKind, asrKind, fmt.Sprintf("/mnt/md0/asr-demo/files/%s", filename)))
	fmt.Printf("%s/run_rec_upload.sh %s %s %s\n", SCRIPT_PREFIX, langKind, asrKind, fmt.Sprintf("/mnt/md0/asr-demo/files/%s", filename))
	out, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", out)
		fmt.Println(err.Error())
		results[filename] = result{Done: true, Data: fmt.Sprintf("辨識失敗, %s", out)}
		return
	}

	err = os.Remove(fmt.Sprintf("files/%s", filename))
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(out))

	var data any
	err = json.Unmarshal(out, &data)
	if err != nil {
		fmt.Println(err.Error())
		results[filename] = result{Done: true, Data: fmt.Sprintf("轉換 json 失敗, %s", err.Error())}
		return
	}
	results[filename] = result{Done: true, Data: data}

}

func (handler apiHandler) youtubeSrt(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var info youtubeInfo

	err := decoder.Decode(&info)
	if err != nil {
		fmt.Printf("%s", err.Error())
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	id := uuid.New().String()
	results[id] = result{Done: false, Data: "辨識中...(2/5)"}
	go handler.handleYoutubeSrt(info.AsrKind, info.Vid, id)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
}

func (handler apiHandler) handleYoutubeSrt(asrKind string, vid string, id string) {
	command := exec.Command("bash", "-c", fmt.Sprintf("%s/run_rec_youtube.sh %s %s srt %s 10", API_SCRIPT_PREFIX, id, asrKind, vid))
	out, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", out)
		results[id] = result{Done: true, Data: fmt.Sprintf("辨識失敗, %s", out)}
		return
	}
	// so, err := command.StdoutPipe()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	w.WriteHeader(500)
	// 	return
	// }

	// scanner := bufio.NewScanner(so)

	// err = command.Start()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	w.WriteHeader(500)
	// 	return
	// }
	// unReadFirstLine := true
	// result := ""
	// for scanner.Scan() {
	// 	if unReadFirstLine {
	// 		result = scanner.Text()
	// 		unReadFirstLine = false
	// 	}
	// 	fmt.Println(scanner.Text())
	// 	fmt.Println(scanner.Err())
	// }

	srtPath := ""
	videoPath := ""

	if splitResult := strings.Split(strings.TrimSpace(string(out)), " "); len(splitResult) > 1 {
		srtPath = splitResult[0]
		videoPath = splitResult[1]
	} else {
		fmt.Println("script wrong output")
		results[id] = result{Done: true, Data: "辨識失敗, script wrong output"}
		return
	}

	results[vid] = result{Done: false, Data: "上傳影片至 YT...(3/5)"}

	videoFile, err := os.Open(videoPath)
	if err != nil {
		fmt.Printf("Error opening %v: %v\n", videoPath, err)
		results[id] = result{Done: true, Data: "上傳失敗, Error opening video"}
		return
	}
	defer videoFile.Close()

	videoRes, err := handler.youtubeService.UploadVideo(path.Base(videoPath), videoFile)
	if err != nil {
		fmt.Println(err.Error())
		results[id] = result{Done: true, Data: fmt.Sprintf("上傳影片至 YT 失敗, %s", err.Error())}
		return
	}

	time.Sleep(1 * time.Minute)

	results[vid] = result{Done: false, Data: "上傳字幕至 YT...(4/5)"}

	srtFile, err := os.Open(srtPath)
	if err != nil {
		fmt.Printf("Error opening %v: %v\n", srtFile, err)
		results[id] = result{Done: true, Data: "上傳失敗, Error opening srt"}
		return
	}
	defer srtFile.Close()

	_, err = handler.youtubeService.UploadCaptions(videoRes.Id, srtFile)
	if err != nil {
		fmt.Println(err.Error())
		results[id] = result{Done: true, Data: fmt.Sprintf("上傳字幕至 YT 失敗, %s", err.Error())}
		return
	}

	results[id] = result{Done: true, Data: map[string]string{"vid": videoRes.Id}}
}

func (handler apiHandler) getResult(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")

	if _, ok := results[filename]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(results[filename])
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(500)
		delete(results, filename)
		return
	}

	if results[filename].Done {
		delete(results, filename)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
