# git-contrib

Creates simple contributors report for git.

Useful for different branches then main.

[![Go](https://github.com/ermanimer/git-contrib/actions/workflows/go.yml/badge.svg)](https://github.com/ermanimer/git-contrib/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ermanimer/git-contrib)](https://goreportcard.com/report/github.com/ermanimer/git-contrib)

# Installation 

```bash
go install github.com/ermanimer/git-contrib@latest
```

# Usage

Run following command in a git repository.

```bash
git-contrib
```

# Sample Output

```bash
+---------------------+------------+------------+
| author              | line count | percentage |
+---------------------+------------+------------+
| ermanimer@gmail.com |        268 |     100.00 |
+---------------------+------------+------------+
```
