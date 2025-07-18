# MeLi_technical_challenge

This repository contains solutions to several technical challenges, each in its own subdirectory:

- [BestInGenre](BestInGenre/solution_best_in_genre.py)
- [Minesweeper](Minesweeper/solution_minesweeper.py)
- [FailuresReport](FailuresReport/FailuresReport.sql)
- [Summarizer](Summarizer/)

---

## BestInGenre

**Description:**  
A Python script that queries a TV series API and returns the series with the highest IMDB rating for a given genre. If two series have the same rating, the one with the lexicographically smaller name is returned.

**Main file:**  
[BestInGenre/solution_best_in_genre.py](BestInGenre/solution_best_in_genre.py)

**Requirements:**  
- Python 3.13.5
- See [BestInGenre/requirements.txt](BestInGenre/requirements.txt) for dependencies.

**How to run:**
```sh
pip install -r [requirements.txt](http://_vscodecontentref_/0)
python [solution_best_in_genre.py](http://_vscodecontentref_/1)

```
---

## Minesweeper

**Description:**
A Python implementation of the classic Minesweeper logic. Given a square matrix of 0s and 1s (mines), it returns a matrix where each cell contains the count of adjacent mines, or a special indicator if the cell itself is a mine.

**Main file:**  
[Minesweeper/solution_minesweeper.py](Minesweeper/solution_minesweeper.py)

**Requirements:**  
- Python 3.13.5
- See [Minesweeper/requirements.txt](BestInGenre/requirements.txt) for dependencies.

**How to run:**
```sh
pip install -r [requirements.txt](http://_vscodecontentref_/0)
python [solution_minesweeper.py](http://_vscodecontentref_/1)

```
---

## FailuresReponrt

**Description:**
A SQL query that reports customers with more than 3 campaign failures, joining customer, campaign, and event data.

**Main file:**  
[FailuresReport/FailuresReport.sql](FailuresReport/FailuresReport.py)

**How to use:**
Run the SQL query in your database environment with the appropriate tables (customers, campaigns, events).

---

## Summarizer

**Description:**  
A Go CLI tool that summarizes text files using the HuggingFace API. It supports different summary types (short, bullet, medium) and includes input validation and prompt injection sanitization.

**Main files:**  
[Summarizer/cmd/summarizer/main.go](Summarizer/cmd/summarize/main.go) entry endpoint
[Summarizer/internal/api/client.go](Summarizer/internal/api/client.go) (API client)
[Summarizer/types/models.go](Summarizer/types/models.go) (types and summary logic)

**Requirements:**  
- Go 1.24.5+
- Set the following environment variables:
    HUGGINGFACE_TOKEN: your HuggingFace API token
    HUGGINGFACE_ENDPOINT: the endpoint URL for the summatization model

**How to run:**
```sh
go run Summarizer\cmd\summarize main.go --type ( bullet || short || medium) --input test.txt
or go run Summarizer\cmd\summarize main.go -t ( bullet || short || medium) --input test.txt
or go run Summarizer\cmd\summarize main.go -t ( bullet || short || medium) test.txt

```

**Testing**
```sh
go test ./cmd/summarize
```
---

**Project structure** 
- Each project is self-contained with its own dependencies and environment (see respective requirements.txt or Go modules).
- For Go, dependencies are managed via Go modules.
