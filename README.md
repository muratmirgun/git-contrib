# git-contrib

git-contrib creates "git blame" based contributors report for git repositories different than the GitHub's Insights->Contributors report.

[![Go](https://github.com/ermanimer/git-contrib/actions/workflows/go.yml/badge.svg)](https://github.com/ermanimer/git-contrib/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ermanimer/git-contrib)](https://goreportcard.com/report/github.com/ermanimer/git-contrib)

# Installation 

```bash
go install github.com/ermanimer/git-contrib@latest
```

# Usage

Run following command in a git repository.

```bash
git contrib
```

or

```go
git-contrib
```

# Sample Output

```bash
+---------------------+------------+------------+
| author              | line count | percentage |
+---------------------+------------+------------+
| ermanimer@gmail.com |        328 |     100.00 |
+---------------------+------------+------------+
```
