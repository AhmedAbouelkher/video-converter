package ffmpeg

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func GenerateHLSFromVideoSource(s string, o string) (string, error) {
	fmt.Println("Started processing\t" + s)
	cmd := exec.Command("./create-vod-hls.sh", s, o)

	if err := cmd.Run(); err != nil {
		return "", err
	}

	fmt.Println("Finished processing")

	return filepath.Base(o) + "/playlist.m3u8", nil
}
