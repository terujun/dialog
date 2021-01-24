package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	modalopenURL = "https://slack.com/api/views.open"
)

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

//modal作成用
/*
type Modalcontent struct {
	TriggerID string `json:"trigger_id"`
	View      struct {
		Title struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"title"`
		Submit struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"submit"`
		Blocks []struct {
			Type    string `json:"type"`
			Element struct {
				Type        string `json:"type"`
				ActionID    string `json:"action_id"`
				Placeholder struct {
					Type string `json:"type"`
					Text string `json:"text"`
				} `json:"placeholder"`
			} `json:"element,omitempty"`
			Label struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"label,omitempty"`
			Elements []struct {
				Type  string `json:"type"`
				Text  string `json:"text"`
				Emoji bool   `json:"emoji"`
			} `json:"elements,omitempty"`
			Element struct {
				Type    string `json:"type"`
				Options []struct {
					Text struct {
						Type  string `json:"type"`
						Text  string `json:"text"`
						Emoji bool   `json:"emoji"`
					} `json:"text"`
					Value string `json:"value"`
				} `json:"options"`
				ActionID string `json:"action_id"`
			} `json:"element,omitempty"`
			Label struct {
				Type  string `json:"type"`
				Text  string `json:"text"`
				Emoji bool   `json:"emoji"`
			} `json:"label,omitempty"`
			Element struct {
				Type        string `json:"type"`
				Placeholder struct {
					Type  string `json:"type"`
					Text  string `json:"text"`
					Emoji bool   `json:"emoji"`
				} `json:"placeholder"`
				Options []struct {
					Text struct {
						Type  string `json:"type"`
						Text  string `json:"text"`
						Emoji bool   `json:"emoji"`
					} `json:"text"`
					Value string `json:"value"`
				} `json:"options"`
				ActionID string `json:"action_id"`
			} `json:"element,omitempty"`
			Label struct {
				Type  string `json:"type"`
				Text  string `json:"text"`
				Emoji bool   `json:"emoji"`
			} `json:"label,omitempty"`
			Element struct {
				Type        string `json:"type"`
				Placeholder struct {
					Type  string `json:"type"`
					Text  string `json:"text"`
					Emoji bool   `json:"emoji"`
				} `json:"placeholder"`
				Options []struct {
					Text struct {
						Type  string `json:"type"`
						Text  string `json:"text"`
						Emoji bool   `json:"emoji"`
					} `json:"text"`
					Value string `json:"value"`
				} `json:"options"`
				ActionID string `json:"action_id"`
			} `json:"element,omitempty"`
			Label struct {
				Type  string `json:"type"`
				Text  string `json:"text"`
				Emoji bool   `json:"emoji"`
			} `json:"label,omitempty"`
			Element struct {
				Type      string `json:"type"`
				Multiline bool   `json:"multiline"`
				ActionID  string `json:"action_id"`
			} `json:"element,omitempty"`
			Label struct {
				Type  string `json:"type"`
				Text  string `json:"text"`
				Emoji bool   `json:"emoji"`
			} `json:"label,omitempty"`
		} `json:"blocks"`
		Type string `json:"type"`
	} `json:"view"`
}
*/

/*
func SendSlackModal(webhookurl string, TriggerID string) error {
	//trigger iD取得
	var modalcontent Modalcontent
	modalcontent.TriggerID = appstart.TriggerID
	//構造体からJSONを作成
	slackBody, _ := json.Marshal(modalcontent)
	//Modalようリクエスト作成
	req, err := http.NewRequest(http.MethodPost, webhookurl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}
	//httpヘッダ追加
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}

	//送信
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	//レスポンス確認
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil

}
*/

func postarticleHandler(w http.ResponseWriter, req *http.Request) {
	//受信確認用
	fmt.Fprintf(w, "caught")
	fmt.Println("I display req!!!!")

	//ボディ(JSON)取得
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	postarticlejsonbody, _ := url.QueryUnescape(string(body)[8:])
	fmt.Println("string new body")
	fmt.Println(string(postarticlejsonbody))
	if err != nil {
		log.Fatal(err)
	}

	// jsonを構造体へデコード
	var appstart Appstart
	err = json.Unmarshal([]byte(postarticlejsonbody), &appstart)
	if err != nil {
		log.Fatal(err)
	}

	//取得確認出力
	fmt.Println(appstart.TriggerID)

	//Modalの送信
	/*
		err = SendSlackModal(modalopenURL, apps)
		if err != nil {
			log.Fatal(err)
		}
	*/

}

func main() {
	fmt.Println("Hello!")

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		log.Printf("defaulting to port %s", port)
	}

	http.HandleFunc("/postarticle", postarticleHandler)
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
