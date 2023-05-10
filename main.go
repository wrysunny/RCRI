package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	searchurl = "http://82.157.138.16:8091/CRAC/app/exam_exam/getExamList"
	pushurl   = "https://push.showdoc.com.cn/server/api/push/3bd91d80f82bc69860d56a8526dd4cbd1721300601"
)

type Respose struct {
	Code int `json:"code"`
	Res  struct {
		List []struct {
			Id                string `json:"id"`
			Remarks           string `json:"remarks"`
			CreateDate        string `json:"createDate"`
			UpdateDate        string `json:"updateDate"`
			Area              string `json:"area"`
			State             string `json:"state"`
			ExamDate          string `json:"examDate"`
			MaxNum            int    `json:"maxNum"`
			Mode              string `json:"mode"`
			SignUpEndDate     string `json:"signUpEndDate"`
			SupplementEndDate string `json:"supplementEndDate"`
			Telephone         string `json:"telephone"`
			SignUpStartDate   string `json:"signUpStartDate"`
			Organizer         string `json:"organizer"`
			Type              string `json:"type"`
			AdviceId          string `json:"adviceId"`
			ExamId            string `json:"examId"`
			Province          struct {
				Id         string `json:"id"`
				Remarks    string `json:"remarks"`
				CreateDate string `json:"createDate"`
				UpdateDate string `json:"updateDate"`
				ParentIds  string `json:"parentIds"`
				Name       string `json:"name"`
				Sort       int    `json:"sort"`
				Code       string `json:"code"`
				Type       string `json:"type"`
				ParentId   string `json:"parentId"`
			} `json:"province"`
			City struct {
				Id         string `json:"id"`
				Remarks    string `json:"remarks"`
				CreateDate string `json:"createDate"`
				UpdateDate string `json:"updateDate"`
				ParentIds  string `json:"parentIds"`
				Name       string `json:"name"`
				Sort       int    `json:"sort"`
				Code       string `json:"code"`
				Type       string `json:"type"`
				ParentId   string `json:"parentId"`
			} `json:"city"`
			Street struct {
				Id         string `json:"id"`
				Remarks    string `json:"remarks"`
				CreateDate string `json:"createDate"`
				UpdateDate string `json:"updateDate"`
				ParentIds  string `json:"parentIds"`
				Name       string `json:"name"`
				Sort       int    `json:"sort"`
				Code       string `json:"code"`
				Type       string `json:"type"`
				ParentId   string `json:"parentId"`
			} `json:"street"`
			CanSingUp      string `json:"canSingUp"`
			SignUpState    string `json:"signUpState"`
			Weixin         string `json:"weixin"`
			Email          string `json:"email"`
			ExamArea       string `json:"examArea"`
			SignUpIsDelete string `json:"signUpIsDelete"`
			AdviceName     string `json:"adviceName"`
			ExamNumFlag    string `json:"examNumFlag"`
		} `json:"list"`
		Page      string `json:"page"`
		Count     int    `json:"count"`
		PageSize  int    `json:"page_size"`
		PageNo    int    `json:"page_no"`
		TotalPage int    `json:"total_page"`
	} `json:"res"`
	ResMeta struct {
	} `json:"res_meta"`
	Msg string `json:"msg"`
}

func main() {
	for {
		Search()
	}
}

func Search() {
	defaultCipherSuites := []uint16{0xc02f, 0xc030, 0xc02b, 0xc02c, 0xcca8, 0xcca9, 0xc013, 0xc009,
		0xc014, 0xc00a, 0x009c, 0x009d, 0x002f, 0x0035, 0xc012, 0x000a}
	httpclient := &http.Client{
		Timeout:   20 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{CipherSuites: append(defaultCipherSuites[8:], defaultCipherSuites[:8]...)}},
	}

	postdata := `{"req":{"province":"811","page_no":1,"page_size":10,"type":"A"},"req_meta":{"user_id":""}}`

	req, _ := http.NewRequest("POST", searchurl, bytes.NewBufferString(postdata))
	// add header value
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh-Hans;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	resp, err := httpclient.Do(req)
	if err != nil {
		log.SetPrefix("[Error] ")
		log.Println("request failed.", err.Error())
	}
	defer resp.Body.Close()

	all, err := io.ReadAll(resp.Body)
	var result Respose
	json.Unmarshal(all, &result)
	if resp.StatusCode == 200 && result.Code == 10000 {
		if len(result.Res.List) > 0 {
			for x := range result.Res.List {
				var cmd *exec.Cmd
				if os.Getenv("GOOS") == "windows" {
					cmd = exec.Command("cmd", "/c", "cls")
				} else {
					cmd = exec.Command("clear")
				}
				cmd.Stdout = os.Stdout
				cmd.Run()

				//createDate := result.Res.List[x].CreateDate
				//updateDate := result.Res.List[x].UpdateDate
				adviceName := result.Res.List[x].AdviceName               //标题
				city := result.Res.List[x].City.Name                      //城市
				area := result.Res.List[x].Area                           //地址
				examDate := result.Res.List[x].ExamDate                   //考试日期
				maxnum := result.Res.List[x].MaxNum                       //最大人数
				signUpEndDate := result.Res.List[x].SignUpStartDate       //报名截止日期
				supplementEndDate := result.Res.List[x].SupplementEndDate //补充材料截止日期
				fmt.Println(adviceName)
				fmt.Println("City:", city)
				fmt.Println("地址:", area)
				fmt.Println("考试日期:", examDate)
				fmt.Println("考试最大人数:", maxnum)
				fmt.Println("报名截止日期:", signUpEndDate)
				fmt.Println("补充材料截止日期:", supplementEndDate)
				fmt.Println()
				fmt.Println()
				if strings.Contains(city, "南京") {
					buf := ""
					buf += adviceName + "\r\n" + "城市:" + city + "\r\n" + "地址:" + area + "\r\n" +
						"考试日期:" + examDate + "\r\n" + "考试最大人数:" + strconv.Itoa(maxnum) + "\r\n" +
						"报名截止日期:" + signUpEndDate + "\r\n" + "补充材料截止日期:" + supplementEndDate + "\r\n"
					SendToWechat(buf)
				}
			}
			Next()
		} else {
			Next()
		}
	} else {
		Next()
	}
}

func Next() {
	time.Sleep(time.Hour * 24)
}

func SendToWechat(buf string) {
	var content = "title=" + "发现符合报名条件-" + time.Now().Format("01-02 15:04:05") + "&content=" + url.QueryEscape(buf)
	req, _ := http.NewRequest("POST", pushurl, bytes.NewBufferString(content))
	req.Header.Set("Accept-Language", "zh-CN,zh-Hans;q=0.9")
	req.Header.Set("langType", "zh_CN")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) weapp/4.3.21/public//1.0//2")
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	
	log.Println(resp.Body)
}
