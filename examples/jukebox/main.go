// tinygo build -target=pico -size=short -o jukebox.uf2 .
// tinygo flash -target=pico -size=short -monitor .

// 複数の楽曲を順次再生していくデモ

package main

import (
	"fmt"
	"machine"
	"time"
	// "transmitter" // ローカルのディレクトリに置かれたtransmitterのパッケージをインポートする場合
	"github.com/triring/transmitter" // githubで公開しているパッケージをインポートする場合
	"github.com/triring/transmitter/music"
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
// 楽曲データのリスト
song_lists := []*music.Song{
	music.IevanPolkka(),
	music.CheCheKoolay(),
	music.JingleBells(),
	music.KimiGaYo(),
	music.Le_temps_des_cerises(),
	music.Ode_to_Joy(),
	music.Hell_and_heaven(),
	music.TheBearSong(),
	music.Goin_Home(),
	music.SeeTheConqueringHeroComes(),
	music.When_the_saints_go_marching_in(),
	music.YuyakeKoyake(),
	music.Etenraku(),

//	music.MitoKoumon(),
//	music.KonoKiNannoKi(),
//	music.ShortShorts(),
//	music.WelcomeJapariPark(),
//	music.BonMawashi(),
//	music.DarthVader(),
//	music.Godzilla(),
//	music.KnightRider(),
}

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
	var count int = 1

	for {
		for _, song_data := range song_lists {
			fmt.Printf("%s\n", song_data.Title)	// 楽曲のタイトルを表示する。
			// 別ファイルで定義している楽譜データを読み込み、1音づつ読み出し、再生していく。
			count = 1
			for _, n := range song_data.Notes {
				fmt.Printf("%3d : %14.8f,%12.8f\n", count, n.Tone, n.Duration)
				count++
				// 音の高さと長さを算出する。
				tone := (1.0 / (2.0 * n.Tone)) * 1000000.0
				// 楽譜データで定義されているBPMから、tempoを算出する。
				tempo := ((60 / song_data.BPM) * (n.Duration * 1000))
				//	Low  := pwm.Set(ch, pwm.Top()>>4) // 音にノイズが交じる。ゼロに設定した時と同じ
				// 休符の場合は、音の再生は行わず、指定時間スリープする。
				if n.Tone == music.R {
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
		}	
	}
	//	pwm.Set(ch, 0) 完全に出力を止めると、ラジオから未受信の状態でノイズが発生する。
	tx.SetDutyRatio(CarrierWave) // これを防止するために、搬送波(変調していない電波)だけを出し続ける。
	for {
		time.Sleep(1 * time.Hour)
	}
}