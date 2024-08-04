package server

import (
	"fmt"
	"html/template"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"os"
	"polish_notation/Opportunities/calculate"
	"polish_notation/Opportunities/convert"
)

func RunServer() {
	http.HandleFunc("/", index)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("."))))
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./Web/client/"))))
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
	link := ""
	if data != "" {
		err := drawImage(data)
		if err != nil {
			send_data = uncorrectData
			log.Println(err.Error())
		} else {
			send_data = correctData
			link = "/images/someimage.png"
		}
	}

	err = t.Execute(w, &page{Title: "Just Page", Ans: send_data, Src: link})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func drawImage(data string) error {
	var teal color.Color = color.RGBA{0, 200, 200, 255}
	var red color.Color = color.RGBA{200, 30, 30, 255}
	file, err := os.Create("someimage.png")

	if err != nil {
		log.Printf("%v", err)
	}

	img := image.NewRGBA(image.Rect(0, -100, 360, 100))
	draw.Draw(img, img.Bounds(), &image.Uniform{teal}, image.ZP, draw.Src)

	for x := 0.0; x < 360.0; x += 1.0 {
		formula, err := convert.SplitOnTokens(data, fmt.Sprintf("%.8f", x))
		if err != nil {
			return err
		}

		polish_formula := convert.ConvertToPolish(formula)
		y := calculate.Calculate(polish_formula)

		img.Set(int(x), int(y)*-1, red)
		img.Set(int(x-1), int(y)*-1, red)
		img.Set(int(x+1), int(y)*-1, red)
	}
	png.Encode(file, img)
	return nil
}
