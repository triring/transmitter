package effect

// TITLE:Westminster Chimes

func Westminster_chimes() *Song {
	return &Song{
		Title:       "Westminster Chimes",	// 楽曲名
		BPM:         80.0,					// 楽曲のテンポ
		Repetitions: 1,						// 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。
		Notes: []Note{
			{R , L2  },      // 何故か、最初にノイズが再生されるので、ここに、ダミーの休符を置く。
			{ C5 , L4  },
			{ E5 , L4  },
			{ D5 , L4  },
			{ G4 , L2d },
			{ C5 , L4  },
			{ D5 , L4  },
			{ E5 , L4  },
			{ C5 , L2d },
			{ E5 , L4  },
			{ C5 , L4  },
			{ D5 , L4  },
			{ G4 , L2d },
			{ G4 , L4  },
			{ D5 , L4  },
			{ E5 , L4  },
			{ C5 , L2d },
		},
	}
}

