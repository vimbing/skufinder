package api

import "github.com/dop251/goja"

func getAdditionalCheckFuncFromJs(script string) func(string) bool {
	return func(body string) bool {
		vm := goja.New()

		_, err := vm.RunString(script)

		if err != nil {
			return false
		}

		sum, ok := goja.AssertFunction(vm.Get("check"))

		if !ok {
			return false
		}

		res, err := sum(goja.Undefined(), vm.ToValue(body))

		if err != nil {
			return false
		}

		return res.ToBoolean()
	}
}
