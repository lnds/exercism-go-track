package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Score struct {
	Wins   int
	Losses int
	Draws  int
}

func (score Score) Played() int {
	return score.Wins + score.Losses + score.Draws
}

func (score Score) Points() int {
	return score.Wins*3 + score.Draws
}

type ScoreTable map[string]Score

func (table ScoreTable) AddWin(team string) {
	score := table[team]
	score.Wins++
	table[team] = score
}

func (table ScoreTable) AddLoss(team string) {
	score := table[team]
	score.Losses++
	table[team] = score
}

func (table ScoreTable) AddDraw(team string) {
	score := table[team]
	score.Draws++
	table[team] = score
}

func (table ScoreTable) Output(writer *bufio.Writer) error {
	header := fmt.Sprintf("%-31.31s| MP |  W |  D |  L |  P\n", "Team")
	writer.WriteString(header)
	keys := []string{}
	for team := range table {
		keys = append(keys, team)
	}
	sort.Slice(keys, func(i, j int) bool {
		if table[keys[i]].Points() == table[keys[j]].Points() {
			return keys[i] < keys[j]
		}
		return table[keys[i]].Points() > table[keys[j]].Points()
	})

	for _, team := range keys {
		score := table[team]
		line := fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", team, score.Played(), score.Wins, score.Draws, score.Losses, score.Points())
		writer.WriteString(line)
	}
	return writer.Flush()
}

func Tally(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)
	scores := ScoreTable{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		cols := strings.Split(line, ";")
		if len(cols) != 3 {
			return fmt.Errorf("need three columns per line found len = %d [%s]", len(cols), line)
		}
		switch cols[2] {
		case "win":
			scores.AddWin(cols[0])
			scores.AddLoss(cols[1])
		case "draw":
			scores.AddDraw(cols[0])
			scores.AddDraw(cols[1])
		case "loss":
			scores.AddLoss(cols[0])
			scores.AddWin(cols[1])
		default:
			return fmt.Errorf("invalid result")
		}
	}
	writer := bufio.NewWriter(output)
	return scores.Output(writer)
}
