package wg

import (
	rand2 "crypto/rand"
	"fmt"
	"github.com/dop251/goja"
	"github.com/pkg/errors"
)

type WGKeyPair struct {
	PrivateKeyB64 string
	PublicKeyB64  string
}

func GenerateSelfKeyPair() *WGKeyPair {
	var rt = goja.New()
	var windowObject = rt.NewObject()
	if err := rt.Set("window", windowObject); err != nil {
		panic(err)
	}
	var cryptoObject = rt.NewObject()
	if err := cryptoObject.Set("getRandomValues", rt.ToValue(getRandomValuesJSFunc)); err != nil {
		panic(err)
	}
	if err := windowObject.Set("crypto", cryptoObject); err != nil {
		panic(err)
	}
	_, err := rt.RunString(KeyGeneratorJS)
	if err != nil {
		panic(err)
	}
	fmt.Println(windowObject.Keys())
	var generateKeyPairValueJSValue = windowObject.Get("wireguard").
		ToObject(rt).
		Get("generateKeypair")
	generateKeyPairValue, ok := goja.AssertFunction(generateKeyPairValueJSValue)
	if !ok {
		panic("unsupported js generator")
	}
	resultJSValue, err := generateKeyPairValue(windowObject)
	if err != nil {
		panic(errors.Wrap(err, "call generate func"))
	}
	var resultJSObj = resultJSValue.ToObject(rt)
	var keyPair = &WGKeyPair{
		PrivateKeyB64: resultJSObj.Get("privateKey").String(),
		PublicKeyB64:  resultJSObj.Get("publicKey").String(),
	}
	return keyPair
}

func getRandomValuesJSFunc(funcCall goja.FunctionCall, rt *goja.Runtime) goja.Value {
	var resultUInt8Array = funcCall.Argument(0)
	var export = resultUInt8Array.Export().([]byte)
	if _, err := rand2.Read(export); err != nil {
		panic(errors.Wrap(err, "read crypto random"))
	}
	return goja.Undefined()
}
