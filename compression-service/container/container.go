package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

func main() {

	http.HandleFunc("POST /probe", handleProbe)
	http.HandleFunc("POST /compress", handleCompress)

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("SERVER_START_FAILED")
	}
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
		log.Println("PROBE_FAILED")
		WriteError(w, http.StatusInternalServerError, "could not run ffprobe", "ffprobe_error", err)
		return
	}

	var probeOutput fFProbeOutput
	if err = json.Unmarshal(output, &probeOutput); err != nil {
		log.Println("PROBE_JSON_UNMARSHAL_FAILED")
		WriteError(w, http.StatusInternalServerError, "could not parse ffprobe output", "json_unmarshal_error", err)
		return
	}

	probeBytes, err := json.Marshal(probeOutput)
	if err != nil {
		log.Println("PROBE_JSON_MARSHAL_FAILED")
		WriteError(w, http.StatusInternalServerError, "could not marshal ffprobe output", "json_marshal_error", err)
		return
	}

	fmt.Println(string(probeBytes))
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
		WriteError(w, http.StatusBadRequest, "invalid request body", "invalid_request_body", err)
		return
	}

	cmd, err := compress(
		fmt.Sprintf("./input.%s", req.InputContainer),
		fmt.Sprintf("./output.%s", req.InputContainer),
		req.MaxWidth,
		req.MaxHeight,
		req.Codec,
		req.Crf,
		req.Preset,
		req.AudioBitrate,
	)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "could not start compression", "internal_error", err)
		return
	}

	WriteSuccess(w, http.StatusCreated, "compression started", nil)

	go waitForCompletion(cmd)
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
		return nil, fmt.Errorf("failed to start command: %w", err)
	}

	return cmd, nil
}

func waitForCompletion(cmd *exec.Cmd) {
	err := cmd.Wait()
	if err != nil {
		log.Fatal("COMPRESSION_FAILED")
	} else {
		log.Println("compression completed successfully")
	}
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
