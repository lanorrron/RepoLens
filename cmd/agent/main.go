package main

import (
	"RepoLens/internal/config"
	"RepoLens/internal/inspector"
	"RepoLens/internal/runner"
	"RepoLens/internal/utils"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		fmt.Println("config error:", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the repository path: ")

	repoPath, error := reader.ReadString('\n')
	if error != nil {
		fmt.Println("Input error: ", error)
	}
	repoPath = strings.TrimSpace(repoPath)
	finalPath, err := utils.VerifyPath(repoPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	lang := inspector.DetectLanguage(finalPath)
	runner.Log("Detected language", lang)

	code := runner.GenerateScript(cfg, lang, finalPath)
	output := runner.RunWithRetries(cfg, code, lang, 3)

	report := runner.AnalyzeRepo(cfg, output)
	runner.SaveReport("report.md", report)
	repoReportPath := filepath.Join(finalPath, "AI_REPORT.md")
	runner.SaveReport(repoReportPath, report)

	runner.Log("Report generated", "report.md")

}
