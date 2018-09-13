package main


import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"math/big"
	"github.com/golang/glog"
	"github.com/gangwn/dvc/pkg/eth/events"
					"github.com/gangwn/dvc/pkg/net/basic"
	"github.com/gangwn/dvc/pkg/eth/basic"
	"strconv"
	"encoding/json"
				"github.com/ethereum/go-ethereum/common"
	"github.com/gangwn/dvc/pkg/eth"
	"github.com/gangwn/dvc/tests/client/ws"
	"github.com/gangwn/dvc/pkg/protocol"
	"github.com/satori/go.uuid"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/gangwn/dvc/pkg/protocol/pb"
	"github.com/gangwn/dvc/internal/pkg/ccs/config"
	"github.com/gangwn/dvc/pkg/util"
)

var ethClient eth.Client
var wsClient *ws.WSClient
var node *basic.BasicNode
var userId string
var ccsId peer.ID

type Request struct {
	Type string `json:"type"`
	ScheduleConference scheduleConference `json:"scheduleConference""`
	ListCCS listCCSJson `json:"setCCS""`
	GetCCS getCCSJson `json:"getCCS""`
	Join joinJson `json:"join""`
	End endJson `json:"end""`
	Leave leaveJson `json:"leave""`

}

type scheduleConference struct {
	ConfId string `json:"confId"`
	ConfName string `json:"confName""`
	Duration int64 `json:"duration"`
	StartTime int64 `json:"startTime"`
	Invitees []string `json:"invitees"`
}

type listCCSJson struct {
	ConfId string `json:"confId"`
	CCS string `json:"ccs""`
}

type getCCSJson struct {
	ConfId string `json:"confId"`
}

type joinJson struct {
	ConfId string `json:"confId"`
	CCS ccsJson `json:"ccs"`
}

type endJson struct {
	ConfId string `json:"confId"`
}

type leaveJson struct {
	ConfId string `json:"confId"`
}

type ccsJson struct {
	IP string `json:"ip"`
	Port *big.Int `json:"port"`
	PeerId string `json:"peerId"`

}

var HTTPAddr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
},} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	wsClient = &ws.WSClient{conn, make(chan []byte),handleMessage}

	go wsClient.Read()
	go wsClient.Write()
}

func handleMessage(message []byte) {
	glog.Infof("Receive message, request: %v", string(message))

	request := &Request{}
	json.Unmarshal(message, request)
	handleRequest(request)
}

func handleRequest(request *Request) {
	if request.Type == "scheduleConference" {

		invitees :=[] common.Address{common.HexToAddress(request.ScheduleConference.Invitees[0])}

		ethClient.ScheduleConference(request.ScheduleConference.ConfId, request.ScheduleConference.ConfName,
			big.NewInt(request.ScheduleConference.StartTime), big.NewInt(request.ScheduleConference.Duration), invitees)
	} else if  request.Type == "listCCS" {
		err, ccss := ethClient.ListAllCCS()
		if err != nil {
			glog.Errorf("Error Get CCS: %v", err)
		} else {
			handleListAllCCS(ccss);
		}
	} else if  request.Type == "setCCS" {
		ethClient.NewJob(request.ListCCS.ConfId, common.HexToAddress(request.ListCCS.CCS))
	} else if request.Type == "getCCS" {
		_, ccs := ethClient.GetCCS(request.GetCCS.ConfId)
		handleGetAllCCS(request.GetCCS.ConfId, ccs)
	} else if request.Type == "join" {
		//lisAddr :=[] config.ListenAddress{{"127.0.0.1", 8101}}
		//
		//addr := util.NewMultiaddr(lisAddr)
		//
		//_, priv, err := util.LoadKeyPair("/Users/gangwan/client")
		//if err != nil {
		//	return
		//}
		//
		//node := basic.NewBasicNode(addr, priv)
		//if node == nil {
		//	return
		//}
		//
		//node.SetMessageHandler(MessageHandler)
		var err error
		remote := "/ip4/" + request.Join.CCS.IP + "/tcp/" + strconv.Itoa(int(request.Join.CCS.Port.Int64())) + "/ipfs/" + request.Join.CCS.PeerId
		ccsId, err = node.Connect(remote)
		if err != nil {
			glog.Errorf("Error connect CCS: %v", err)
		}

		userId = uuid.NewV4().String()

		pbPack := &protocol.ProtobufPack{}
		message := pbPack.CreateJoinConferenceRequest(request.Join.ConfId, userId,  userId)
		node.SendMessage(ccsId, message)
	} else if request.Type == "end" {
		glog.Infof("end conference, confid: %v, ccsid: %v, userid: %v", request.End.ConfId, ccsId.Pretty(), userId)
		pbPack := &protocol.ProtobufPack{}
		message := pbPack.CreateEndConferenceRequest(request.End.ConfId, userId)
		err := node.SendMessage(ccsId, message)
		if err != nil {
			glog.Errorf("Error send end request to CCS: %v", err)
		}
	} else if request.Type == "leave" {
		pbPack := &protocol.ProtobufPack{}
		message := pbPack.CreateLeaveConferenceRequest(request.Leave.ConfId, userId)
		node.SendMessage(ccsId, message)
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	flag.Set("logtostderr", "true")
	flag.Parse()

	lisAddr :=[] config.ListenAddress{{"127.0.0.1", 8101}}
	ethAccount := "0x673bf560ba83e2fc3bbb0c15c341918176785c7c"
	keystoreDir := "/Users/gangwan/blockchain/testnet/keystore"
	gethRPCPath := "/Users/gangwan/blockchain/testnet/geth.ipc"
	contractAddr := "0x15253be34b7ee592970dfc09d5b6a9d3205c6c34"
	cfgGasLimit := "0x8000000000"
	cfgGasPrice := 20
	password := "pass"

	addr := util.NewMultiaddr(lisAddr)

	_, priv, err := util.LoadKeyPair("/Users/gangwan/client")
	if err != nil {
		return
	}

	node = basic.NewBasicNode(addr, priv)
	if node == nil {
		return
	}

	node.SetMessageHandler(MessageHandler)

	//id, err := node.Connect("/ip4/127.0.0.1/tcp/8100/ipfs/QmVzDgD8kiU57EPSZBYV45bzHYEiRXhQMats1ENHMfELVR")
	//
	//pbPack := &protocol.ProtobufPack{}
	//message := pbPack.CreateJoinConferenceRequest(uuid.NewV4().String(), uuid.NewV4().String(),  "client1")
	//node.SendMessage(id, message)


	glog.Infof("EthAccount: %s, KeystoreDir: %s, GethRPCPath: %s, ContractAddr: %s", ethAccount, keystoreDir, gethRPCPath, contractAddr)
	ethClient, err = basiceth.NewBasicClient(ethAccount, keystoreDir, gethRPCPath, contractAddr)
	if err != nil {
		glog.Errorf("Create basic eth client fail: ", err)
		return
	}

	gasLimit, err := strconv.ParseUint(cfgGasLimit, 0, 64)
	if err != nil {
		glog.Errorf("Error parse gas limit: ", err)
		return
	}

	var gasPrice *big.Int
	if cfgGasPrice > 0 {
		gasPrice = big.NewInt(int64(cfgGasPrice))
	}

	glog.Infof("Initialze ethClient, gasLimit: %v, gasPrice: %v", gasLimit, *gasPrice)

	err = ethClient.SetUp(password, gasLimit, gasPrice)
	if err != nil {
		glog.Errorf("Error initialize ethClient: ", err)
		return
	}

	conferenceScheduledEvent := events.NewConferenceScheduledEvent(ethClient)
	if conferenceScheduledEvent == nil {
		glog.Errorf("create conferenceScheduledEvent fail")
		return
	}

	conferenceScheduledEvent.Handler = ConferenceScheduledEventHandler
	err = conferenceScheduledEvent.Listen()
	if err != nil {
		glog.Errorf("Error listen conferenceScheduledEvent: ", err)
		return
	}

	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe(*HTTPAddr, nil))
}

