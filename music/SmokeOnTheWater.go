package music

// 楽譜データ
// TITLE:Smoke on the water (スモーク・オン・ザ・ウォーター)
// イギリスのロックバンド Deep Purple (ディープ・パープル) の代表曲

func SmokeOnTheWater() *Song {
	return &Song{
		Title:       "Smoke on the water",	// 楽曲名
		BPM:         112.0,					// 楽曲のテンポ
		Repetitions: 3,						// 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。
		Notes: []Note{
			{R , L2}, // 何故か、最初にノイズが再生されるので、ここに、ダミーの休符を置く。
			{B6, L8}, // 0 B
			{R,  L8},
			{D6, L8}, // 3 D
			{R,  L8},
			{E6, L4}, // 5 E
			{R,  L8},
			{B6, L8}, // 0 B
			
			{R,  L8},
			{D6, L8}, // 3 D
			{R,  L8},
			{F6, L8}, // 6 F
			{E6, L4}, // 5 E
			{R,  L4},

			{B6, L8}, // 0 B
			{R,  L8},
			{D6, L8}, // 3 D
			{R,  L8},
			{E6, L4}, // 5 E
			{R,  L8},
			{D6, L8}, // 3 D

			{R,  L8},
			{B6, L8}, // 0 B
			{B6, L4}, // 0 B
			{B6, L4}, // 0 B
			{R,  L4}, // 0 B
		},
		/*
		Notes: []Note{
			{R , L2}, // 何故か、最初にノイズが再生されるので、ここに、ダミーの休符を置く。
			{B5, L8}, // 0 B
			{R,  L8},
			{D5, L8}, // 3 D
			{R,  L8},
			{E5, L4}, // 5 E
			{R,  L8},
			{B5, L8}, // 0 B
			
			{R,  L8},
			{D5, L8}, // 3 D
			{R,  L8},
			{F5, L8}, // 6 F
			{E5, L4}, // 5 E
			{R,  L4},

			{B5, L8}, // 0 B
			{R,  L8},
			{D5, L8}, // 3 D
			{R,  L8},
			{E5, L4}, // 5 E
			{R,  L8},
			{D5, L8}, // 3 D

			{R,  L8},
			{B5, L8}, // 0 B
			{B5, L4}, // 0 B
			{B5, L4}, // 0 B
			{R,  L4}, // 0 B
		},
	*/
	}
}
