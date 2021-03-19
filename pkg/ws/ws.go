package ws

import (
	"log"
	"time"

	"github.com/centrifugal/centrifuge-go"
)

type Ws struct {
	config *Config
	conn   *centrifuge.Client
}

func NewWS(config *Config) *Ws {
	return &Ws{
		config: config,
	}
}

func (w *Ws) Connect() error {
	//var err error
	w.conn = centrifuge.New(w.config.url, centrifuge.DefaultConfig())
	//w.conn, _, err = websocket.DefaultDialer.Dial(w.config.url, nil)
	//if err != nil {
	//	return err
	//}
	return nil
}

type LocalPublishHandler struct {
	ops        *uint64
	timeReport *uint64
	channel    string
}

func (s *LocalPublishHandler) OnPublish(c *centrifuge.Subscription, e centrifuge.PublishEvent) {
	//log.Printf("recv: %v", e.Data)
	now := time.Now()
	*s.timeReport = uint64(now.UnixNano())
	//fmt.Printf("Channel: %s received %d \n", s.channel, now.UnixNano())
	//atomic.AddUint64(s.ops, 1)
}

/* func (s *LocalPublishHandler) Counter() {
	log.Printf("recv: %v", e.Data)
}
*/
func (w *Ws) Run(stop chan int, ops *uint64, sender bool, timeReport *uint64) {
	w.conn.SetToken(w.config.token)
	err := w.conn.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer w.conn.Disconnect()
	if sender {
		time.Sleep(10 * time.Millisecond)
		now := time.Now()
		*timeReport = uint64(now.UnixNano())
		_, err = w.conn.Publish("public:"+w.config.channel, []byte("test"))
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("Channel: %s public %d \n", "public:"+w.config.channel, now.UnixNano())

		//fmt.Printf("Channel: %s send \n", w.config.channel)
	} else {
		sub, _ := w.conn.NewSubscription("public:" + w.config.channel)
		handler := &LocalPublishHandler{ops, timeReport, w.config.channel}
		sub.OnPublish(handler)
		_ = sub.Subscribe()
	}
	<-stop
}
