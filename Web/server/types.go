package server

type page struct {
	Title string
	Ans   string
	Src   string
}

const (
	uncorrectData = "Формат введенных данных не соответствует необходимому! Пожалуйста, вводите формулу без пробелов и используйте только разрешенные опрерации и функции!"
	correctData   = "Ваш график: "
)
