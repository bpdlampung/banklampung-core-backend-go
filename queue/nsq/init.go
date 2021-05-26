package nsq

import (
	"banklampung-core/logs"
	"fmt"
	goNsq "github.com/nsqio/go-nsq"
)

var producer *goNsq.Producer

type NsqConfig struct {
	channel           string
	nsqAddress        string
	nsqLookupdAddress string
	logger            logs.Collections
}

func InitConfig(nsqAddress, nsqLookupAddress, channel string, logger logs.Collections) NsqConfig {
	logger.Info("NSQ Generate Config --Success")

	return NsqConfig{
		channel:           fmt.Sprintf("%s-channel", channel),
		nsqAddress:        nsqAddress,       // NsqAddress for produce to nsqd
		nsqLookupdAddress: nsqLookupAddress, // NsqLookupdAddress for Get Data from nsqd
		logger:            logger,
	}
}
