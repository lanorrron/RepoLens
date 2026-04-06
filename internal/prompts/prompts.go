package prompts

const ANALIZE_REPO = `
You are a senior Go engineer.

Analyze the following repository data and generate a report.

Classify issues into:
- High Risk
- Medium Risk
- Low Risk

Also provide recommendations.
IMPORTANT:
- Identify potential security vulnerabilities
- If applicable, relate issues to known vulnerability patterns (similar to CVEs)
- Prioritize real risks over style issues

Return in markdown format.

DATA:
`

func FixScriptPrompt(prevCode string, execError string) string {
	return `
The following Go script failed.

ERROR:
` + execError + `

CODE:
` + prevCode + `

Fix the code.

RULES:
- Only use standard library
- Only read files
- Do NOT delete or modify files
- Do NOT execute external commands

Return ONLY corrected code.
`
}

const ANALYSIS_SCRIPT = `
You are a senior Go engineer.

Generate a Go program that analyzes a repository.

REPO PATH:
%s

LANGUAGE DETECTED:
%s

REQUIREMENTS:
- Walk the repository recursively
- Read source files depending on language:
  - go → .go
  - typescript → .ts, .tsx
  - javascript → .js, .jsx
  - python → .py
- Ignore directories:
  node_modules, .git, dist, build, vendor

OUTPUT:
- Print file path
- Print first 200 lines of each file

STRICT RULES:
- ONLY Go
- ONLY standard library
- DO NOT delete or modify files
- DO NOT execute external commands

Return ONLY Go code.
`
