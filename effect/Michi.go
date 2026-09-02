package effect

// TITLE:未知との遭遇
// Close Encounters of the Third Kind
// 〔米1977《監督》スティーブン・スピルバーグ
//         《出演》リチャード・ドレイファス、テリー・ガー

func Michi() *Song {
	return &Song{
		Title:       "未知との遭遇",	// 楽曲名
		BPM:         120.0,			// 楽曲のテンポ
		Repetitions: 5,				// 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。
		Notes: []Note{
			{ R  , L2 },      // 何故か、最初にノイズが再生されるので、ここに、ダミーの休符を置く。
			{ G5 , L4 },	// ５ソ  
			{ A5 , L4 },	// ５ラ  
			{ F5 , L4 },	// ５ファ
			{ F4 , L4 },	// ４ファ
			{ C5 , L2 },	// ５Ｃ  
			{ R  , L1 },	// ５Ｃ  
		},
	}
}

