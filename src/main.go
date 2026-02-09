package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	// "google.golang.org/api/option"
	"google.golang.org/genai"
	"log"
	"path/filepath"
	// "strings"
)

func main() {

	// Initialize flags and flag values
	pathPtr := flag.String("path", "./src", "Path to your application code.")
	// modelPtr := flag.String("model", "gemini-3-flash-preview", "Choose a Google Gemini AI model.")
	flag.Parse()

	targetDirectory := *pathPtr

	// Initialize Gemini client
	ctx := context.Background()
	// The client gets the API key from the environment variable `GEMINI_API_KEY`.
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Read template files
	systemInstruction, err := os.ReadFile("templates/system_instruction.d")
	if err != nil {
		log.Fatal(err)
	}

	outputTemplate, err := os.ReadFile("templates/output_template.md")
	if err != nil {
		log.Fatal(err)
	}

	// Read Architectural Decision Records (ADRs)
	adrPath := filepath.Join(targetDirectory, "docs", "adr")
	adrContext := collectDesignDecisions(adrPath)

	codeContext, err := scanFiles(targetDirectory)
	if err != nil {
		log.Fatal(err)
	}

	// ---------------------------------------------------------
	// 5. BYGG PROMPTEN (OPPDATERT SANDWICH)
	// ---------------------------------------------------------
	// var sb strings.Builder

	// A. System Instruction (Rollen)
	// sb.Write(systemInstruction)
	// sb.WriteString("\n---\n")

	// B. Template (Formatet)
	// sb.WriteString("Please use this template:\n")
	// sb.Write(outputTemplate)
	// sb.WriteString("\n---\n")

	// C. Design Decisions (Kontekst) - NYTT!
	// Her primer vi AI-en med hvorfor ting er som de er.
	// sb.WriteString("Here are the Architecture Decision Records (ADR):\n")
	// sb.WriteString(adrContext)
	// sb.WriteString("\n---\n")

	// D. Koden (Fakta)
	// sb.WriteString("Here is the source code:\n")
	// sb.WriteString(codeContext)

	// fullPrompt := sb.String()

	// ---------------------------------------------------------
	// 6. SEND TIL AI (API KALL)
	// ---------------------------------------------------------
	// model := client.GenerativeModel(modelName)
	// resp, err := model.GenerateContent(ctx, genai.Text(fullPrompt))
	// Hent tekst fra responsen.

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-3-flash-preview",
		genai.Text("Explain why cats are cool in one sentence."),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Text())

	// ---------------------------------------------------------
	// 7. LAGRE RESULTATET
	// ---------------------------------------------------------
	// os.MkdirAll(filepath.Join(targetDirectory, "docs"), ...)
	// os.WriteFile(..., data, 0644)
	// Print "Done".
}

// ---------------------------------------------------------
// HJELPEFUNKSJON 1: FIL-SCANNER (Rekursiv)
// ---------------------------------------------------------
func scanFiles(rootPath string) (string, error) {
	// strings.Builder...
	// filepath.WalkDir(rootPath, func(...) {
	// 1. Ignorer .git, node_modules, etc (return filepath.SkipDir)
	// 2. Sjekk filendelser (.go, .tf, .yaml)
	// 3. Les fil
	// 4. Formater: "--- FILE: [navn] ---\n [innhold]"
	// })
	// return builder.String()
}

// ---------------------------------------------------------
// HJELPEFUNKSJON 2: LES DESIGNBESLUTNINGER (Flat liste)
// ---------------------------------------------------------
func collectDesignDecisions(adrPath string) string {
	// 1. SJEKK OM MAPPEN FINNES
	// GO-LÆRDOM: Bruk os.Stat() for å sjekke eksistens.
	// if _, err := os.Stat(adrPath); os.IsNotExist(err) {
	//     return "No ADRs found." // Helt ok, vi bare returnerer tomt.
	// }

	// 2. LES MAPPEN (IKKE REKURSIVT)
	// GO-LÆRDOM: Bruk os.ReadDir() når du bare vil ha filene i én mappe.
	// entries, err := os.ReadDir(adrPath)

	// strings.Builder...

	// 3. LOOP GJENNOM FILENE
	// for _, entry := range entries {
	//     if entry.IsDir() { continue } // Hopp over undermapper
	//     if !strings.HasSuffix(entry.Name(), ".md") { continue } // Kun markdown

	//     fullPath := filepath.Join(adrPath, entry.Name())
	//     content, _ := os.ReadFile(fullPath)

	//     sb.WriteString("\n--- DECISION RECORD: " + entry.Name() + " ---\n")
	//     sb.Write(content)
	// }

	// return sb.String()
}
