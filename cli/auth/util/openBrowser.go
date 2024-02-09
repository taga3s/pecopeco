package util

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform.")
	}
	if err != nil {
		return errors.New("\nSorry, could not open a browser. Open the link above in your browser to manually complete authentication.")
	}
	return nil
}
