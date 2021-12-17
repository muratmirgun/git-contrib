package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/olekukonko/tablewriter"
)

func main() {
	report, err := createReport()
	if err != nil {
		fmt.Println(err.Error())
	}

	printReport(report)
}

func createReport() (map[string]int, error) {
	repository, err := git.PlainOpen("./")
	if err != nil {
		return nil, err
	}

	commits, err := repository.Log(&git.LogOptions{})
	if err != nil {
		return nil, err
	}

	lastCommit, err := commits.Next()
	if err != nil {
		return nil, err
	}

	report := make(map[string]int)

	tree, err := lastCommit.Tree()
	if err != nil {
		return nil, err
	}

	tree.Files().ForEach(func(file *object.File) error {
		blameResult, err := git.Blame(lastCommit, file.Name)
		if err != nil {
			fmt.Printf("git blame failed for %s, %s\n", file.Name, err.Error())
			return nil
		}

		for _, line := range blameResult.Lines {
			author := line.Author
			lineCount, ok := report[author]
			if ok {
				report[author] = lineCount + 1
				continue
			}

			report[author] = 1
		}

		return nil
	})

	return report, nil
}

func printReport(report map[string]int) {
	totalCount := 0
	for _, lineCount := range report {
		totalCount += lineCount
	}

	tableData := [][]string{}
	for author, lineCount := range report {
		sLineCount := strconv.Itoa(lineCount)
		percentage := float32(lineCount) / float32(totalCount) * 100
		sPercentage := fmt.Sprintf("%.2f", percentage)

		tableData = append(tableData, []string{author, sLineCount, sPercentage})
	}

	sort.Slice(tableData, func(i, j int) bool {
		lineCountI, _ := strconv.Atoi(tableData[i][1])
		lineCountJ, _ := strconv.Atoi(tableData[j][1])

		return lineCountI > lineCountJ
	})

	tableData = append([][]string{{"author", "line count", "percentage"}}, tableData...)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.AppendBulk(tableData)
	table.Render()
}
