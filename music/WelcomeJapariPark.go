package music

// 楽譜データ
// TITLE:ようこそジャパリパークへ

func WelcomeJapariPark() *Song {
	return &Song{
		Title:       "Welcome to Jafari Park",	// 楽曲名
		BPM:         170.0,							// 楽曲のテンポ
		Repetitions: 3,								// 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。
		Notes: []Note{
			{R,  L1}, // 何故か、最初にノイズが再生されるので、ここに、ダミーの休符を置く。
			{G4, L8},
			{F4, L8},
			{G4, L8},
			{A4, L4},
			{G4, L8},
			{A4, L8},
			{B4, L4},

			{C5, L8},
			{CS5, L8},
			{D5, L4},
			{B4, L8},
			{A4, L8},
			{G4, L8},

			{G4, L4},
			{E5, L4},
			{D5, L4},
			{G4, L4},

			{E4, L4},
			{C5, L4},
			{B4, L4},
			{A4, L4},

			{G4, L4},
			{R, L4},
			{R, L2},
		},
	}
}

