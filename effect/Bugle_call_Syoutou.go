package effect

// TITLE:消灯ラッパ(帝国陸軍)

func Bugle_call_Syoutou() *Song {
	return &Song{
		Title:       "消灯ラッパ(帝国陸軍)",	// 楽曲名
		BPM:         108.0,							// 楽曲のテンポ
		Repetitions: 1,								// 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。
		Notes: []Note{
			{ R   , L2  },		// 何故か、最初にノイズが再生されるので、ここに、ダミーの休符を置く。
			{ A5  , L8  },
			{ A5  , L8  },
			{ A5  , L8  },
			{ A5  , L8  },
			{ A5  , L8  },
			{ CS6 , L8  },
			{ A5  , L8  },
			{ CS6 , L8  },
			{ E6  , L1  },
			{ E6  , L16 },
			{ CS6 , L8  },

			{ E6  , L8  },
			{ CS6 , L8  },
			{ A5  , L8  },
			{ CS6 , L8  },
			{ A5  , L8  },
			{ E5  , L8  },
			{ A5  , L1d },
		},
	}
}

