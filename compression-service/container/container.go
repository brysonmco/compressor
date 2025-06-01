package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	err, process := compress("./input.mkv", "./output.mp4", 1280, 720, "libx264", 23, "medium", 128)
	if err != nil {
		log.Fatalf("Error compressing video: %v", err)
	}
	log.Printf("Compression started on process: %v", process)

	http.HandleFunc("POST /probe", handleProbe)
	http.HandleFunc("POST /compress", handleCompress)

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

// POST /probe
func handleProbe(w http.ResponseWriter, r *http.Request) {

}

// POST /compress
func handleCompress(w http.ResponseWriter, r *http.Request) {

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
) (*os.Process, error) {
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
	err := cmd.Start()

	// Optional but good for debugging
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Process, err
}
