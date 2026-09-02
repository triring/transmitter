package effect

// TITLE:金魚売り

func KingyoUri() *Song {
	return &Song{
		Title:       "金魚売り",	// 楽曲名
		BPM:         120.0,			// 楽曲のテンポ
		Repetitions: 3,				// 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。
		Notes: []Note{
			{ R  , L2  },      // 何故か、最初にノイズが再生されるので、ここに、ダミーの休符を置く。
			{ E5 , L8d },
			{ G5 , L16 },
			{ G5 , L2  },
			{ E5 , L4d },
			{ G5 , L16 },
			{ A5 , L16 },
			{ A5 , L1  },
		},
	}
}

