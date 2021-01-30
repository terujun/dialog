/*
The Main Package starts "Echo Server"
It Controls the behavior of app according to the data
received from Slack.
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/terujun/dialog/pkg/meal-slack-bot/config"
	"github.com/terujun/dialog/pkg/meal-slack-bot/file"
)

//config.jsonの読み込み処理
func readConfig(configsDirPath string, token string) (config.Config, error) {

	//config構造体を宣言
	config := config.Config{}
	configFilePath := filepath.Join(configsDirPath, "config.json")

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
		configsDirPath = "/home/go/dialog/configs/config/"
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
	log.Printf("listening on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}

func gateway(c echo.Context, appConfig config.Config, configsDirPath string) error {
	//payloadJSON := c.FormValue("payload")
	//var payload interface{}

	//受信確認あとで消す。
	log.Printf("受信したで")

	/*
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

		//type別にコールバックIDを取得する
		var iCallbackID interface{}
		switch requestType {
		case "shortcut":
			iCallbackID, _ = jsonpointer.Get(payload, "/callback_id")
		case "view_submit":
			iCallbackID, _ = jsonpointer.Get(payload, "/view/callback_id")
		}
		callbackID := iCallbackID.(string)

		//callbackID種類ごとの処理を記載
		if len(callbackID) > 0 {
			switch callbackID {
			case "meal_reg_call":
				return HandleOpenHydrationForm(c, appConfig, configsDirPath, payload)
			default:
				c.Echo().Logger.Warn("Unrecognized callbackID:", callbackID)
			}

		}
		return c.String(http.StatusForbidden, "Error")
	*/

	//とりあえずOK.あとで消す
	return c.String(http.StatusOK, "Ok")
}

/*適宜追加していく
func HandleOpenHydrationForm(c echo.Context, appConfig config.Config, configsDirPath string, payload interface{}) error {

	//非同期処理を記載
		go func (){
			slackRepo := &repositories.SlackRepository{
				Token:	appConfig.Slack.Token,
				ViewDirPath:filepath.Join(configsDirPath,"views"),
			}

		}
	// ステータス200でレスポンスを返す。
	return c.String(http.StatusOK, "Ok")
}

func (repo *SlackRepository) OpenHydrationAddView(triggerID string) ([]byte, error) {
	var err error
	var res []byte
	var requestparam, view interface{}

	viewpath := repositories.SlackRepository
}*/
