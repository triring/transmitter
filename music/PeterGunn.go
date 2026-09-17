package music

// 楽譜データ
// TITLE:Peter Gunn (ピーター・ガン)
// 1958年から1961年までアメリカで放送された私立探偵テレビドラマのテーマ曲
// 作曲は、ヘンリー・マンシーニ(Henry Mancini)

func PeterGunn() *Song {
	return &Song{
		Title:       "Peter Gunn",	// 楽曲名
		BPM:         111.0,					// 楽曲のテンポ
		Repetitions: 16,						// 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。
		Notes: []Note{
			{R , L64}, // 何故か、最初にノイズが再生されるので、ここに、ダミーの休符を置く。
			{E3, L8}, // 0 E
			{E3, L8}, // 0 E
			{FS3, L8}, // 2 F#
			{E3, L8}, // 0 E
		
			{G3, L16}, // 3 G
			{GS3, L16}, // 4 G#
			{E3, L8}, // 0 E

			{A3, L8}, // 5 A
			{GS3, L8}, // 4 G#
		},
	}
}
