package music

// 楽譜データ
// TITLE:Another One Bites the Dust (地獄へ道づれ)
// イギリスのロックバンド Queen (クイーン) の代表曲

func AnotherOneBitesTheDust() *Song {
	return &Song{
		Title:       "Another One Bites the Dust",	// 楽曲名
		BPM:         111.0,					// 楽曲のテンポ
		Repetitions: 3,						// 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。
		Notes: []Note{
			{R , L32}, // 何故か、最初にノイズが再生されるので、ここに、ダミーの休符を置く。
			{E4, L8}, // 0 E
			{R,  L8},
			{E4, L8}, // 0 E
			{R,  L8},
			{E4, L8}, // 0 E
			{R,  L8},
			{R,  L8},
			{E4, L16}, // 0 E

			{E4, L8}, // 0 E
			{E4, L8}, // 0 E
			{G4, L8}, // 3 G
			{E4, L16}, // 0 E
			{A4, L16}, // 5 A
			{R,  L2},
		},
	}
}
