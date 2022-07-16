package ffmpeg

import (
	"fmt"
	"os/exec"
	"time"
)


type ResolutionConversion struct {
	Input string
	Output string
	Resolution *Resolution
}

type Resolution struct {
	Width int32
	Hight int32
}

func (r *Resolution) Size() int32 {
	return r.Hight * r.Width
}

func (r *Resolution) String() string {
	return fmt.Sprintf("%d:%d", r.Width, r.Hight)
}

func (r *Resolution) Scale() string {
	return fmt.Sprintf("scale=%s", r.String())
}

func (r *Resolution) AspectRatio() int32 {
	return r.Width / r.Hight
}


const (
	DefaultProcessTimeOutDurationInSeconds = 30
)


func ConvertVideoResolution(conversion *ResolutionConversion) error {
	done := make(chan error, 1)

	cmd := exec.Command("ffmpeg", "-i", conversion.Input, "-vf", conversion.Resolution.Scale(),  conversion.Output, "-y")
	cmd.Start()

	go func(){
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(time.Second * DefaultProcessTimeOutDurationInSeconds):
		cmd.Process.Kill()
		return fmt.Errorf("process with pid %d timed out", cmd.Process.Pid)
	case <-done:
		return nil
	}
}