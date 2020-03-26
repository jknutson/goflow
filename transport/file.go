package transport

import (
	"flag"
	"log"
	"os"

	flowmessage "github.com/cloudflare/goflow/v3/pb"
	"github.com/cloudflare/goflow/v3/utils"
)

var (
	// LogFile path to log file
	LogFile *string
)

// DefaultFileTransport is a transport for writing to a file.
type DefaultFileTransport struct {
}

// RegisterFileFlags (TODO: document this)
func RegisterFileFlags() {
	LogFile = flag.String("log.file", "goflow.log", "Path to logfile")
}

// WriteLogLine writes a flowMessage to a new line in a local file.
func WriteLogLine(flowMessage *flowmessage.FlowMessage) {
}

// Publish writes FlowMessages to a File.
func (t DefaultFileTransport) Publish(msgs []*flowmessage.FlowMessage) {
	f, err := os.OpenFile("goflow.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	for _, msg := range msgs {
		if _, err := f.Write([]byte(utils.FlowMessageToString(msg))); err != nil {
			log.Fatal(err)
		}
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