type ConferenceScheduledEventMessage  struct {
	Type string
	ConferenceScheduledEvent *events.ConferenceScheduledEventData
}

func ConferenceScheduledEventHandler(eventData *events.ConferenceScheduledEventData) {
	message := &ConferenceScheduledEventMessage{}
	message.Type = "ConferenceScheduledEvent"
	message.ConferenceScheduledEvent = eventData

	data, _ := json.Marshal(message)

	glog.Infof("ConferenceScheduledEventHandler: %v", message)

	wsClient.Send <- data
}

type ListAllCCSMessage  struct {
	Type string
	CCSS []eth.CCS
}

func handleListAllCCS(ccss []eth.CCS) {
	message := &ListAllCCSMessage{}
	message.Type = "ListAllCCSResponse"
	message.CCSS = ccss

	data, _ := json.Marshal(message)

	glog.Infof("handleListAllCCS: %v", message)

	wsClient.Send <- data
}

type GetCCSMessage struct {
	Type string
	CCS eth.CCS
	ConfId string
}

func handleGetAllCCS(confId string, ccs eth.CCS) {
	message := &GetCCSMessage{}
	message.Type = "GetCCSResponse"
	message.CCS = ccs
	message.ConfId = confId

	data, _ := json.Marshal(message)

	glog.Infof("handleGetAllCCS: %v", message)

	wsClient.Send <- data
}

type RspMessage  struct {
	Type string
	Result int32
}

type EndIndicationMessage  struct {
	Type string
	ConfId string
}

type RosterMessage  struct {
	Type string
	Participants         []*dvc_protocol.Participant
}

func  MessageHandler(id peer.ID, message *dvc_protocol.DVCMessage) {
	glog.Infof("Receive message type %s from id %s", message.Type, id.Pretty() )

	if message.GetType() == dvc_protocol.DVCMessage_JoinConferenceResponse {
		rsp := message.GetJoinConfRsp()

		message := &RspMessage{}
		message.Type = "JoinConfRsp"
		message.Result = rsp.Result

		data, _ := json.Marshal(message)
		wsClient.Send <- data
	} else if message.GetType() == dvc_protocol.DVCMessage_EndConferenceResponse {
		endRsp := message.GetEndConfRsp()

		message := &RspMessage{}
		message.Type = "EndConfRsp"
		message.Result = endRsp.Result

		data, _ := json.Marshal(message)
		wsClient.Send <- data
	} else if message.GetType() == dvc_protocol.DVCMessage_LeaveConferenceResponse {
		leaveRsp := message.GetLeaveConfRsp()

		message := &RspMessage{}
		message.Type = "LeaveConfRsp"
		message.Result = leaveRsp.Result

		data, _ := json.Marshal(message)
		wsClient.Send <- data

	} else if message.GetType() == dvc_protocol.DVCMessage_EndConferenceIndication {
		endInd := message.GetEndConfInd()

		message := &EndIndicationMessage{}
		message.Type = "EndConfIndication"
		message.ConfId = endInd.ConferenceId

		data, _ := json.Marshal(message)
		wsClient.Send <- data
	} else if message.GetType() == dvc_protocol.DVCMessage_RosterMessage {
		rosterMessage := message.GetRosterMsg()

		message := &RosterMessage{}
		message.Type = "RosterMessage"
		message.Participants = rosterMessage.Participants

		data, _ := json.Marshal(message)
		wsClient.Send <- data
	}
}