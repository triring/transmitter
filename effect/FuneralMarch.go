package effect

// 楽譜データ
// TITLE:葬送行進曲

func FuneralMarch() *Song {
	return &Song{
		Title:       "Funeral March",	// 楽曲名
		BPM:         100.0,							// 楽曲のテンポ
		Repetitions: 1,								// 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。
		Notes: []Note{
			{ R   ,   L2  },	// 何故か、最初にノイズが再生されるので、ここに、ダミーの休符を置く。
			{ AS3 ,   L4  },
			{ R   ,   L32 },
			{ AS3 ,   L8d },
			{ R   ,   L32 },
			{ AS3 ,   L16 },
			{ R   ,   L32 },
			{ AS3 ,   L4  },
			{ R   ,   L32 },
			{ CS2 ,   L8d },
			{ R   ,   L32 },
			{ C2  ,   L16 },
			{ R   ,   L32 },
			{ CS2 ,   L32 },
			{ R   ,   L32 },
			{ C2  ,   L8  },
		// 	{ R   ,   L32 },
			{ C2  ,   L32 },
			{ R   ,   L32 },
			{ AS3 ,   L16 },
			{ R   ,   L32 },
			{ AS3 ,   L8d },
			{ R   ,   L32 },
			{ AS3 ,   L16 },
			{ R   ,   L32 },
			{ AS3 ,   L2  },
			{ R   ,   L1  },
			{ R   ,   L1  },
		},
	}
}

