package music

// 曲データの共通構造体

type Note struct {
	Tone     float64
	Duration float64
}

type Song struct {
    Title       string
    BPM         float64
    Repetitions int
    Notes       []Note
}
