package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	name          string
	matchesPlayed int
	matchesWon    int
	matchesDrawn  int
	matchesLost   int
	points        int
}

func (t Team) String() string {

	return fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d", t.name, t.matchesPlayed, t.matchesWon, t.matchesDrawn, t.matchesLost, t.points)
}

func Tally(reader io.Reader, writer io.Writer) error {
	lineReader := bufio.NewScanner(reader)
	var table = map[string]Team{}
	for lineReader.Scan() {
		l := lineReader.Text()

		if l == "" || strings.HasPrefix(l, "#") {
			continue
		}

		match := strings.Split(l, ";")
		// fmt.Println(l)
		team1, ok := table[match[0]]
		if !ok {
			team1 = Team{
				name: match[0],
			}
		}
		team2, ok := table[match[1]]
		if !ok {
			team2 = Team{
				name: match[1],
			}
		}
		if len(match) != 3 {
			return fmt.Errorf("incorrect format")
		}
		result := match[2]

		switch result {
		case "win":

			team1.matchesPlayed++
			team1.matchesWon++
			team2.matchesPlayed++
			team2.matchesLost++
			team1.points += 3
			table[match[0]] = team1
			table[match[1]] = team2

		case "loss":

			team2.matchesPlayed++
			team2.matchesWon++
			team1.matchesPlayed++
			team1.matchesLost++
			team2.points += 3
			table[match[0]] = team1
			table[match[1]] = team2

		case "draw":
			team1.matchesPlayed++
			team1.matchesDrawn++
			team2.matchesPlayed++
			team2.matchesDrawn++
			team1.points += 1
			team2.points += 1
			table[match[0]] = team1
			table[match[1]] = team2
		default:
			return fmt.Errorf("incorrect result specified")
		}

	}
	var teams []Team
	io.WriteString(writer, "Team                           | MP |  W |  D |  L |  P\n")
	for _, team := range table {
		teams = append(teams, team)
	}

	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points == teams[j].points {
			return teams[i].name < teams[j].name
		}
		return teams[i].points > teams[j].points
	})
	for _, team := range teams {
		io.WriteString(writer, team.String()+"\n")
	}

	return nil

}
