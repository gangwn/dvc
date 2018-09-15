
var socket;
var ccsMap = new Map();
var ccsList = new Array();
var confId;

function start_connect() {
	socket = new WebSocket("ws://127.0.0.1:8080/");

	socket.onmessage = function(data) {
        console.log("We get signal:");
        console.log(data);
        handle_data(data.data);
    };

    socket.onopen= function() {
        console.log("Socket opened.");
    };
}

function guid() {
  function s4() {
    return Math.floor((1 + Math.random()) * 0x10000)
      .toString(16)
      .substring(1);
  }
  return s4() + s4() + '-' + s4() + '-' + s4() + '-' + s4() + '-' + s4() + s4() + s4();
}

function schedule_conference() {
	var confId = guid();
	var confName = "test";
	var duration = 30000;
	var startTime = moment().unix();
	var invitees = new Array();
	invitees[0] = "0xc48910bb385b65150dc96ca3e3dd8687b770d106";

	var info = {
		"type":"scheduleConference",
		"scheduleConference":{
			"confId": confId,
			"confName": confName,
			"duration": duration,
			"startTime": startTime,
			"invitees": invitees
		}
	};

	socket.send(JSON.stringify(info));
}

function list_ccs() {
	var info = {
		"type":"listCCS"
	};

	socket.send(JSON.stringify(info));
}

function get_ccs() {
	var confId = prompt("Enter conf id:");
	var info = {
		"type":"getCCS",
		"getCCS": {
			"confId": confId
		}
	};

	socket.send(JSON.stringify(info));
}

function handle_data(data) {
	json_data = JSON.parse(data)

	console.log(json_data);

	if (json_data.Type == "ConferenceScheduledEvent") {
		
		event_data = json_data.ConferenceScheduledEvent;
		time = moment(moment.unix(event_data.StartTime)).format('YYYY-MM-DD HH:mm:ss');
		var notification = "Conference scheduled, creator: " + event_data.CreatorAddr
			+ ", ConfId: " + event_data.ConfId + ", Topic: " + event_data.Topic
			+ ", Start time:" + time + ", Duration: " + event_data.Duration + "s"
			+ ", Invitees: " + event_data.Invitees[0] ;

		 $("#append").append(notification);
		 $("#append").append("<br / >");
	} else if (json_data.Type == "ListAllCCSResponse") {
		event_data = json_data.CCSS;
		for (var i=0;i<event_data.length;i++)
		{
			var notification = "Available CSS, address: " + event_data[i].Addr
			+ ", IP: " + event_data[i].IP + ", Port: " + event_data[i].Port
			+ ", peerId: " + event_data[i].PeerId ;

			ccsList.push(event_data[i]);

		 $("#append").append(notification);
		 $("#append").append("<br / >");
		}

	} else if (json_data.Type == "GetCCSResponse") {
		var ccs = json_data.CCS;
		ccsMap.set(json_data.ConfId, ccs);

		var notification = "Get CSS, address: " + json_data.CCS.Addr
			+ ", IP: " + json_data.CCS.IP + ", Port: " + json_data.CCS.Port
			+ ", ConfId: " + json_data.ConfId;

		$("#append").append(notification);
		$("#append").append("<br / >");
	} else if (json_data.Type == "JoinConfRsp") {
		var result = json_data.Result;

		var notification = "Join conference result: " + result;

		$("#append").append(notification);
		$("#append").append("<br / >");
	}

	else if (json_data.Type == "LeaveConfRsp") {
		var result = json_data.Result;

		var notification = "Leave conference result: " + result;

		$("#append").append(notification);
		$("#append").append("<br / >");
	}

	else if (json_data.Type == "EndConfRsp") {
		var result = json_data.Result;

		var notification = "End conference result: " + result;

		$("#append").append(notification);
		$("#append").append("<br / >");
	}

}

function set_ccs() {
	var conf_id = prompt("Enter conf id:");
	var ccs = prompt("Enter CCS address:");

	console.log(conf_id);
	console.log(ccs);

	for (var i=0;i<ccsList.length;i++)
	{
		if(ccsList[i].Addr == ccs) {
			ccsMap[conf_id] = ccsList[i];
			break;
		}
	}

	
	var info = {
		"type":"setCCS",
		"setCCS": {
			"confId": conf_id,
			"ccs": ccs
		}
	};

	socket.send(JSON.stringify(info));
}

function join_conference() {
	var conf_id = prompt("Enter conf id:");
	var ccs = ccsMap.get(conf_id);

	confId = conf_id;

	var info = {
		"type":"join",
		"join": {
			"confId": conf_id,
			"ccs": {
				"ip": ccs.IP,
				"port": ccs.Port,
				"peerId": ccs.PeerId
			}
		}
	};

	socket.send(JSON.stringify(info));
}

function end_conference() {
	var info = {
		"type":"end",
		"end": {
			"confId": confId
		}
	};

	socket.send(JSON.stringify(info));
}

function leave_conference() {
	var info = {
		"type":"leave",
		"leave": {
			"confId": confId
		}
	};

	socket.send(JSON.stringify(info));
}
