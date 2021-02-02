/*
This packege describes the format for posting message.
*/
package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/mattn/go-jsonpointer"
	"github.com/terujun/dialog/pkg/meal-slack-bot/file"
)

//Appstart is struct of first message from slack
type Appstart struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ActionTs string `json:"action_ts"`
	Team     struct {
		ID     string `json:"id"`
		Domain string `json:"domain"`
	} `json:"team"`
	User struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		TeamID   string `json:"team_id"`
	} `json:"user"`
	IsEnterpriseInstall bool        `json:"is_enterprise_install"`
	Enterprise          interface{} `json:"enterprise"`
	CallbackID          string      `json:"callback_id"`
	TriggerID           string      `json:"trigger_id"`
}

// SlackRepository controls posts to Slack.
type SlackRepository struct {
	Token        string
	ViewsDirPath string
}

type Config struct {
	Token string `json:"token"`
}

func (repo *SlackRepository) OpenHydrationAddView(triggerID string) ([]byte, error) {
	var err error
	var resp []byte
	var requestParams, view interface{}

	fmt.Println("OpenHydrationAddViewまできたよ")
	//modalのviewフォーマットファイル確認
	viewPath := filepath.Join(repo.ViewsDirPath, "modal_view.json")
	if !file.FileExists(viewPath) {
		return resp, fmt.Errorf("View file does not exeist: %s", viewPath)
	}

	//triggerIDとviewを記載したjson→構造体への変換
	requestJSON := `{"trigger_id": "", "view":{}}`
	err = json.Unmarshal([]byte(requestJSON), &requestParams)
	if err != nil {
		return resp, err
	}

	//viewのJSONファイル読み込み
	viewJSON, err := ioutil.ReadFile(viewPath)
	if err != nil {
		return resp, err
	}

	//view構造体への変換
	err = json.Unmarshal([]byte(viewJSON), &view)
	if err != nil {
		return resp, err
	}

	//view→requestへの追加
	err = jsonpointer.Set(requestParams, "/view", view)
	if err != nil {
		return resp, err
	}

	//triggerID→requestへの追加
	err = jsonpointer.Set(requestParams, "trigger_id", triggerID)
	if err != nil {
		return resp, err
	}

	//requestParams→JSONへの変換
	requestParamasJSON, err := json.Marshal(requestParams)
	if err != nil {
		return resp, err
	}

	//request確認
	fmt.Println(string(requestParamasJSON))
	//res作成
	resp, err = PostJSON(repo.Token, "views.open", string(requestParamasJSON))

	return resp, err
}

func PostJSON(token string, command string, requestParamasJSON string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://slack.com/api/"+command, strings.NewReader(requestParamasJSON))
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
