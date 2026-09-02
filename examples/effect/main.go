// tinygo build -target=pico -size=short -o effect.uf2 .
// tinygo flash -target=pico -size=short -monitor .

package main

import (
	"fmt"
	"machine"
	"time"
	// "transmitter" // ローカルのディレクトリに置かれたtransmitterのパッケージをインポートする場合
	"github.com/triring/transmitter" // githubで公開しているパッケージをインポートする場合
	"github.com/triring/transmitter/effect"
)

// raspberry pi pico 用のgpioピンとPWMの対応表
var pinToPWM = map[machine.Pin]transmitter.PWM{
	machine.GPIO0:  machine.PWM0,
	machine.GPIO1:  machine.PWM0,
	machine.GPIO2:  machine.PWM1,
	machine.GPIO3:  machine.PWM1,
	machine.GPIO4:  machine.PWM2,
	machine.GPIO5:  machine.PWM2,
	machine.GPIO6:  machine.PWM3,
	machine.GPIO7:  machine.PWM3,
	machine.GPIO8:  machine.PWM4,
	machine.GPIO9:  machine.PWM4,
	machine.GPIO10: machine.PWM5,
	machine.GPIO11: machine.PWM5,
	machine.GPIO12: machine.PWM6,
	machine.GPIO13: machine.PWM6,
	machine.GPIO14: machine.PWM7,
	machine.GPIO15: machine.PWM7,
	machine.GPIO16: machine.PWM0,
	machine.GPIO17: machine.PWM0,
	machine.GPIO18: machine.PWM1,
	machine.GPIO19: machine.PWM1,
	machine.GPIO20: machine.PWM2,
	machine.GPIO21: machine.PWM2,
	machine.GPIO22: machine.PWM3,
	machine.GPIO23: machine.PWM3,
	machine.GPIO24: machine.PWM4,
	machine.GPIO25: machine.PWM4,
	machine.GPIO26: machine.PWM5,
	machine.GPIO27: machine.PWM5,
	machine.GPIO28: machine.PWM6,
	machine.GPIO29: machine.PWM6,
}

// 送信機を初期化するメソッド
// 引数	: アンテナ線に接続するGPIOのピン名
func initTransmitter(pwm transmitter.PWM, antPin machine.Pin) (transmitter.Transmitter, error) {
	wave, err := transmitter.New(pwm, antPin)
	if err != nil {
		return wave, err
	}
	return wave, nil
}

func main() {
	time.Sleep(500 * time.Millisecond)
	// 出力アンテナに設定するGPIOピンに対応するpwmのチャンネルを取得する。
	pwm := pinToPWM[machine.GPIO15]
	// 微弱電波を出力するGPIOピンを指定して、送信機を設定する。
	tx, err := initTransmitter(pwm, machine.GPIO15)
	if err != nil {
		fmt.Printf("failed to configure PWM\r\n")
		return
	}
	// 搬送波として使用する周波数の設定
	tx.SetFrequency(999 * machine.KHz)

	fmt.Printf("Transmitter Test\n")	
	var CarrierWave uint32 = uint32(pwm.Top() / 2)       // 搬送波の設定
	var IntermittentWave uint32 = uint32(pwm.Top() / 10) // 断続波の設定

	// 効果音等のデータを取り込み
	sound := effect.Westminster_chimes()			// Westminster Chimes
	// sound := effect.Ultraman()					// ウルトラマンのカラータイマー
	// sound := effect.SaoDakeYa()					// 竿竹屋さん
	// sound := effect.Michi()						// 未知との遭遇
	// sound := effect.KingyoUri()					// 金魚売り
	// sound := effect.Jihou()						// 時報
	// sound := effect.Charamela()					// らーめん屋さんのチャルメラ
	// sound := effect.Fanfare()					// 京都&阪神G1ファンファーレ
	// sound := effect.FuneralMarch()				// 葬送行進曲
	// sound := effect.Famima()						// Family Mart 入店音
	// sound := effect.Bugle_call_Souinokosi()		// 起床ラッパ(海上自衛隊)
	// sound := effect.Bugle_call_Totsugeki()		// 突撃ラッパ
	// sound := effect.Bugle_call_Syoutou()			// 消灯ラッパ(帝国陸軍)
	// sound := effect.Bugle_call_Syokuji()			// 食事ラッパ(正露丸のテーマ曲)


	var i int = 0
	var count int = 1
	for {
		fmt.Println(sound.Title) // 楽曲のタイトルを表示する。
		// 別ファイルで定義している楽譜データを読み込み、1音づつ読み出し、再生していく。
		for _, n := range sound.Notes {
			fmt.Printf("%3d : %14.8f,%12.8f\n", count, n.Tone, n.Duration)
			count++
			// 音の高さと長さを算出する。
			tone := (1.0 / (2.0 * n.Tone)) * 1000000.0
			// 楽譜データで定義されているBPMから、tempoを算出する。
			tempo := ((60 / sound.BPM) * (n.Duration * 1000))
			//	Low  := pwm.Set(ch, pwm.Top()>>4) // 音にノイズが交じる。ゼロに設定した時と同じ
			// 休符の場合は、音の再生は行わず、指定時間スリープする。
			if n.Tone == effect.R {
				tx.SetDutyRatio(IntermittentWave)
				time.Sleep(time.Duration(tempo) * time.Millisecond)
				continue
			}
			DUR := time.Duration(tone) * time.Microsecond
			for i := 0.0; i < tempo*1000; i += tone * 2.0 {
				tx.SetDutyRatio(CarrierWave)
				time.Sleep(DUR)
				tx.SetDutyRatio(IntermittentWave)
				time.Sleep(DUR)
			}
		}
		// 楽譜データで停止されているRepetitionsが0の場合は、無限ループ
		if 0 == sound.Repetitions {
			continue
		}
		// 楽譜データで停止されているRepetitionsに設定された回数だけ楽曲を再生する。
		i++
		if i >= sound.Repetitions {
			break
		}
	}
	//	pwm.Set(ch, 0) 完全に出力を止めると、ラジオから未受信の状態でノイズが発生する。
	tx.SetDutyRatio(CarrierWave) // これを防止するために、搬送波(変調していない電波)だけを出し続ける。
	for {
		time.Sleep(1 * time.Hour)
	}
}