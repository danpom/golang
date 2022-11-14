package tournament

/*
example input:
Allegoric Alaskans;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskans;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskans;Courageous Californians;win


example output:
Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  3 |  2 |  1 |  0 |  7
Allegoric Alaskans             |  3 |  2 |  0 |  1 |  6
Blithering Badgers             |  3 |  1 |  0 |  2 |  3
Courageous Californians        |  3 |  0 |  1 |  2 |  1
*/

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name          string
	matchesPlayed int
	wins          int
	draws         int
	losses        int
	points        int
}

type teamSlice []*team

type teamRecords map[string]*team

func (tr *teamRecords) addTeamScore(n string, w int, d int, l int) {
	if val, ok := (*tr)[n]; ok {
		val.matchesPlayed++
		val.wins += w
		val.draws += d
		val.losses += l
		val.points += w*3 + d*1

		(*tr)[n] = val

	} else {
		(*tr)[n] = &team{name: n, matchesPlayed: 1, wins: w, draws: d, losses: l, points: w*3 + d*1}
	}
}

func Tally(reader io.Reader, writer io.Writer) error {
	tr := make(teamRecords)

	s := bufio.NewScanner(reader)
	for s.Scan() {
		l := s.Text()

		// Skip newlines and comments
		if l == "" || l[0] == '#' {
			continue
		}

		teamsScore := strings.Split(l, ";")
		if len(teamsScore) != 3 {
			return fmt.Errorf("wring line format: %s (has %d fields)", l, len(teamsScore))
		}
		t1, t2, result := teamsScore[0], teamsScore[1], teamsScore[2]

		switch result {
		case "win":
			tr.addTeamScore(t1, 1, 0, 0)
			tr.addTeamScore(t2, 0, 0, 1)
		case "draw":
			tr.addTeamScore(t1, 0, 1, 0)
			tr.addTeamScore(t2, 0, 1, 0)
		case "loss":
			tr.addTeamScore(t1, 0, 0, 1)
			tr.addTeamScore(t2, 1, 0, 0)
		default:
			return fmt.Errorf("invalid outcome %s", result)
		}

	}
	//sort - maps are unordered so creating a slice to do the sorting
	ts := make(teamSlice, 0, len(tr))

	for _, t := range tr {
		ts = append(ts, t)
	}

	//sorting by points
	sort.Slice(ts, func(i, j int) bool {
		//on even points sorting by name
		if ts[i].points == ts[j].points {
			return ts[i].name < ts[j].name
		}
		return ts[i].points > ts[j].points
	})

	fmt.Fprintf(writer, "%-31s| %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	for _, v := range ts {
		fmt.Fprintf(writer, "%-31s| %2d | %2d | %2d | %2d | %2d\n", v.name, v.matchesPlayed, v.wins, v.draws, v.losses, v.points)
	}

	return nil
}
