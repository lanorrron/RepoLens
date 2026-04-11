# RepoLens

AI-powered CodeAct agent in Go that analyzes repositories by generating and executing Go scripts.

## 🚀 What it does

* Generates Go code to inspect a repo
* Executes the code locally
* Uses the output to create a report

## 🧠 Code-as-Action

The agent uses **generated Go code as actions**, executes it, and retries if it fails.

## 🛠️ Run

```bash
go run ./cmd/agent/main.go
```

## 📄 Output

* `report.md` (agent)
* `AI_REPORT.md` (inside analyzed repo)

## ⚠️ Notes

* Uses OpenAI API (`.env`)
* Scripts are read-only
