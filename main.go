package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func getFlags() (token, port string) {
	flag.StringVar(&token, "token", "", "Your Telegram bot token")
	flag.StringVar(&port, "port", "3000", "Port that the server runs on")

	flag.Usage = func() {
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if token == "" {
		flag.Usage()
		os.Exit(1)
	}

	return
}

func main() {
	_, port := getFlags()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := exec.CommandContext(ctx, "npx", "localtunnel", "--port", port)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Ошибка подключения stdout: %v", err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal("Ошибка запуска:", err)
	}

	go func() {
		sig := <-sigChan
		fmt.Println("Signal:", sig)
		cancel()
	}()

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("LT:", line)

			if checkConnection(line) {
				password, err := getPassword()
				if err != nil {
					log.Fatal("Ошибка получения пароля:", err)
				}
				fmt.Println("Пароль:", password)
			}
		}
	}()

	if err := cmd.Wait(); err != nil {
		if ctx.Err() != context.Canceled {
			fmt.Println("Ошибка выполнения:", err)
		}
	}
}

func checkConnection(scannedText string) bool {
	return strings.Contains(scannedText, "your url is: https://")
}

func getPassword() (string, error) {
	const timeout = 300 * time.Millisecond
	var lastErr error

	for range 3 {
		resp, err := http.Get("https://loca.lt/mytunnelpassword")
		if err != nil {
			lastErr = fmt.Errorf("HTTP запрос failed: %w", err)
			time.Sleep(timeout)
			continue
		}

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			lastErr = fmt.Errorf("чтение тела ответа: %w", err)
			time.Sleep(timeout)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			lastErr = fmt.Errorf("status code: %d, body: %s", resp.StatusCode, string(respBody))
			time.Sleep(timeout)
			continue
		}

		return string(respBody), nil
	}

	return "", fmt.Errorf("после 3 попыток не удалось получить пароль: %w", lastErr)
}

// func setWebAppUrl(token, url string) error {
// 	apiUrl := fmt.Sprintf("https://api.telegram.org/bot%s/setWebAppUrl", token)
// }
