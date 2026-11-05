package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Exercise struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Track       string `json:"track"`
}

var exercisesFile = "exercises.json"

func loadExercises() ([]Exercise, error) {
	data, err := os.ReadFile(exercisesFile)
	if err != nil {
		return nil, err
	}
	var ex []Exercise
	if err := json.Unmarshal(data, &ex); err != nil {
		return nil, err
	}
	return ex, nil
}

func findExercise(id int) (*Exercise, error) {
	ex, err := loadExercises()
	if err != nil {
		return nil, err
	}
	for _, e := range ex {
		if e.ID == id {
			return &e, nil
		}
	}
	return nil, fmt.Errorf("exercise %d not found", id)
}

func runTests(dir string) error {
	cmd := exec.Command("go", "test", "-count=1", dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	verify := flag.Bool("verify-solutions", false, "verify all solutions")
	list := flag.Bool("list", false, "list exercises")
	reset := flag.Int("reset", 0, "reset exercise by id")
	id := flag.Int("id", 0, "exercise id to check")
	only := flag.Int("only", 0, "exercise id to check (alias for --id)")
	checkAll := flag.Bool("check-all", false, "check all exercises in exercises/")
	flag.Parse()

	ex, err := loadExercises()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if *list {
		for _, e := range ex {
			fmt.Printf("[%d] %s - %s\n", e.ID, e.Title, e.Name)
		}
		return
	}

	if *verify {
		fmt.Println("Verifying solutions...")
		for _, e := range ex {
			solPath := "./solutions/" + e.Name
			cmd := exec.Command("go", "test", "-count=1", "-run", "Test", solPath)
			err := cmd.Run()
			status := "OK"
			if err != nil {
				status = "FAIL"
			}
			fmt.Printf("[%d] %s .................. %s\n", e.ID, e.Title, status)
		}
		return
	}

	if *checkAll {
		fmt.Println("Checking all exercises...")
		passed := 0
		failed := 0
		for _, e := range ex {
			path := "./exercises/" + e.Name
			cmd := exec.Command("go", "test", "-count=1", "-run", "Test", path)
			err := cmd.Run()
			status := "PASS"
			if err != nil {
				status = "FAIL"
				failed++
			} else {
				passed++
			}
			fmt.Printf("[%d] %s .................. %s\n", e.ID, e.Title, status)
		}
		fmt.Printf("\nSummary: %d PASS, %d FAIL / %d total\n", passed, failed, len(ex))
		return
	}

	if *reset > 0 {
		e, err := findExercise(*reset)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		src := filepath.Join(".pristine", e.Name, "main.go")
		dst := filepath.Join("exercises", e.Name, "main.go")
		content, err := os.ReadFile(src)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := os.WriteFile(dst, content, 0644); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("Reset %d %s\n", e.ID, e.Title)
		return
	}

	if *only > 0 {
		*id = *only
	}

	if *id > 0 {
		e, err := findExercise(*id)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		path := "./exercises/" + e.Name
		fmt.Printf("Running tests for %d %s ...\n", e.ID, e.Title)
		if err := runTests(path); err != nil {
			fmt.Printf("Result: FAIL\n")
			os.Exit(1)
		}
		fmt.Printf("Result: PASS\n")
		return
	}

	fmt.Println("Use --list, --verify-solutions, --check-all, --reset <id>, or --id <id>")
}
