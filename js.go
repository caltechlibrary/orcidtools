//
// Package ot wraps the data structures and services returned by the Orcid API.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2016, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package ot

import
// 3rd Party packages

// Caltech Library packages
(
	"fmt"
	"log"

	"github.com/caltechlibrary/ostdlib"
	"github.com/robertkrimen/otto"
)

//
// JavaScript integration via otto
//

// AddExtensions adds the ot API wrapper to a JavaScriptVM
func (api *OrcidAPI) AddExtensions(js *ostdlib.JavaScriptVM) *otto.Otto {
	vm := js.VM
	errorObject := func(obj *otto.Object, msg string) otto.Value {
		if obj == nil {
			obj, _ = vm.Object(`({})`)
		}
		log.Println(msg)
		obj.Set("status", "error")
		obj.Set("error", msg)
		return obj.Value()
	}

	// responseObject := func(data interface{}) otto.Value {
	//      src, _ := json.Marshal(data)
	//      obj, _ := vm.Object(fmt.Sprintf(`(%s)`, src))
	//      return obj.Value()
	// }

	obj, _ := vm.Object(`Orcid = {}`)
	obj.Set("login", func(call otto.FunctionCall) otto.Value {
		data, err := api.Login()
		if err != nil {
			return errorObject(nil, fmt.Sprintf("Orcid.login() failed %s, %s", call.CallerLocation(), err))
		}
		result, err := vm.ToValue(data)
		if err != nil {
			return errorObject(nil, fmt.Sprintf("Orcid.login() failed %s, %s", call.CallerLocation(), err))
		}
		return result
	})
	obj.Set("get", func(call otto.FunctionCall) otto.Value {
		if len(call.ArgumentList) != 2 {
			return errorObject(nil, fmt.Sprintf("Missing args Orcid.get(path, orcid) %s", call.CallerLocation()))
		}
		p := call.Argument(0).String()
		orcid := call.Argument(1).String()
		data, err := api.Get(p, orcid)
		if err != nil {
			return errorObject(nil, fmt.Sprintf("Orcid.get(path, orcid) failed %s, %s", call.CallerLocation(), err))
		}
		result, err := vm.ToValue(data)
		if err != nil {
			return errorObject(nil, fmt.Sprintf("Orcid.get(path, orcid) failed %s, %s", call.CallerLocation(), err))
		}
		return result
	})
	obj.Set("getBio", func(call otto.FunctionCall) otto.Value {
		if len(call.ArgumentList) != 1 {
			return errorObject(nil, fmt.Sprintf("Missing arg Orcid.getBio(orcid) %s", call.CallerLocation()))
		}
		orcid := call.Argument(0).String()
		data, err := api.GetBio(orcid)
		if err != nil {
			return errorObject(nil, fmt.Sprintf("Orcid.getBio(orcid) failed %s, %s", call.CallerLocation(), err))
		}
		result, err := vm.ToValue(data)
		if err != nil {
			return errorObject(nil, fmt.Sprintf("Orcid.getBio(orcid) failed %s, %s", call.CallerLocation(), err))
		}
		return result
	})
	obj.Set("getWorks", func(call otto.FunctionCall) otto.Value {
		if len(call.ArgumentList) != 1 {
			return errorObject(nil, fmt.Sprintf("Missing arg Orcid.getWorks(orcid) %s", call.CallerLocation()))
		}
		orcid := call.Argument(0).String()
		data, err := api.GetWorks(orcid)
		if err != nil {
			return errorObject(nil, fmt.Sprintf("Orcid.getWorks(orcid) failed %s, %s", call.CallerLocation(), err))
		}
		result, err := vm.ToValue(data)
		if err != nil {
			return errorObject(nil, fmt.Sprintf("Orcid.getWorks(orcid) failed %s, %s", call.CallerLocation(), err))
		}
		return result
	})
	obj.Set("getProfile", func(call otto.FunctionCall) otto.Value {
		if len(call.ArgumentList) != 1 {
			return errorObject(nil, fmt.Sprintf("Missing arg Orcid.getProfile(orcid) %s", call.CallerLocation()))
		}
		orcid := call.Argument(0).String()
		data, err := api.GetProfile(orcid)
		if err != nil {
			return errorObject(nil, fmt.Sprintf("Orcid.getProfile(orcid) failed %s, %s", call.CallerLocation(), err))
		}
		result, err := vm.ToValue(data)
		if err != nil {
			return errorObject(nil, fmt.Sprintf("Orcid.getProfile(orcid) failed %s, %s", call.CallerLocation(), err))
		}
		return result
	})

	return vm
}

// AddHelp builds the help structures for use in the REPL
func (api *OrcidAPI) AddHelp(js *ostdlib.JavaScriptVM) {
	js.SetHelp("Orcid", "login", []string{}, "log into the API setting up access token")
	js.SetHelp("Orcid", "get", []string{"path string", "orcid string"}, "General API access with provided path and orcid strings")
	js.SetHelp("Orcid", "getBio", []string{"orcid string"}, "Make an orcid bio request")
	js.SetHelp("Orcid", "getWorks", []string{"orcid string"}, "Make an orcid works request")
	js.SetHelp("Orcid", "getProfile", []string{"orcid string"}, "Make an orcid profile request")
}
