# Gemini Docs
A CLI tool and GitHub Composite Action that generates automated architecture documentation, C4 diagrams, and code analysis using Google Gemini AI.

The tool parses your source code and Architecture Decision Records (ADRs) to generate an up-to-date `AI_GENERATED.md` file. This file includes architecture diagrams, integration data flows, and automated security and FinOps insights.

## Prerequisites & Folder Structure
This tool requires no configuration files for prompts or templates, as these are managed centrally by the Action.

To provide the AI with context regarding your design decisions, you may optionally create the following directory in your repository:
- `docs/adr/*.md` **(Optional)**: Place your Markdown Architecture Decision Records here. The tool will read them to understand the design intent behind your code.

The tool will output the following file:
- `docs/AI_GENERATED.md`: The generated documentation. This file is overwritten on every run.

## GitHub Action Usage
The recommended way to use Gemini Docs is as a CI step.
```yaml
name: Generate Architecture Docs

on:
  push:
    branches: [ main ]

jobs:
  update-docs:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Source Code
        uses: actions/checkout@v6

      - name: Run Gemini Docs
        uses: vetlekise/gemini-docs-cli@main
        with:
          api_key: ${{ secrets.GEMINI_API_KEY }}
          model: 'gemini-3-flash-preview'
          target_dir: '.'
```

> [!TIP]
> Chain this action with `peter-evans/create-pull-request` to automatically open a pull request whenever the generated documentation changes.

### Action Inputs
| Name | Description | Required | Default |
| -- | -- | -- | -- |
| `api_key` | Your Google Gemini API key. | Yes | |
| `target_dir` | The directory to scan, relative to the repository root. | No | `.` |
| `model` | The Google Gemini model to use for generation. | No | `gemini-3-flash-preview` |

## Local CLI Usage

You can run the tool locally using Go. Ensure your Google Gemini API key is exported in your environment.

```bash
export GEMINI_API_KEY="your_api_key_here"
go run src/main.go -path ./my-app -model gemini-3-flash-preview
```

### CLI Flags
| Flag | Description | Default |
| -- | -- | -- |
| `-path` | Path to the application code you want to scan. | `./` |
| `-model` | Choose a Google Gemini AI model. | `gemini-3-flash-preview` |

## How It Works

1. **File Scanning**: Recursively scans the target directory for relevant file extensions (`.go`, `.tf`, `.yaml`, `.py`, `.md`, `.ts`, `.js`).

2. **Exclusions**: Automatically parses and respects your repository's `.gitignore` rules. Additionally, it enforces a hardcoded blacklist for directories like `node_modules`, `vendor`, `.git`, and `docs`.

3. **ADR Ingestion**: Checks `docs/adr/` for markdown files to include as design context.

4. **Generation**: Sends the aggregated context to the Gemini API with strict system instructions to output structured Markdown and valid Mermaid C4 diagrams.

5. **Output**: Saves the result to `docs/AI_GENERATED.md`.

[!IMPORTANT]
Data Privacy: This tool sends the raw text of scanned source code and ADRs to a third-party LLM (Google Gemini API). Ensure you do not have hardcoded secrets committed to your repository before running the scan.