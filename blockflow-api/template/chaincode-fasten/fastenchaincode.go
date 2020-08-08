/**/

package main
import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

 
type SampleChaincode struct {
}

type Printer struct {
	ObjectType 			string  `json:"docType"`
	Id 					string  `json:"id"`
	Type				string  `json:"type"`
	Timestamp			int     `json:"timestamp"`
	Temp_tool1_actual 	float64 `json:"temp_tool1_actual"`
	Temp_tool2_actual 	float64 `json:"temp_tool2_actual"`
	Temp_bed_actual 	float64 `json:"temp_bed_actual"`
	File_name 			string  `json:"file_name"`
	Printer3d_state		string  `json:"printer3d_state"`
	Progress			float64 `json:"progress"`
	Time_left			float64 `json:"time_left"`
	Time_elapsed		float64 `json:"time_elapsed"`
	Temp_tool1_goal		float64 `json:"temp_tool1_goal"`
	Temp_tool2_goal		float64 `json:"temp_tool2_goal"`
	Temp_bed_goal		float64 `json:"temp_bed_goal"`
}

type Navigation struct {
	ObjectType 			string  `json:"docType"`
	Id 					string  `json:"id"`
	Type				string  `json:"type"`
	Time 				float64 `json:"time"`
	Skill				int     `json:"skill"`
	X 					float64 `json:"x"`
	Y 					float64 `json:"y"`
	Theta 				float64 `json:"theta"`
	V 					float64 `json:"v"`
	W 					float64 `json:"w"`
}

type PickAndPlace struct {
	ObjectType 			string  `json:"docType"`
	Id 					string  `json:"id"`
	Type				string  `json:"type"`
	Time 				float64 `json:"theta"`
	Skill 				float64 `json:"skill"`
	Joint1 				float64 `json:"joint1"`
	Joint2 				float64 `json:"joint2"`
	Joint3 				float64 `json:"joint3"`
	Joint4 				float64 `json:"joint4"`
	Joint5 				float64 `json:"joint5"`
	Joint6 				float64 `json:"joint6"`
	Velocity1 			float64 `json:"velocity1"`
	Velocity2 			float64 `json:"velocity2"`
	Velocity3 			float64 `json:"velocity3"`
	Velocity4 			float64 `json:"velocity4"`
	Velocity5 			float64 `json:"velocity5"`
	Velocity6 			float64 `json:"velocity6"`
	X		 			float64 `json:"x"`
	Y		 			float64 `json:"y"`
	Z		 			float64 `json:"z"`
	Ax		 			float64 `json:"ax"`
	Ay		 			float64 `json:"ay"`
	Az		 			float64 `json:"az"`
	Tx		 			float64 `json:"tx"`
	Ty		 			float64 `json:"ty"`
	Tz		 			float64 `json:"tz"`
	Cx		 			float64 `json:"cx"`
	Cy		 			float64 `json:"cy"`
	Cz		 			float64 `json:"cz"`
	Cax		 			float64 `json:"cax"`
	Cay		 			float64 `json:"cay"`
	Caz		 			float64 `json:"caz"`
	GripperPosition		int     `json:"gripperPosition"`
	GripperStatus		int     `json:"gripperStatus"`
}

type Part struct {
	ObjectType 			string  `json:"docType"`
	Id 					string  `json:"id"`
	Type				string  `json:"type"`
	Name 				string  `json:"name"`
	Description			string  `json:"description"`
}

type User struct {
	ObjectType 			string  `json:"docType"`
	Id 					string  `json:"id"`
	Type				string  `json:"type"`
	Name 				string  `json:"name"`
	Description			string  `json:"description"`
	Login 				string  `json:"login"`
	Senha 				string  `json:"senha"`
}

type Admin struct {
	ObjectType 			string  `json:"docType"`
	Id 					string  `json:"id"`
	Type				string  `json:"type"`
	Name 				string  `json:"name"`
	Description			string  `json:"description"`
	Login 				string  `json:"login"`
	Senha 				string  `json:"senha"`
}


