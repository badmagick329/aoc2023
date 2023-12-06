package day6

const FILE = "day6/file.txt"

func Run() {
    RunOne()
    RunTwo()
}

func calculateDistance(totalRaceTime, pressDuration int) int {
	raceTime := totalRaceTime - pressDuration
	return raceTime * pressDuration
}

func nWaysToBeat(totalRaceTime, distanceToBeat int) int {
	winningPressDurations := []int{}
	for pressDuration := 0; pressDuration < totalRaceTime; pressDuration++ {
		distance := calculateDistance(totalRaceTime, pressDuration)
		// fmt.Printf(
		// 	"RACE TIME %d: Pressed %d. Distance: %d\n",
		// 	totalRaceTime,
		// 	pressDuration,
		// 	distance,
		// )
		if distance < distanceToBeat && pressDuration > (totalRaceTime/2) {
			break
		}
		if distance > distanceToBeat {
			winningPressDurations = append(winningPressDurations, pressDuration)
		}
	}
	return len(winningPressDurations)
}

