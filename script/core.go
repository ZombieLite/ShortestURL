package script

import (
	"html/template"
	"net/http"
)

type Data struct {
	Host               string
	Link               string
	IsNotAvailableLink bool
	NewLink            string
}

func CoreLink(w http.ResponseWriter, r *http.Request) {
	strlink := r.FormValue("text")

	data := Data{
		Host:               r.Host,
		Link:               strlink,
		IsNotAvailableLink: IsEmptyString(strlink),
		NewLink:            "",
	}

	if !data.IsNotAvailableLink {
		data.ReplaceLink()
	}

	tmpl, err := template.ParseFiles("templates/changelink.html")

	if err != nil {
		err.Error()
		return
	}

	tmpl.Execute(w, data)
}
