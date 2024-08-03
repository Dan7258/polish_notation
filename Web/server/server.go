package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"polish_notation/Opportunities/calculate"
	"polish_notation/Opportunities/convert"
)


func RunServer() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}


func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	t, err := template.ParseFiles("Web/client/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := r.FormValue("text-input")
	send_data := ""
	if data != "" {
		formula, err := convert.SplitOnTokens(data, "1")
		if err != nil {
			send_data = uncorrectData
			log.Println(err.Error())
		} else {
			polish_formula := convert.ConvertToPolish(formula)
			send_data = fmt.Sprintf("%f", calculate.Calculate(polish_formula))
		}

	} 
	err = t.Execute(w, &page{Title: "Just Page", Ans: send_data})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}