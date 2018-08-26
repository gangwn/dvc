
package net

import (
	inet "github.com/libp2p/go-libp2p-net"
	protobufCodec "github.com/multiformats/go-multicodec/protobuf"
	"github.com/multiformats/go-multicodec"
	"bufio"
	"github.com/gangwn/dvc/pkg/protocol/pb"
	"errors"
)

type OutStream struct {
	Stream  inet.Stream
	encoder multicodec.Encoder
	w       *bufio.Writer
}


var ErrStream = errors.New("ErrorStream")

func NewOutStream(s inet.Stream) *OutStream {
	writer := bufio.NewWriter(s)
	enc := protobufCodec.Multicodec(nil).Encoder(writer)

	stream := &OutStream{s, enc, writer}
	return stream
}

func (s *OutStream) SendMessage(message *dvc_protocol.DVCMessage) error {
	if s == nil {
		return ErrStream
	}

	err := s.encoder.Encode(message)
	if err != nil {
		return err
	}

	err = s.w.Flush()
	if err != nil {
		return err
	}

	return nil
}






