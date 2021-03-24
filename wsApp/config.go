package wsApp

import (
	"log"
	"os"
	"strconv"
)

type config struct {
	sessions                int
	create_user_timeout     int
	run_video_test          int
	video_test_timeout      int
	run_message_test        int
	message_test_timeout    int
	path_video_file         string
	path_image_preview_file string
	//wsConfig *ws.Config
}

func NewConfig() *config {
	var c config
	var err error
	var exists bool

	//   SESSIONS
	sessions, exists := os.LookupEnv("SESSIONS")
	if !exists {
		log.Fatalf("Variable SESSIONS is unknown")
	}
	c.sessions, err = strconv.Atoi(sessions)
	if err != nil {
		log.Fatalf("Variable SESSIONS has wrong format")
	}
	//	CREAT_USER_TIMEOUT
	create_user_timeout, exists := os.LookupEnv("CREATE_USER_TIMEOUT")
	if !exists {
		log.Fatalf("Variable CREATE_USER_TIMEOUT is unknown")
	}
	c.create_user_timeout, err = strconv.Atoi(create_user_timeout)
	if err != nil {
		log.Fatalf("Variable CREATE_USER_TIMEOUT has wrong format")
	}
	//	RUN_VIDEO_TEST
	run_video_test, exists := os.LookupEnv("RUN_VIDEO_TEST")
	if !exists {
		log.Fatalf("Variable RUN_VIDEO_TEST is unknown")
	}
	c.run_video_test, err = strconv.Atoi(run_video_test)
	if err != nil {
		log.Fatalf("Variable RUN_VIDEO_TEST has wrong format")
	}
	//	RUN_MESSAGE_TEST
	run_message_test, exists := os.LookupEnv("RUN_MESSAGE_TEST")
	if !exists {
		log.Fatalf("Variable RUN_MESSAGE_TEST is unknown")
	}
	c.run_message_test, err = strconv.Atoi(run_message_test)
	if err != nil {
		log.Fatalf("Variable RUN_MESSAGE_TEST has wrong format")
	}
	//  VIDEO_TEST_TIMEOUT
	if c.run_video_test == 1 {
		video_test_timeout, exists := os.LookupEnv("VIDEO_TEST_TIMEOUT")
		if !exists {
			log.Fatalf("Variable VIDEO_TEST_TIMEOUT is unknown")
		}
		c.video_test_timeout, err = strconv.Atoi(video_test_timeout)
		if err != nil {
			log.Fatalf("Variable VIDEO_TEST_TIMEOUT has wrong format")
		}
		c.path_video_file, exists = os.LookupEnv("PATH_VIDEO_FILE")
		if !exists {
			log.Fatalf("Variable PATH_VIDEO_FILE is unknown")
		}
		c.path_image_preview_file, exists = os.LookupEnv("PATH_IMAGE_PREVIEW_FILE")
		if !exists {
			log.Fatalf("Variable PATH_IMAGE_PREVIEW_FILE is unknown")
		}

	}
	//  MESSAGE_TEST_TIMEOUT
	if c.message_test_timeout == 1 {
		message_test_timeout, exists := os.LookupEnv("MESSAGE_TEST_TIMEOUT")
		if !exists {
			log.Fatalf("Variable MESSAGE_TEST_TIMEOUT is unknown")
		}
		c.message_test_timeout, err = strconv.Atoi(message_test_timeout)
		if err != nil {
			log.Fatalf("Variable MESSAGE_TEST_TIMEOUT has wrong format")
		}
	}

	//c.wsConfig = ws.NewConfig()
	return &c
}
