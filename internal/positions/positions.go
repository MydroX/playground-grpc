package positions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/MydroX/geotwin-grpctest/internal/model"
	pb "github.com/MydroX/geotwin-grpctest/internal/positions/proto"
)

type positionsServer struct {
	pb.UnimplementedPositionServiceServer
}

func NewPositionsServer() positionsServer {
	return positionsServer{
		UnimplementedPositionServiceServer: pb.UnimplementedPositionServiceServer{},
	}
}

func (s positionsServer) GetEntriesPositionsFromScenarioId(_ *pb.GetEntriesPositionsFromScenarioIdRequest, stream pb.PositionService_GetEntriesPositionsFromScenarioIdServer) error {
	file, err := os.Open("vehicleOutputs.json")
	if err != nil {
		log.Fatalf("fail to open json file: %v", err)
	}

	var data model.SimulationResult

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("%v", err)
	}
	json.Unmarshal(byteValue, &data)

	entries := model.JSONEntriesToGRPC(data.StatesEntries)
	for i := 0; i < 7; i++ {
		entriesDuplicate := model.JSONEntriesToGRPC(data.StatesEntries)
		entries = append(entries, entriesDuplicate...)
	}

	for _, entry := range entries {
		err = stream.Send(
			&pb.GetEntriesPositionsFromScenarioIdResponse{
				Entry: entry,
			})
	}

	if err != nil {
		log.Fatalf("error in during streaming: %v", err)
	}

	log.Println("Data send")

	return nil
}

// func addEntries(data *model.SimulationResult, newEntriesNumber uint64) *model.SimulationResult {
// 	lastEntry := data.StatesEntries[len(data.StatesEntries)-1]

// 	for i := 0; i < int(newEntriesNumber)-1; i++ {
// 		entry := model.StateEntry{
// 			AgentId: ,
// 		}
// 	}
// }
