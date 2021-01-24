package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
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

func SendSlackModal(webhookurl string, TriggerID string) error {
	//trigger iD登録
	modalcontent := `{"trigger_id":` + TriggerID + `,"view":{"title":{"type":"plain_text","text":"Regiteryourmeal"},"submit":{"type":"plain_text","text":"Submit"},"blocks":[{"type":"input","element":{"type":"plain_text_input","action_id":"imageURL","placeholder":{"type":"plain_text","text":"https://〜Pasteyouruploadedimage'sURL"}},"label":{"type":"plain_text","text":"ImageURL"}},{"type":"context","elements":[{"type":"plain_text","text":"事前にslackへ画像up→画像タップ→その他タップ→リンクをコピー","emoji":true}]},{"type":"input","element":{"type":"radio_buttons","options":[{"text":{"type":"plain_text","text":"美味すぎ","emoji":true},"value":"umai"},{"text":{"type":"plain_text","text":"美味しくない","emoji":true},"value":"mazui"}],"action_id":"radio_buttons-action"},"label":{"type":"plain_text","text":"お勧めする？","emoji":true}},{"type":"input","element":{"type":"static_select","placeholder":{"type":"plain_text","text":"Selectagenre","emoji":true},"options":[{"text":{"type":"plain_text","text":"酒","emoji":true},"value":"sake"},{"text":{"type":"plain_text","text":"アジア","emoji":true},"value":"asia"},{"text":{"type":"plain_text","text":"フレンチ","emoji":true},"value":"french"},{"text":{"type":"plain_text","text":"イタリアン","emoji":true},"value":"italian"},{"text":{"type":"plain_text","text":"その他","emoji":true},"value":"other_genres"}],"action_id":"static_select-action"},"label":{"type":"plain_text","text":"ジャンル選択","emoji":true}},{"type":"input","element":{"type":"static_select","placeholder":{"type":"plain_text","text":"Selectawebsite","emoji":true},"options":[{"text":{"type":"plain_text","text":"Uber","emoji":true},"value":"uber"},{"text":{"type":"plain_text","text":"Doordash","emoji":true},"value":"doordash"},{"text":{"type":"plain_text","text":"Grubhub","emoji":true},"value":"grubhub"},{"text":{"type":"plain_text","text":"Yelp","emoji":true},"value":"yelp"},{"text":{"type":"plain_text","text":"EaterSF","emoji":true},"value":"eater"},{"text":{"type":"plain_text","text":"その他","emoji":true},"value":"other_websites"}],"action_id":"static_select-action"},"label":{"type":"plain_text","text":"注文サイト","emoji":true}},{"type":"input","element":{"type":"plain_text_input","action_id":"store_name","placeholder":{"type":"plain_text","text":"storename"}},"label":{"type":"plain_text","text":"店名"}},{"type":"input","element":{"type":"plain_text_input","multiline":true,"action_id":"plain_text_input-action"},"label":{"type":"plain_text","text":"一言","emoji":true}}],"type":"modal"}}`

	//modalcontent = []byte(strings.NewReplacer("¥n", "").Replace(string(modalcontent)))
	fmt.Println("modal request")
	fmt.Println(modalcontent)

	//Modalリクエスト作成
	req, err := http.NewRequest(http.MethodPost, webhookurl, strings.NewReader(modalcontent))
	//req, err := http.NewRequest(http.MethodPost, webhookurl, nil)
	if err != nil {
		return err
	}
	//httpヘッダ追加
	req.Header.Add("Content-Type", "application/json")

	dump, _ := httputil.DumpRequestOut(req, true)
	fmt.Println("dumpを出力するぞ")
	fmt.Printf("%s", dump)

	//body確認用
	httpsendbody, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(httpsendbody))

	//client := &http.Client{Timeout: 10 * time.Second}
	client := new(http.Client)

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

func postarticleHandler(w http.ResponseWriter, req *http.Request) {
	//受信確認用
	fmt.Println("I display req!!!!")

	//ボディ(JSON)取得
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	//URLデコード
	postarticlejsonbody, _ := url.QueryUnescape(string(body)[8:])
	//テスト出力
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

	//Modalの送信
	err = SendSlackModal(modalopenURL, appstart.TriggerID)
	if err != nil {
		log.Fatal(err)
	}

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
