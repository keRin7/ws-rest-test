package wsApp

import (
	"log"
	"os"
	"strconv"
)

type config struct {
	sessions              int
	secret                string
	create_user_timeout   int
	wait_end_work_timeout int
	//wsConfig *ws.Config
}

func NewConfig() *config {
	var c config
	var err error
	var exists bool

	c.secret, exists = os.LookupEnv("SECRET")
	if !exists {
		log.Fatalf("Variable SECRET is unknown")
	}
	sessions, exists := os.LookupEnv("SESSIONS")
	if !exists {
		log.Fatalf("Variable SESSIONS is unknown")
	}
	c.sessions, err = strconv.Atoi(sessions)
	if err != nil {
		log.Fatalf("Variable SESSIONS has wrong format")
	}
	create_user_timeout, exists := os.LookupEnv("CREATE_USER_TIMEOUT")
	if !exists {
		log.Fatalf("Variable CREATE_USER_TIMEOUT is unknown")
	}
	c.create_user_timeout, err = strconv.Atoi(create_user_timeout)
	if err != nil {
		log.Fatalf("Variable CREATE_USER_TIMEOUT has wrong format")
	}
	wait_end_work_timeout, exists := os.LookupEnv("WAIT_END_WORK_TIMEOUT")
	if !exists {
		log.Fatalf("Variable CONN_TIMEOUT is unknown")
	}
	c.wait_end_work_timeout, err = strconv.Atoi(wait_end_work_timeout)
	if err != nil {
		log.Fatalf("Variable SESSIONS has wrong format")
	}
	//c.wsConfig = ws.NewConfig()
	return &c
}
