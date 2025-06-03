package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	fmt.Println("APPLICATION_STARTED")
	http.HandleFunc("POST /download", handleDownload)
	http.HandleFunc("POST /probe", handleProbe)
	http.HandleFunc("POST /compress", handleCompress)

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("SERVER_FAILED")
	}
}

// POST /download
type downloadRequest struct {
	URL       string `json:"url"`
	Container string `json:"container"`
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	var req downloadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("DOWNLOAD_FAILED")
		WriteError(w, http.StatusBadRequest, "invalid request body", "invalid_request_body", err)
		return
	}

	if req.URL == "" || req.Container == "" {
		fmt.Println("DOWNLOAD_FAILED")
		WriteError(w, http.StatusBadRequest, "missing required fields", "missing_fields", "URL and Container are required")
		return
	}

	// Get the expected length of the file
	headResp, err := http.Head(req.URL)
	if err != nil {
		fmt.Println("DOWNLOAD_FAILED")
		WriteError(w, http.StatusInternalServerError, "could not fetch file info", "fetch_error", err)
		return
	}
	if headResp.StatusCode != http.StatusOK {
		fmt.Println("DOWNLOAD_FAILED")
		WriteError(w, http.StatusBadRequest, "file not found", "file_not_found", fmt.Sprintf("status code: %d", headResp.StatusCode))
		return
	}
	contentLengthStr := headResp.Header.Get("Content-Length")
	if contentLengthStr == "" {
		fmt.Println("DOWNLOAD_FAILED")
		WriteError(w, http.StatusBadRequest, "missing Content-Length header", "missing_header", "Content-Length is required for download")
		return
	}
	contentLength, err := strconv.ParseInt(contentLengthStr, 10, 64)
	if err != nil {
		fmt.Println("DOWNLOAD_FAILED")
		WriteError(w, http.StatusBadRequest, "invalid Content-Length header", "invalid_header", err)
		return
	}

	path := fmt.Sprintf("./input.%s", req.Container)

	// Download the file from the URL
	cmd, err := downloadFile(req.URL, path)
	if err != nil {
		fmt.Println("DOWNLOAD_FAILED")
		WriteError(w, http.StatusInternalServerError, "could not download file", "download_error", err)
		return
	}

	WriteSuccess(w, http.StatusCreated, "file download started", nil)
	go watchDownload(cmd, path, contentLength)
}

func downloadFile(
	url string,
	path string,
) (*exec.Cmd, error) {
	cmd := exec.Command(
		"curl",
		"-L",       // Follow redirects
		"-o", path, // Output to the specified path
		url, // The URL to download
	)

	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd, nil
}

func watchDownload(
	cmd *exec.Cmd,
	filePath string,
	expectedLength int64,
) {
	err := cmd.Wait()
	if err != nil {
		fmt.Println("DOWNLOAD_FAILED")
		return
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("DOWNLOAD_FAILED")
		return
	}

	if fileInfo.Size() != expectedLength {
		fmt.Println("DOWNLOAD_FAILED")
		return
	}

	fmt.Println("DOWNLOAD_COMPLETED")
}

// POST /probe
type fFProbeStream struct {
	CodecName  string `json:"codec_name"`
	CodecType  string `json:"codec_type"`
	Width      int    `json:"width,omitempty"`
	Height     int    `json:"height,omitempty"`
	SampleRate string `json:"sample_rate,omitempty"`
}

type fFProbeFormat struct {
	Filename   string `json:"filename"`
	NbStreams  int    `json:"nb_streams"`
	FormatName string `json:"format_name"`
}

type fFProbeOutput struct {
	Streams []fFProbeStream `json:"streams"`
	Format  fFProbeFormat   `json:"format"`
}

func handleProbe(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("ffprobe",
		"-v", "quiet",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
		"",
	)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("PROBE_FAILED")
		WriteError(w, http.StatusInternalServerError, "could not run ffprobe", "ffprobe_error", err)
		return
	}

	var probeOutput fFProbeOutput
	if err = json.Unmarshal(output, &probeOutput); err != nil {
		fmt.Println("PROBE_FAILED")
		WriteError(w, http.StatusInternalServerError, "could not parse ffprobe output", "json_unmarshal_error", err)
		return
	}

	probeBytes, err := json.Marshal(probeOutput)
	if err != nil {
		fmt.Println("PROBE_FAILED")
		WriteError(w, http.StatusInternalServerError, "could not marshal ffprobe output", "json_marshal_error", err)
		return
	}

	fmt.Println("START_PROBE_DATA")
	fmt.Println(string(probeBytes))
	fmt.Println("END_PROBE_DATA")

	WriteSuccess(w, http.StatusOK, "probe data retrieved", probeOutput)
}

