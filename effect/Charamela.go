package effect

// TITLE:らーめん屋さんのチャルメラ

func Charamela() *Song {
	return &Song{
		Title:       "らーめん屋さんのチャルメラ",	// 楽曲名
		BPM:         120.0,							// 楽曲のテンポ
		Repetitions: 3,								// 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。
		Notes: []Note{
			{ R   , L2  },	 // 何故か、最初にノイズが再生されるので、ここに、ダミーの休符を置く。
			{ C5  , L8d },
			{ D5  , L16 },
			{ E5  , L2d },
			{ D5  , L16 },
			{ C5  , L8  },
			{ R   , L8  },
			{ C5  , L8  },
			{ D5  , L8  },
			{ E5  , L8  },
			{ D5  , L8  },
			{ C5  , L8  },
			{ D5  , L2d },
		},
	}
}

