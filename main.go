package main

const url = "http://82.157.138.16:8091/CRAC/app/exam_exam/getExamList"

func main() {
	defaultCipherSuites = []uint16{0xc02f, 0xc030, 0xc02b, 0xc02c, 0xcca8, 0xcca9, 0xc013, 0xc009,
		0xc014, 0xc00a, 0x009c, 0x009d, 0x002f, 0x0035, 0xc012, 0x000a}
	httpclient = &http.Client{
		Timeout:   20 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{CipherSuites: append(defaultCipherSuites[8:], defaultCipherSuites[:8]...)}},
	}
	postdata := `{"req":{"province":"811","page_no":1,"page_size":10,"type":"A"},"req_meta":{"user_id":""}}`

	req, _ := http.NewRequest("POST", url, bytes.NewBufferString(postdata))
	// add header value
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh-Hans;q=0.9")
	req.Header.Set("timeZoneOffset", "-480")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("langType", "zh_CN")
	req.Header.Set("Cookie", "langType=zh_CN")
	req.Header.Set("X-Forwarded-For", clientip)
	req.Header.Set("Origin", "https://passport.eteams.cn")
	req.Header.Set("Referer", "https://passport.eteams.cn/login?v=111111")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone 16_2 like Mac OS F) AppleWebKit/605.1.15 (KHTML, like Gecko) weapp/4.3.21/public//1.0//2")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	resp, err := httpclient.Do(req)
	if err != nil {
		log.SetPrefix("[Error] ")
		log.Println("login send req failed.", err.Error())
	}
	defer resp.Body.Close()

	Cookies = resp.Cookies()
	all, err := io.ReadAll(resp.Body)

}