// Init function
func (s *SmartContractBlockflow) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke funcion
func (s *SmartContractBlockflow) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()
	
	if function == "recordPrinter" {
		return s.recordPrinter(APIstub, args)
	}else if function == "queryPrinter" {
		return s.queryPrinter(APIstub, args)
	}else if function == "updatePrinter" {
		return s.updatePrinter(APIstub, args)
	}
	
	
		
	return shim.Error("Invalid Smart Contract function name.")
}


// Record printer values
func (s *SampleChaincode) recordPrinter(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {


	if len(args) != 14 {
		return shim.Error("Incorrect number of arguments for recordPrinter")
	}

	ObjectType := args[0]
	Id:= args[1] 
	Type:= "printer3d,"
	Timestamp:= args[2]
	Temp_tool1_actual:= args[3]
	Temp_tool2_actual:= args[4]
	Temp_bed_actual:= args[5]
	File_name:= args[6]
	Printer3d_state:= args[7]
	Progress:= args[8]
	Time_left:= args[9]
	Time_elapsed:= args[10]
	Temp_tool1_goal:= args[11]
	Temp_tool2_goal:= args[12]
	Temp_bed_goal:= args[13]
	

	printer3d := &Printer{ObjectType, Id, Type, Timestamp, Temp_tool1_actual, Temp_tool2_actual, Temp_bed_actual, File_name, Printer3d_state, Progress, Time_left, Time_elapsed, Temp_tool1_goal, Temp_tool2_goal, Temp_bed_goal}
	
	printer3dAsBytes, _ := json.Marshal(printer3d)

	err := APIstub.PutState(Id, printer3dAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record printer catch: %s", args[0]))
	}

	return shim.Success(nil)
}

// Query printer by ID
func (s *SampleChaincode) queryPrinter(ctx contractapi.TransactionContextInterface, printer3dId string) (*Printer, error) {
	printer3dAsBytes, err := ctx.GetStub().GetState(printer3dId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if printer3dAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", printer3dId)
	}

	printer := new(Printer)
	_ = json.Unmarshal(printer3dAsBytes, printer)

	return printer, nil
}


// Update printer values
func (s *SampleChaincode) updatePrinter(ctx contractapi.TransactionContextInterface, args []string) error {
	printer3dAsBytes, err := s.queryPrinter(ctx, args[1])

	if err != nil {
		return err
	}

	if len(args) != 14 {
		return shim.Error("Incorrect number of arguments for updatePrinter")
	}

	ObjectType := args[0]
	Id:= args[1] 
	Type:= "printer3d,"
	Timestamp:= args[2]
	Temp_tool1_actual:= args[3]
	Temp_tool2_actual:= args[4]
	Temp_bed_actual:= args[5]
	File_name:= args[6]
	Printer3d_state:= args[7]
	Progress:= args[8]
	Time_left:= args[9]
	Time_elapsed:= args[10]
	Temp_tool1_goal:= args[11]
	Temp_tool2_goal:= args[12]
	Temp_bed_goal:= args[13]
	

	printer3d := &Printer{ObjectType, Id, Type, Timestamp, Temp_tool1_actual, Temp_tool2_actual, Temp_bed_actual, File_name, Printer3d_state, Progress, Time_left, Time_elapsed, Temp_tool1_goal, Temp_tool2_goal, Temp_bed_goal}
	
	printer3dAsBytes, _ := json.Marshal(printer3d)

	err := APIstub.PutState(Id, printer3dAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update printer catch: %s", args[1]))
	}

	return shim.Success(nil)
}


//Main
func main() {
	
	err := shim.Start(new(SampleChaincode))
	
	if err != nil {
	
		fmt.Println(“Could not start SampleChaincode”)
	
	} else {

		fmt.Println(“SampleChaincode successfully started”)

	}

}
	