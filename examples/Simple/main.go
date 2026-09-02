// tinygo build -target=pico -size=short -o Simple.uf2 .
// tinygo flash -target=pico -size=short -monitor .

// 中央ラ音 440Hz を再生するデモ

package main

import (
	"fmt"
	"machine"
	"time"
	// "transmitter" // ローカルのディレクトリに置かれたtransmitterのパッケージをインポートする場合
	"github.com/triring/transmitter" // githubで公開しているパッケージをインポートする場合
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

// AM放送の周波数帯(526.5～1606.5KHz)では、9KHz間隔で割り振られている。
// これに適合する誤差の少ない周波数を生成できるPWMの設定は、以下の通り。
// これ以外の周波数を使用する場合は、周波数の誤差は大きくなるが、以下のように出力したい周波数をHzで設定すればよい。
//
//	tx.SetFrequency(1494000)	// 1494MHzの電波を生成する場合
const AM0612KHz uint64 =  612_000
const AM0684KHz uint64 =  684_000
const AM0693KHz uint64 =  693_000
const AM0738KHz uint64 =  738_000
const AM0774KHz uint64 =  774_000
const AM0819KHz uint64 =  819_000
const AM0999KHz uint64 =  999_000
const AM1224KHz uint64 = 1224_000
const AM1269KHz uint64 = 1269_000
const AM1287KHz uint64 = 1287_000
const AM1368KHz uint64 = 1368_000
const AM1548KHz uint64 = 1548_000

// 日本のFM放送の周波数帯は、76.1～94.9MHzである。
// 世界のFM放送の周波数帯は、87.5～108.0 MHz)が一般的である。
// これらの周波数帯の波長をRaspberry pi picoのPWMで生成する不可能である。
// そこで、逓倍波で、これらに適合する周波数となるPWMの設定を探した結果は、以下のとおりである。
const FREQ_2MHz     uint64 =  2_000_000		// 受信できた周波数(MHz) : 76.0, 78.0, 80.0, 82.0, 84.5, 86.0, 90.1, 92.0, 94.0
const FREQ_4MHz     uint64 =  4_000_000		// 受信できた周波数(MHz) : 76.0, 78.5, 84.0, 85.5, 88.0, 92.0
const FREQ_5_20MHz  uint64 =  5_208_333		// 受信できた周波数(MHz) : 84.1, 91.4
const FREQ_8MHz     uint64 =  8_000_000		// 受信できた周波数(MHz) : 88.0, 91.3
const FREQ_10MHz    uint64 = 10_000_000		// 受信できた周波数(MHz) : 80.0, 86.5, 90.0
const FREQ_16MHz    uint64 = 16_000_000		// 受信できた周波数(MHz) : 89.9, 90.1
const FREQ_20MHz    uint64 = 20_000_000		// 受信できた周波数(MHz) : 80.0, 83.9, 86.5, 87.9, 93.5
const FREQ_25MHz    uint64 = 25_000_000		// 受信できた周波数(MHz) : 76.0, 83.5, 89.9

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
//	tx.SetFrequency(AM0999KHz)
	tx.SetFrequency(FREQ_10MHz)

	t := (1000 * time.Millisecond) / 440
//	t := (1000 * time.Millisecond) / 880
//	t := (1000 * time.Millisecond) / 1760
	count := 0
	// Two tone siren.
	for {
		tx.SetDutyRatio(2)
		time.Sleep(t)
		tx.Stop()
		time.Sleep(t)
		count++
		if count > 2200 {
			fmt.Printf("440Hz\r\n")
			count = 0
		}
	}
}