package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	var ReactDir = "ganti/directory/react-mu!!"
	var ApiDir = "ganti/directory/api-mu!!"

	apiCmd, err := startApiServer(ApiDir)
	if err != nil {
		log.Fatalf("Gagal menjalankan server Laravel: %v", err)
	}

	reactCmd, err := startReactServer(ReactDir)
	if err != nil {
		log.Fatalf("Gagal menjalankan server React: %v", err)
	}

	fmt.Println("Menjalankan server Presence dan PresenceAPI...")

	openBrowser("http://localhost:5173")

	err = apiCmd.Wait()
	if err != nil {
		log.Printf("API server berhenti dengan error: %v", err)
	}

	err = reactCmd.Wait()
	if err != nil {
		log.Printf("Presence server berhenti dengan error: %v", err)
	}
}

func startReactServer(projectDir string) (*exec.Cmd, error) {
	cmd := exec.Command("npm", "run", "dev")
	cmd.Dir = projectDir
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	err := cmd.Start()
	return cmd, err
}

func startApiServer(projectDir string) (*exec.Cmd, error) {
	cmd := exec.Command("php", "artisan", "serve")
	cmd.Dir = projectDir
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	err := cmd.Start()
	return cmd, err
}

func openBrowser(url string) {
	var err error

	err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()

	if err != nil {
		log.Printf("Gagal membuka browser: %v\n", err)
	} else {
		fmt.Printf("Membuka browser dengan URL: %s\n", url)
	}
}