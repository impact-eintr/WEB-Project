package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"io/ioutil"
	"net/http"
)

var ip = widget.NewLabel("")
var position = widget.NewLabel("")
var isp = widget.NewLabel("")

type IpInfo struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    `json:"data"`
}

type Data struct {
	IP       string `json:"ip"`
	Position string `json:"pos"`
	Isp      string `json:"isp"`
}

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Demo")
	w.Resize(fyne.NewSize(600, 500))
	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayoutWithRows(2), info(GetIpInfo("")), query()))
	w.ShowAndRun()

}

func GetIpInfo(ip string) string {
	if len(ip) == 0 {
		return ""

	}

	url := fmt.Sprintf("http://v1.alapi.cn/api/ip?ip=%s&format=json", ip)

	resp, err := http.Get(url)
	if err != nil {
		// handle error

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error

	}

	return string(body)

}

func query() fyne.CanvasObject {
	ip := widget.NewEntry()
	ip.SetPlaceHolder("Please input IP address")

	form := &widget.Form{
		OnSubmit: func() {
			info(GetIpInfo(ip.Text))

		},
	}

	form.Append("IP", ip)
	query := widget.NewGroup("Query", form)
	return widget.NewScrollContainer(query)

}

func info(response string) fyne.CanvasObject {
	var i IpInfo
	json.Unmarshal([]byte(response), &i)

	screen := widget.NewForm(
		&widget.FormItem{Text: "IPAddr:", Widget: ip},
		&widget.FormItem{Text: "Position:", Widget: position},
		&widget.FormItem{Text: "ISP测试:", Widget: isp},
	)

	ip.SetText(i.IP)
	position.SetText(i.Position)
	isp.SetText(i.Isp)

	info := widget.NewGroup("Info", screen)
	return widget.NewScrollContainer(info)

}