// POST /compress
type compressRequest struct {
	InputContainer  string `json:"inputContainer"`
	OutputContainer string `json:"outputContainer"`
	MaxWidth        int    `json:"maxWidth"`
	MaxHeight       int    `json:"maxHeight"`
	Codec           string `json:"codec"`
	Crf             int    `json:"crf"`
	Preset          string `json:"preset"`
	AudioBitrate    int    `json:"audioBitrate"`
}

func handleCompress(w http.ResponseWriter, r *http.Request) {
	var req compressRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("COMPRESSION_FAILED")
		WriteError(w, http.StatusBadRequest, "invalid request body", "invalid_request_body", err)
		return
	}

	outputPath := fmt.Sprintf("./output.%s", req.InputContainer)

	cmd, err := compress(
		fmt.Sprintf("./input.%s", req.InputContainer),
		outputPath,
		req.MaxWidth,
		req.MaxHeight,
		req.Codec,
		req.Crf,
		req.Preset,
		req.AudioBitrate,
	)
	if err != nil {
		fmt.Println("COMPRESSION_FAILED")
		WriteError(w, http.StatusInternalServerError, "could not start compression", "internal_error", err)
		return
	}

	WriteSuccess(w, http.StatusCreated, "compression started", nil)
	fmt.Println("STARTED_COMPRESSION")

	go watchCompression(cmd, outputPath)
}

// TODO: currently only works for H.264 and H.265 codecs
func compress(
	inputPath string,
	outputPath string,
	maxWidth int,
	maxHeight int,
	codec string,
	crf int,
	preset string,
	audioBitrate int,
) (*exec.Cmd, error) {
	vf := fmt.Sprintf(
		"scale='min(%d,iw)':'min(%d,ih)':force_original_aspect_ratio=decrease,pad=ceil(iw/2)*2:ceil(ih/2)*2",
		maxWidth, maxHeight,
	)

	cmd := exec.Command(
		"ffmpeg",
		"-i", inputPath,
		"-vf", vf,
		"-c:v", codec,
		"-crf", strconv.Itoa(crf),
		"-preset", preset,
		"-c:a", "aac",
		"-b:a", fmt.Sprintf("%dk", audioBitrate),
		"-ar", "44100",
		outputPath,
	)
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func watchCompression(
	cmd *exec.Cmd,
	filePath string,
) {
	err := cmd.Wait()
	if err != nil {
		fmt.Println("COMPRESSION_FAILED")
	}

	// Ensure file is present
	_, err = os.Stat(filePath)
	if err != nil {
		fmt.Println("COMPRESSION_FAILED")
	}

	fmt.Println("COMPRESSION_COMPLETED")
}

// POST /upload
type uploadRequest struct {
	URL       string `json:"url"`
	Container string `json:"container"`
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	var req uploadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("UPLOAD_FAILED")
		WriteError(w, http.StatusBadRequest, "invalid request body", "invalid_request_body", err)
		return
	}

	if req.URL == "" || req.Container == "" {
		fmt.Println("UPLOAD_FAILED")
		WriteError(w, http.StatusBadRequest, "missing required fields", "missing_fields", "URL and Container are required")
		return
	}

	// Upload the file to the specified URL
	cmd := exec.Command(
		"curl",
		"-X", "POST",
		"-F", fmt.Sprintf("file=@%s", fmt.Sprintf("./output.%s", req.Container)),
		"-F", fmt.Sprintf("container=%s", req.Container))
	fmt.Println(cmd)
}

type ErrorResponse struct { // Human-readable
	Error   string      `json:"error"` // Machine-readable
	Details interface{} `json:"details"`
}

func WriteError(
	w http.ResponseWriter,
	code int,
	message string,
	error string,
	details interface{},
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"success":   false,
		"status":    code,
		"timestamp": time.Now(),
		"message":   message,
		"error": ErrorResponse{
			Error:   error,
			Details: details,
		},
	}); err != nil {
		log.Printf("error encoding JSON response: %v", err)
	}
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

func WriteSuccess(
	w http.ResponseWriter,
	code int,
	message string,
	data interface{},
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"success":   true,
		"status":    code,
		"timestamp": time.Now(),
		"message":   message,
		"data":      data,
	}); err != nil {
		log.Printf("error encoding JSON response: %v", err)
	}
}
