package duk_bridge

/**
 * a wrapper of ecmascript function or wrapper
 * Rosbit Xu <me@rosbit.cn>
 * Dec. 4, 2018
 */

/*
#include "duk_bridge.h"
*/
import "C"
import (
	"unsafe"
)

type EcmaObject struct {
	ecmaObj unsafe.Pointer
	isFunc bool
}

func wrapEcmaObject(ecmaObj unsafe.Pointer, isFunc bool) *EcmaObject {
	return &EcmaObject{ecmaObj, isFunc}
}

func (m *EcmaObject) destroy(env unsafe.Pointer) {
	C.js_destroy_ecmascript_obj(env, m.ecmaObj);
}
