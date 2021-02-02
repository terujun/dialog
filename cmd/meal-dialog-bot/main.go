/*
The Main Package starts "Echo Server"
It Controls the behavior of app according to the data
received from Slack.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/mattn/go-jsonpointer"
	"github.com/terujun/dialog/pkg/meal-slack-bot/config"
	"github.com/terujun/dialog/pkg/meal-slack-bot/file"
	"github.com/terujun/dialog/pkg/meal-slack-bot/slack"
)

//config.jsonの読み込み処理
func readConfig(configsDirPath string, token string) (config.Config, error) {

	//config構造体を宣言
	config := config.Config{}
	configFilePath := filepath.Join(configsDirPath, "/config/config.json")

	//ファイルの存在確認
	if !file.FileExists(configFilePath) {
		return config, fmt.Errorf("Config file does not exist: %s", configFilePath)
	}

	//config.jsonファイルを読み込み
	//	jsonContent, err := ioutil.ReadFile(configFilePath)
	//	if err != nil {
	//	return config, err
	//}

	//config.json→config構造体へ読み込み
	/*今は読み込む物なし
	if err := json.Unmarshal(jsonContent, &config); err != nil {
		return config, err
	}
	*/
	//tokenセット
	config.Slack.Token = token

	return config, nil
}

func main() {
	//ポート情報取得
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		log.Printf("defaulting to port %s", port)
	}

	//token情報取得
	token := os.Getenv("TOKEN")
	if token == "" {
		//errorハンドリングを後で記載
		log.Printf("Tokenが設定されていません")
	}

	//config場所取得
	configsDirPath := os.Getenv("CONFIGDIRPATH")
	if configsDirPath == "" {
		configsDirPath = "/home/go/dialog/configs/"
		log.Printf("defaulting to configsDirPath %s", configsDirPath)
	}

	//config読み込み
	appConfig, err := readConfig(configsDirPath, token)
	if err != nil {
		log.Printf("error! %s", err)
	}

	//サーバスタート
	e := echo.New()
	e.POST("/postarticle", func(c echo.Context) error {
		return gateway(c, appConfig, configsDirPath)
	})
	/*テスト用
	e.GET("/test", func(c echo.Context) error {
		client := &http.Client{}
		req, _ := http.NewRequest("POST", "https://slack.com/api/views.open", nil)

		resp, err := client.Do(req)
		if err != nil {
			return c.String(http.StatusOK, "Error!!")
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		return c.String(http.StatusOK, string(string(body)))
	})*/
	log.Printf("listening on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}

func gateway(c echo.Context, appConfig config.Config, configsDirPath string) error {
	payloadJSON := c.FormValue("payload")
	var payload interface{}

	//中身確認用
	fmt.Println("payloadJSON is")
	fmt.Println(string([]byte(payloadJSON)))

	//payloadをJSONとして取得
	err := json.Unmarshal([]byte(payloadJSON), &payload)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	//type取得
	pointRequesttype, err := jsonpointer.Get(payload, "/type")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	requestType := pointRequesttype.(string)

	//受信確認あとで消す。
	log.Printf("受信したで %s", requestType)

	//type別にコールバックIDを取得する
	var iCallbackID interface{}
	switch requestType {
	case "shortcut":
		iCallbackID, _ = jsonpointer.Get(payload, "/callback_id")
	case "view_submit":
		iCallbackID, _ = jsonpointer.Get(payload, "/view/callback_id")
	}
	callbackID := iCallbackID.(string)
	fmt.Printf("callbackID is  %s", callbackID)

	//callbackID種類ごとの処理を記載
	if len(callbackID) > 0 {
		switch callbackID {
		case "meal_reg_call":
			//fmt.Printf("callbackID is %s", callbackID)
			//triggerID, _ := jsonpointer.Get(payload, "/trigger_id")
			return HandleOpenMealmodalForm(c, appConfig, configsDirPath, payload)
		case "mealreg_modal_receive":
			return HandleMealmodalFormSubmission(c, appConfig, configsDirPath, payload)
		default:
			c.Echo().Logger.Warn("Unrecognized callbackID:", callbackID)
		}

	}
	//return c.String(http.StatusForbidden, "Error")

	//とりあえずOK.あとで消す
	return c.String(http.StatusOK, "Ok")
}

func HandleOpenMealmodalForm(c echo.Context, appConfig config.Config, configsDirPath string, payload interface{}) error {

	//非同期処理を記載
	//go func(c echo.Context, configsDirPath string, payload interface{}) {
	slackRepo := &slack.SlackRepository{
		Token:        appConfig.Slack.Token,
		ViewsDirPath: filepath.Join(configsDirPath, "/views"),
	}

	//triggerID取得
	triggerID, err := jsonpointer.Get(payload, "/trigger_id")
	if err != nil {
		c.Echo().Logger.Error(err)
	}
	//fmt.Println("triggerID は")
	//fmt.Println(triggerID.(string))

	_, err = slackRepo.OpenMealmodalAddView(triggerID.(string))
	if err != nil {
		c.Echo().Logger.Error(err)
	}

	//}(c, configsDirPath, payload)
	// ステータス200でレスポンスを返す。
	return c.String(http.StatusOK, "Ok")
}

func HandleMealmodalFormSubmission(c echo.Context, appConfig config.Config, configsDirPath string, payload interface{}) error {

	//ここでひたすら欲しい情報取得

	return c.String(http.StatusOK, "")
}
