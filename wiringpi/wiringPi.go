/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 2.0.8
 * 
 * This file is not intended to be easily readable and contains a number of 
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG 
 * interface file instead. 
 * ----------------------------------------------------------------------------- */

package wiringPi



import "syscall"
import "unsafe"

type _ syscall.Sockaddr

type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

//extern wiringPiSwigCgocall
func SwigCgocall()
//extern wiringPiSwigCgocallDone
func SwigCgocallDone()
//extern wiringPiSwigCgocallBack
func SwigCgocallBack()
//extern wiringPiSwigCgocallBackDone
func SwigCgocallBackDone()

//extern go__wrap_wiringPiSetup
func _swig_wrap_wiringPiSetup() int

func WiringPiSetup() int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiSetup()
}

//extern go__wrap_wiringPiSetupSys
func _swig_wrap_wiringPiSetupSys() int

func WiringPiSetupSys() int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiSetupSys()
}

//extern go__wrap_wiringPiSetupGpio
func _swig_wrap_wiringPiSetupGpio() int

func WiringPiSetupGpio() int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiSetupGpio()
}

//extern go__wrap_wiringPiSetupPhys
func _swig_wrap_wiringPiSetupPhys() int

func WiringPiSetupPhys() int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiSetupPhys()
}

//extern go__wrap_piFaceSetup
func _swig_wrap_piFaceSetup(int) int

func PiFaceSetup(arg1 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_piFaceSetup(arg1)
}

//extern go__wrap_pinMode
func _swig_wrap_pinMode(int, int)

func PinMode(arg1 int, arg2 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_pinMode(arg1, arg2)
}

//extern go__wrap_pullUpDnControl
func _swig_wrap_pullUpDnControl(int, int)

func PullUpDnControl(arg1 int, arg2 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_pullUpDnControl(arg1, arg2)
}

//extern go__wrap_digitalRead
func _swig_wrap_digitalRead(int) int

func DigitalRead(arg1 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_digitalRead(arg1)
}

//extern go__wrap_digitalWrite
func _swig_wrap_digitalWrite(int, int)

func DigitalWrite(arg1 int, arg2 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_digitalWrite(arg1, arg2)
}

//extern go__wrap_pwmWrite
func _swig_wrap_pwmWrite(int, int)

func PwmWrite(arg1 int, arg2 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_pwmWrite(arg1, arg2)
}

//extern go__wrap_analogRead
func _swig_wrap_analogRead(int) int

func AnalogRead(arg1 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_analogRead(arg1)
}

//extern go__wrap_analogWrite
func _swig_wrap_analogWrite(int, int)

func AnalogWrite(arg1 int, arg2 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_analogWrite(arg1, arg2)
}

//extern go__wrap_piBoardRev
func _swig_wrap_piBoardRev() int

func PiBoardRev() int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_piBoardRev()
}

//extern go__wrap_wpiPinToGpio
func _swig_wrap_wpiPinToGpio(int) int

func WpiPinToGpio(arg1 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wpiPinToGpio(arg1)
}

//extern go__wrap_setPadDrive
func _swig_wrap_setPadDrive(int, int)

func SetPadDrive(arg1 int, arg2 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_setPadDrive(arg1, arg2)
}

//extern go__wrap_getAlt
func _swig_wrap_getAlt(int) int

func GetAlt(arg1 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_getAlt(arg1)
}

//extern go__wrap_digitalWriteByte
func _swig_wrap_digitalWriteByte(int)

func DigitalWriteByte(arg1 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_digitalWriteByte(arg1)
}

//extern go__wrap_pwmSetMode
func _swig_wrap_pwmSetMode(int)

func PwmSetMode(arg1 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_pwmSetMode(arg1)
}

//extern go__wrap_pwmSetRange
func _swig_wrap_pwmSetRange(uint)

func PwmSetRange(arg1 uint) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_pwmSetRange(arg1)
}

//extern go__wrap_pwmSetClock
func _swig_wrap_pwmSetClock(int)

func PwmSetClock(arg1 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_pwmSetClock(arg1)
}

//extern go__wrap_gpioClockSet
func _swig_wrap_gpioClockSet(int, int)

func GpioClockSet(arg1 int, arg2 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_gpioClockSet(arg1, arg2)
}

//extern go__wrap_waitForInterrupt_set
func _swig_wrap_waitForInterrupt_set(_swig_fnptr)

func SetWaitForInterrupt(arg1 _swig_fnptr) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_waitForInterrupt_set(arg1)
}

//extern go__wrap_waitForInterrupt_get
func _swig_wrap_waitForInterrupt_get() _swig_fnptr

func GetWaitForInterrupt() _swig_fnptr {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_waitForInterrupt_get()
}

//extern go__wrap_wiringPiISR
func _swig_wrap_wiringPiISR(int, int, _swig_fnptr) int

func WiringPiISR(arg1 int, arg2 int, arg3 _swig_fnptr) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiISR(arg1, arg2, arg3)
}

//extern go__wrap_piThreadCreate
func _swig_wrap_piThreadCreate(_swig_fnptr) int

