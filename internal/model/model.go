package model

import (
	"time"

	pb "github.com/MydroX/geotwin-grpctest/internal/positions/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SimulationResult struct {
	ScenarioId         string       `json:"scenarioId"`
	StatesEntries      []StateEntry `json:"statesEntries"`
	Period             Period       `json:"period"`
	MobilityProviderId string       `json:"mobilityProviderId"`
}

type StateEntry struct {
	AgentId    uint64      `json:"agentId"`
	Capacity   uint16      `json:"capacity"`
	Load       uint16      `json:"load"`
	Route      [][]float64 `json:"route"`
	StateId    string      `json:"stateId"`
	Timestamps []uint64    `json:"timestamps"`
}

// type Position struct {
// 	Longitude string `json:"longitude"`
// 	Latitude  string `json:"latitude"`
// }

type Period struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func JSONEntriesToGRPC(entries []StateEntry) []*pb.Entry {
	var res []*pb.Entry
	for _, entry := range entries {
		var tempEntry pb.Entry

		tempEntry.AgentId = entry.AgentId
		tempEntry.Capacity = uint32(entry.Capacity)
		tempEntry.Load = uint32(entry.Load)
		tempEntry.StateId = pb.Entry_StateId(pb.Entry_StateId_value[entry.StateId])

		var tempTimestamps []*timestamppb.Timestamp
		for _, timestamp := range entry.Timestamps {
			tm := time.Unix(int64(timestamp), 0)
			tempTimestamps = append(tempTimestamps, timestamppb.New(tm))
		}
		tempEntry.Timestamps = tempTimestamps

		res = append(res, &tempEntry)
	}
	return res
}

// func stateIdJSONToGRPC(idStr string) (id int32) {
// 	switch idStr {
// 	case "waiting":
// 		id = 1
// 	case "waitingForMission":
// 		id = 2
// 	case "waitingForPassengers":
// 		id = 3
// 	case "alightingPassengers":
// 		id = 4
// 	case "boardingPassengers":
// 		id = 5
// 	case "drivingOnNetwork":
// 		id = 6
// 	default:
// 		id = 0
// 	}
// 	return pb.Entry_StateId_value[]
// }
