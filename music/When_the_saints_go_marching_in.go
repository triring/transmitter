package music

// 楽譜データ
// TITLE:聖者の行進（せいじゃのこうしん、When The Saints Go Marching In）
// 「聖者の行進」は、黒人霊歌の一つ。ディキシーランド・ジャズのナンバーでもある。

func When_the_saints_go_marching_in() *Song {
	return &Song{
		Title:       "聖者の行進",	// 楽曲名
		BPM:         100.0,							// 楽曲のテンポ
		Repetitions: 1,								// 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。
		Notes: []Note{
			{ R  , L8  }, // 何故か、最初にノイズが再生されるので、ここに、ダミーの休符を置く。
			{ C5 , L8  },
			{ E5 , L8  },
			{ F5 , L8  },

			{ G5 , L2  },

			{ R  , L8  },
			{ C5 , L8  },
			{ E5 , L8  },
			{ F5 , L8  },

			{ G5 , L2  },

			{ R  , L8  },
			{ C5 , L8  },
			{ E5 , L8  },
			{ F5 , L8  },

			{ G5 , L4  },
			{ E5 , L4  },

			{ C5 , L4  },
			{ E5 , L4  },

			{ D5 , L2  },

			{ R  , L8  },
			{ E5 , L8  },
			{ E5 , L8  },
			{ D5 , L8  },

			{ C5 , L4d },
			{ C5 , L8  },

			{ E5 , L4  },
			{ G5 , L4  },

			{ G5 , L8  },
			{ F5 , L8  },
			{ F5 , L4  },

			{ F5 , L4  },
			{ E5 , L8  },
			{ F5 , L8  },

			{ G5 , L4  },
			{ E5 , L4  },

			{ C5 , L4  },
			{ D5 , L4  },

			{ C5 , L1  },
		},
	}
}
