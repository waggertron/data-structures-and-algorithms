package graphs

import "sort"

type flightMap map[string][]flight

type flight struct {
	origin        string
	destination   string
	leave, arrive int
}

var present = struct{}{}

func NewFlightMap(flights []flight) flightMap {
	result := make(flightMap)

	for _, f := range flights {
		origin := f.origin
		flights, found := result[origin]
		if !found {
			flights = make([]flight, 0)
		}

		result[origin] = append(flights, f)
	}

	for _, flights := range result {
		sort.Slice(flights, func(i, j int) bool {
			return flights[i].leave < flights[j].leave
		})
	}

	return result

}

func GetFlightPath(startLocation, destination string, flights []flight) bool {
	graph := NewFlightMap(flights)
	skip := make(map[string]struct{})
	curTime := 0

	result := dfs(startLocation, destination, curTime, graph, skip)

	return result
}

func dfs(location, destination string, curTime int, graph flightMap, skip map[string]struct{}) bool {
	if location == destination {
		return true
	}

	flights := graph[location]
	skip[location] = present
	for _, f := range flights {
		if _, shouldSkip := skip[f.destination]; shouldSkip || f.leave < curTime {
			continue
		}

		if dfs(f.destination, destination, f.arrive, graph, skip) {
			return true
		}
	}

	return false
}
