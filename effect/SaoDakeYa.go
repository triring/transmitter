package effect

// TITLE:竿竹屋さん

func SaoDakeYa() *Song {
	return &Song{
		Title:       "竿竹屋さん",	// 楽曲名
		BPM:         120.0,			// 楽曲のテンポ
		Repetitions: 3,				// 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。
		Notes: []Note{
			{ R  , L2  },      // 何故か、最初にノイズが再生されるので、ここに、ダミーの休符を置く。
			{ E5 , L8d },
			{ G5 , L16 },
			{ G5 , L2d },
			{ G5 , L8  },
			{ A5 , L8  },
			{ A5 , L8  },
			{ A5 , L1  },
		},
	}
}

