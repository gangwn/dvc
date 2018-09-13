package main

import (
	"github.com/gangwn/dvc/pkg/util"
	"github.com/gangwn/dvc/pkg/net/basic"
	"github.com/gangwn/dvc/internal/pkg/ccs/config"
	"github.com/satori/go.uuid"
		"fmt"
	"bufio"
	"os"
	"github.com/golang/glog"
	"github.com/gangwn/dvc/pkg/eth/basic"
	"strconv"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"time"
	"strings"
	"github.com/gangwn/dvc/pkg/eth/events"
	"flag"
	"github.com/gangwn/dvc/pkg/eth"
)

var ConfId string

var ccsMap map[common.Address]eth.CCS

var ccs eth.CCS

func main() {
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

	ccsMap = make(map[common.Address]eth.CCS)

	addr := util.NewMultiaddr(lisAddr)

	_, priv, err := util.LoadKeyPair("/Users/gangwan/client")
	if err != nil {
		return
	}

	node := basic.NewBasicNode(addr, priv)
	if node == nil {
		return
	}

	//id, err := node.Connect("/ip4/127.0.0.1/tcp/8100/ipfs/QmVzDgD8kiU57EPSZBYV45bzHYEiRXhQMats1ENHMfELVR")
	//
	//pbPack := &protocol.ProtobufPack{}
	//message := pbPack.CreateJoinConferenceRequest(uuid.NewV4().String(), uuid.NewV4().String(),  "client1")
	//node.SendMessage(id, message)

	glog.Infof("EthAccount: %s, KeystoreDir: %s, GethRPCPath: %s, ContractAddr: %s", ethAccount, keystoreDir, gethRPCPath, contractAddr)
	ethClient, err := basiceth.NewBasicClient(ethAccount, keystoreDir, gethRPCPath, contractAddr)
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

	err = conferenceScheduledEvent.Listen()
	if err != nil {
		glog.Errorf("Error listen  conferenceScheduledEvent: ", err)
		return
	}

	go readCommands(ethClient)

	//schedule(nil, ethClient)
	select {}
}

func readCommands(client *basiceth.BasicClient) {
	glog.Infof(`Start to test the dvc system
  schedule   Schedule a conference
  join       Join a conference
  leave      Leave conference
  setCCS     set a CCS for a conference
  getCCS     get all CCS`)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Input command: ")
		command := readString(reader)

		if strings.Compare(command, "schedule") == 0 ||  strings.Compare(command, "s") == 0 {
			schedule(reader, client)
		} else if strings.Compare(command, "setCCS") == 0 {
			setCCS(reader, client)
		} else if strings.Compare(command, "getCCS") == 0 {
			getAllCCS(reader, client)
		}
	}
}

func readString(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}


func schedule(reader *bufio.Reader, client *basiceth.BasicClient)  {
	fmt.Print("Conference name:")
	//confName := readString(reader)
	confName := "test"

	fmt.Print("Duration (mins):")
	//durationStr := readString(reader)
	fmt.Print("Invitee:")
	//invitee := readString(reader)

	confId := uuid.NewV4().String()
	startTime := big.NewInt(time.Now().Unix())
	//duration, _ := strconv.ParseInt(durationStr, 0, 64)
	duration:= int64(121)
	//invitees :=[] common.Address{common.HexToAddress(invitee)}
	invitees :=[] common.Address{common.HexToAddress("0xc48910bb385b65150dc96ca3e3dd8687b770d106")}

	err := client.ScheduleConference(confId, confName, startTime, big.NewInt(duration), invitees)
	if err != nil {
		glog.Errorf("Error Schedule Conference: %v", err)
	}

	ConfId = confId
}


func setCCS(reader *bufio.Reader, client *basiceth.BasicClient) {
	fmt.Print("Set CCS:")

	ccsAddr := readString(reader)
	ccs = ccsMap[common.HexToAddress(ccsAddr)]
	client.NewJob(ConfId, ccs.Addr)
}

func getAllCCS(reader *bufio.Reader, client *basiceth.BasicClient) {
	fmt.Print("Get all CCS")

	confId := readString(reader)
	if confId == "" {
		err, ccss := client.ListAllCCS()
		if err != nil {
			glog.Errorf("Error Get CCS: %v", err)
		} else {
			for i := range ccss {
				ccsMap[ccss[i].Addr] = ccss[i]
				glog.Infof("available css, address :%v, ip: %v, port:%v", ccss[i].Addr.Hex(), ccss[i].IP, ccss[i].Port)
			}
		}
	} else {
		err, ccs := client.GetCCS(confId)
		if err != nil {
			glog.Errorf("Error Get CCS: %v", err)
		} else {
			glog.Infof("available css for conf, address :%v, ip: %v, port:%v", ccs.Addr.Hex(), ccs.IP, ccs.Port)
		}
	}

}