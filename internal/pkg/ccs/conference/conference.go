package conference

import (
	"github.com/gangwn/dvc/pkg/protocol"
)

type Conference interface {
	JoinConference(message dvc_protocol.JoinConferenceMessage)
}