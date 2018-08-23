package eth

// Generate contract code via below commands
//go:generate abigen --sol  ../../api/contracts/WebexMeeting.sol --pkg contracts --out contracts/webex_meeting.go

type Client struct {

}

func (c *Client) Test() {

}