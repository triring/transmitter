package effect

// TITLE:起床ラッパ(海上自衛隊)

func Bugle_call_Souinokosi() *Song {
	return &Song{
		Title:       "起床ラッパ(海上自衛隊)",	// 楽曲名
		BPM:         240.0,							// 楽曲のテンポ
		Repetitions: 1,								// 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。
		Notes: []Note{
			{ R   , L2  },		// 何故か、最初にノイズが再生されるので、ここに、ダミーの休符を置く。
			{ G4  , L4  },
			{ G4  , L8  },
			{ D4  , L4  },
			{ G4  , L8  },
			{ B4  , L4  },
			{ B4  , L8  },
			{ G4  , L4  },
			{ B4  , L8  },
			{ D5  , L4  },
			{ D5  , L8  },
			{ B4  , L4  },
			{ G4  , L8  },
			{ D4  , L4  },
			{ D4  , L8  },
			{ D4  , L4d },
			{ G4  , L4  },
			{ G4  , L8  },
			{ D4  , L4  },
			{ G4  , L8  },
			{ B4  , L4  },
			{ B4  , L8  },
			{ G4  , L4  },
			{ B4  , L8  },
			{ D4  , L4  },
			{ D4  , L8  },
			{ G4  , L4  },
			{ B4  , L8  },
			{ D5  , L4  },
			{ D5  , L8  },
			{ D5  , L1  },
			{ R   , L32 },
			{ D5  , L16d},
		},
	}
}

