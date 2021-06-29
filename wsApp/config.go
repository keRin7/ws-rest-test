package wsApp

import (
	"ws-rest-test/pkg/rest_client"
)

type Config struct {
	Sessions                   int    `env:"SESSIONS"`
	Create_user_timeout        int    `env:"CREATE_USER_TIMEOUT"`
	Run_video_test             int    `env:"RUN_VIDEO_TEST"`
	Video_test_timeout         int    `env:"VIDEO_TEST_TIMEOUT"`
	Run_message_test           int    `env:"RUN_MESSAGE_TEST"`
	Message_test_timeout       int    `env:"MESSAGE_TEST_TIMEOUT"`
	Path_video_file            string `env:"PATH_VIDEO_FILE"`
	Path_image_preview_file    string `env:"PATH_IMAGE_PREVIEW_FILE"`
	Send_video_message_to_chat int    `env:"SEND_VIDEO_MESSAGE_TO_CHAT"`
	Report_timeout             int    `env:"REPORT_TIMEOUT"`
	Rest_client                *rest_client.Config
}

func NewConfig() *Config {
	return &Config{
		Rest_client: rest_client.NewConfig(),
	}
}