func PiThreadCreate(arg1 _swig_fnptr) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_piThreadCreate(arg1)
}

//extern go__wrap_piLock
func _swig_wrap_piLock(int)

func PiLock(arg1 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_piLock(arg1)
}

//extern go__wrap_piUnlock
func _swig_wrap_piUnlock(int)

func PiUnlock(arg1 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_piUnlock(arg1)
}

//extern go__wrap_piHiPri
func _swig_wrap_piHiPri(int) int

func PiHiPri(arg1 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_piHiPri(arg1)
}

//extern go__wrap_delay
func _swig_wrap_delay(uint)

func Delay(arg1 uint) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_delay(arg1)
}

//extern go__wrap_delayMicroseconds
func _swig_wrap_delayMicroseconds(uint)

func DelayMicroseconds(arg1 uint) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_delayMicroseconds(arg1)
}

//extern go__wrap_millis
func _swig_wrap_millis() uint

func Millis() uint {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_millis()
}

//extern go__wrap_micros
func _swig_wrap_micros() uint

func Micros() uint {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_micros()
}

//extern go__wrap_serialOpen
func _swig_wrap_serialOpen(string, int) int

func SerialOpen(arg1 string, arg2 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_serialOpen(arg1, arg2)
}

//extern go__wrap_serialClose
func _swig_wrap_serialClose(int)

func SerialClose(arg1 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_serialClose(arg1)
}

//extern go__wrap_serialFlush
func _swig_wrap_serialFlush(int)

func SerialFlush(arg1 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_serialFlush(arg1)
}

//extern go__wrap_serialPutchar
func _swig_wrap_serialPutchar(int, byte)

func SerialPutchar(arg1 int, arg2 byte) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_serialPutchar(arg1, arg2)
}

//extern go__wrap_serialPuts
func _swig_wrap_serialPuts(int, string)

func SerialPuts(arg1 int, arg2 string) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_serialPuts(arg1, arg2)
}

//extern go__wrap_serialPrintf
func _swig_wrap_serialPrintf(int, string)

func SerialPrintf(arg1 int, arg2 string) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_serialPrintf(arg1, arg2)
}

//extern go__wrap_serialDataAvail
func _swig_wrap_serialDataAvail(int) int

func SerialDataAvail(arg1 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_serialDataAvail(arg1)
}

//extern go__wrap_serialGetchar
func _swig_wrap_serialGetchar(int) int

func SerialGetchar(arg1 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_serialGetchar(arg1)
}

//extern go__wrap_shiftOut
func _swig_wrap_shiftOut(byte, byte, byte, byte)

func ShiftOut(arg1 byte, arg2 byte, arg3 byte, arg4 byte) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_shiftOut(arg1, arg2, arg3, arg4)
}

//extern go__wrap_shiftIn
func _swig_wrap_shiftIn(byte, byte, byte) byte

func ShiftIn(arg1 byte, arg2 byte, arg3 byte) byte {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_shiftIn(arg1, arg2, arg3)
}

//extern go__wrap_wiringPiSPIGetFd
func _swig_wrap_wiringPiSPIGetFd(int) int

func WiringPiSPIGetFd(arg1 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiSPIGetFd(arg1)
}

//extern go__wrap_wiringPiSPIDataRW
func _swig_wrap_wiringPiSPIDataRW(int, string) int

func WiringPiSPIDataRW(arg1 int, arg2 string) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiSPIDataRW(arg1, arg2)
}

//extern go__wrap_wiringPiSPISetup
func _swig_wrap_wiringPiSPISetup(int, int) int

func WiringPiSPISetup(arg1 int, arg2 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiSPISetup(arg1, arg2)
}

//extern go__wrap_wiringPiI2CSetupInterface
func _swig_wrap_wiringPiI2CSetupInterface(string, int) int

func WiringPiI2CSetupInterface(arg1 string, arg2 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiI2CSetupInterface(arg1, arg2)
}

//extern go__wrap_wiringPiI2CSetup
func _swig_wrap_wiringPiI2CSetup(int) int

func WiringPiI2CSetup(arg1 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiI2CSetup(arg1)
}

//extern go__wrap_wiringPiI2CRead
func _swig_wrap_wiringPiI2CRead(int) int

func WiringPiI2CRead(arg1 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiI2CRead(arg1)
}

//extern go__wrap_wiringPiI2CReadReg8
func _swig_wrap_wiringPiI2CReadReg8(int, int) int

func WiringPiI2CReadReg8(arg1 int, arg2 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiI2CReadReg8(arg1, arg2)
}

//extern go__wrap_wiringPiI2CReadReg16
func _swig_wrap_wiringPiI2CReadReg16(int, int) int

func WiringPiI2CReadReg16(arg1 int, arg2 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiI2CReadReg16(arg1, arg2)
}

//extern go__wrap_wiringPiI2CWrite
func _swig_wrap_wiringPiI2CWrite(int, int) int

func WiringPiI2CWrite(arg1 int, arg2 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiI2CWrite(arg1, arg2)
}

