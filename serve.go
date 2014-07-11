package view

import (
	"reflect"
	// "strings"
)

var (
	endpointType = reflect.TypeOf(Endpoint{})
	pathArgType  = reflect.TypeOf((*pathArg)(nil)).Elem()
)

// func Serve(site interface{}) error {
// 	return ServeConfig(site, &Config)
// }

// func ServeConfig(site interface{}, config *Configuration) error {
// 	err := config.Init()
// 	if err != nil {
// 		return err
// 	}

// 	v := reflect.ValueOf(site)
// 	initSiteStructRecursive(v, "/")

// 	return nil
// }

// func initSiteStructRecursive(v reflect.Value, path string) {
// 	if v.Kind() != reflect.Struct {
// 		panic("site must be a struct with view.Endpoint members")
// 	}
// 	t := v.Type()
// 	for i := 0; i < t.NumField(); i++ {
// 		fv := v.Field(i)
// 		f := t.Field(i)
// 		name := f.Tag.Get("name")
// 		if name == "" {
// 			name = strings.ToLower(f.Name)
// 		}
// 		switch {
// 		case f.Type == endpointType:
// 			endpoint := fv.Addr().Interface().(*Endpoint)
// 			endpoint.URL = URL(path + name + "/")

// 		case f.Type.Implements(pathArgType):

// 		case f.Type.Kind() == reflect.Struct:
// 			initSiteStructRecursive(fv, path+name+"/")

// 		default:
// 			panic("site struct members must be view.Endpoint or view.StringArg")
// 		}
// 	}
// }
