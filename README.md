# GoMatchMaker

My first Go application. A CLI compatibility quiz that compares your answers to a set of pre-defined responses to calculate a match percentage.

## How It Works

The program presents 5 opinion questions on a Likert scale (1–5, strongly disagree to strongly agree). Your answers are scored against the developer's pre-set answers — the closer your responses, the higher your compatibility score.

**Scoring:** Each question is worth up to 4 points based on the absolute difference from the expected answer.

## Requirements

- Go 1.21.6 or later

## Running

```bash
go run matchmaker.go
```

## Example Output

```
******************************************************
                    Matchmaker 3.0
******************************************************
This program figures out if you and I are meant to be.
...
Your score is: 85%
You are a great match!
```