//extern go__wrap_wiringPiI2CWriteReg8
func _swig_wrap_wiringPiI2CWriteReg8(int, int, int) int

func WiringPiI2CWriteReg8(arg1 int, arg2 int, arg3 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiI2CWriteReg8(arg1, arg2, arg3)
}

//extern go__wrap_wiringPiI2CWriteReg16
func _swig_wrap_wiringPiI2CWriteReg16(int, int, int) int

func WiringPiI2CWriteReg16(arg1 int, arg2 int, arg3 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_wiringPiI2CWriteReg16(arg1, arg2, arg3)
}

//extern go__wrap_softToneCreate
func _swig_wrap_softToneCreate(int) int

func SoftToneCreate(arg1 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_softToneCreate(arg1)
}

//extern go__wrap_softToneWrite
func _swig_wrap_softToneWrite(int, int)

func SoftToneWrite(arg1 int, arg2 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_softToneWrite(arg1, arg2)
}

//extern go__wrap_softServoWrite
func _swig_wrap_softServoWrite(int, int)

func SoftServoWrite(arg1 int, arg2 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_softServoWrite(arg1, arg2)
}

//extern go__wrap_softServoSetup
func _swig_wrap_softServoSetup(int, int, int, int, int, int, int, int) int

func SoftServoSetup(arg1 int, arg2 int, arg3 int, arg4 int, arg5 int, arg6 int, arg7 int, arg8 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_softServoSetup(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

//extern go__wrap_softPwmCreate
func _swig_wrap_softPwmCreate(int, int, int) int

func SoftPwmCreate(arg1 int, arg2 int, arg3 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_softPwmCreate(arg1, arg2, arg3)
}

//extern go__wrap_softPwmWrite
func _swig_wrap_softPwmWrite(int, int)

func SoftPwmWrite(arg1 int, arg2 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_softPwmWrite(arg1, arg2)
}

//extern go__wrap_mcp23s17Setup
func _swig_wrap_mcp23s17Setup(int, int, int) int

func Mcp23s17Setup(arg1 int, arg2 int, arg3 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_mcp23s17Setup(arg1, arg2, arg3)
}

//extern go__wrap_mcp23017Setup
func _swig_wrap_mcp23017Setup(int, int) int

func Mcp23017Setup(arg1 int, arg2 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_mcp23017Setup(arg1, arg2)
}

//extern go__wrap_mcp23s08Setup
func _swig_wrap_mcp23s08Setup(int, int, int) int

func Mcp23s08Setup(arg1 int, arg2 int, arg3 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_mcp23s08Setup(arg1, arg2, arg3)
}

//extern go__wrap_mcp23008Setup
func _swig_wrap_mcp23008Setup(int, int) int

func Mcp23008Setup(arg1 int, arg2 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_mcp23008Setup(arg1, arg2)
}

//extern go__wrap_sr595Setup
func _swig_wrap_sr595Setup(int, int, int, int, int) int

func Sr595Setup(arg1 int, arg2 int, arg3 int, arg4 int, arg5 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_sr595Setup(arg1, arg2, arg3, arg4, arg5)
}

//extern go__wrap_lcdHome
func _swig_wrap_lcdHome(int)

func LcdHome(arg1 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_lcdHome(arg1)
}

//extern go__wrap_lcdClear
func _swig_wrap_lcdClear(int)

func LcdClear(arg1 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_lcdClear(arg1)
}

//extern go__wrap_lcdSendCommand
func _swig_wrap_lcdSendCommand(int, byte)

func LcdSendCommand(arg1 int, arg2 byte) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_lcdSendCommand(arg1, arg2)
}

//extern go__wrap_lcdPosition
func _swig_wrap_lcdPosition(int, int, int)

func LcdPosition(arg1 int, arg2 int, arg3 int) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_lcdPosition(arg1, arg2, arg3)
}

//extern go__wrap_lcdPutchar
func _swig_wrap_lcdPutchar(int, byte)

func LcdPutchar(arg1 int, arg2 byte) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_lcdPutchar(arg1, arg2)
}

//extern go__wrap_lcdPuts
func _swig_wrap_lcdPuts(int, string)

func LcdPuts(arg1 int, arg2 string) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_lcdPuts(arg1, arg2)
}

//extern go__wrap_lcdPrintf
func _swig_wrap_lcdPrintf(int, string)

func LcdPrintf(arg1 int, arg2 string) {
	defer SwigCgocallDone()
	SwigCgocall()
	_swig_wrap_lcdPrintf(arg1, arg2)
}

//extern go__wrap_lcdInit
func _swig_wrap_lcdInit(int, int, int, int, int, int, int, int, int, int, int, int, int) int

func LcdInit(arg1 int, arg2 int, arg3 int, arg4 int, arg5 int, arg6 int, arg7 int, arg8 int, arg9 int, arg10 int, arg11 int, arg12 int, arg13 int) int {
	defer SwigCgocallDone()
	SwigCgocall()
	return _swig_wrap_lcdInit(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

