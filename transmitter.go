package transmitter

import (
	"machine"
)

// PWM is the interface necessary for controlling a transmitter.
type PWM interface {
	Configure(config machine.PWMConfig) error
	Channel(pin machine.Pin) (channel uint8, err error)
	Top() uint32
	Set(channel uint8, value uint32)
	SetPeriod(period uint64) error
}

// Transmitter is a configured audio output channel based on a PWM.
type Transmitter struct {
	pwm PWM
	ch  uint8
}

// New returns a new transmitter instance simply configured for the specified PWM and pin combination.
// Since PWM is used for the waveform output, the waveform forms a square wave.
func New(pwm PWM, pin machine.Pin) (Transmitter, error) {
	err := pwm.Configure(machine.PWMConfig{
		Period: uint64(1e9) / (999 * machine.KHz),
	})
	if err != nil {
		return Transmitter{}, err
	}
	ch, err := pwm.Channel(pin)
	if err != nil {
		return Transmitter{}, err
	}
	pwm.Set(ch, 0)	// いきなり電波を送出しないように、出力を停止状態にしておく。
	return Transmitter{pwm, ch}, nil
}

// Get PWM’s top value.
func (tx Transmitter) GetTop() uint32 {
	return tx.pwm.Top()
}


// Stop disables the transmitter, setting the output to low continuously.
func (tx Transmitter) Stop() {
	tx.pwm.Set(tx.ch, 0)
}

// 
func (tx Transmitter) SetDutyRatio(ratio uint32) {
	tx.pwm.Set(tx.ch, tx.pwm.Top()/ratio)
}

// SetPeriod sets the period for the signal in nanoseconds. Use the following
// formula to convert frequency to period:
//
//	period = 1e9 / frequency
func (tx Transmitter) SetPeriod(period uint64) {
	// Disable output.
	tx.Stop()

	if period == 0 {
		// Assume a period of 0 is intended as "no output".
		return
	}

	// Reconfigure period.
	tx.pwm.SetPeriod(period)

	// Make this a square wave by setting the channel position to half the
	// period.
	tx.pwm.Set(tx.ch, tx.pwm.Top()/2)
}


// SetFrequency generates radio waves at the specified frequency.
// For example, s.SetFrequency(999 * machine.KHz) produces a 999KHz square wave.
func (tx Transmitter) SetFrequency(frequency uint64) {
	tx.SetPeriod(uint64(1e9) / frequency)
}

