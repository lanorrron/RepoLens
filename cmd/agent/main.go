package main

import (
	"RepoLens/internal/config"
	"RepoLens/internal/inspector"
	"RepoLens/internal/runner"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		fmt.Println("config error:", err)
		return
	}
	repoPath := "repo path"

	lang := inspector.DetectLanguage(repoPath)
	runner.Log("Detected language", lang)

	code := runner.GenerateScript(cfg, lang, repoPath)
	output := runner.RunWithRetries(cfg, code, lang, 2)

	report := runner.AnalyzeRepo(cfg, output)
	runner.SaveReport("report.md", report)
	repoReportPath := repoPath + "/AI_REPORT.md"
	runner.SaveReport(repoReportPath, report)

	runner.Log("Report generated", "report.md")

}
