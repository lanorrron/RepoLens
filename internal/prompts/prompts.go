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

GOAL:
Analyze ONLY real developer-written source code.
You MUST aggressively filter out non-source files.

REQUIREMENTS:
- Walk the repository recursively
- Only process files matching the detected language:

  go -> .go
  typescript -> .ts, .tsx
  javascript -> .js, .jsx
  python -> .py

STRICT SOURCE FILTERING:

ALWAYS INCLUDE:
- Files inside: src/, app/, pages/, internal/, pkg/, cmd/, lib/, server/, api/

ALWAYS IGNORE DIRECTORIES:
- node_modules
- .git
- dist
- build
- .next
- out
- coverage
- tmp
- temp
- vendor
- .cache
- .turbo
- .vercel

ALWAYS IGNORE FILE TYPES:
- .map
- .min.js
- .log
- .lock
- .sum
- .meta
- .sst
- .previewinfo
- .tsbuildinfo
- .ico
- .png
- .jpg
- .svg
- .woff
- .woff2

HEURISTICS:
- If file size > 1MB → SKIP
- If file contains long single-line code → SKIP (likely minified)
- If file path contains "generated", "bundle", "chunk", "vendor" → SKIP

OUTPUT:
- Print file path
- Print only meaningful source code
- Skip files without relevant logic
- Limit output to 150 lines per file
- Prioritize:
  - business logic
  - API handlers
  - services
  - database code
  - auth
  - configuration
  
STRICT IMPORT RULES:
- Only include imports that are actually used
- DO NOT include unused imports
- The code MUST compile without modification

STRICT RULES:
- ONLY Go
- ONLY standard library
- DO NOT modify files
- DO NOT execute commands
- MUST compile

Return ONLY Go code.
`
