package email

import (
	"bytes"
	"html/template"
	"server/utils"
)

func SendCodeTPl(code string) string {
	appDir := utils.GetAppDir()
	tpl := template.Must(template.ParseFiles(appDir + "/email/template/send_code.html"))

	buf := new(bytes.Buffer)
	_ = tpl.ExecuteTemplate(
		buf,
		"sendcode",
		map[string]interface{}{
			"Code": code,
		},
	)
	return buf.String()
}
