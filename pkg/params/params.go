package params

import "log"

var (
    _params = map[string]struct{}{} // register params
)

func NewParams(param string) string {
    if param == "" {
        log.Panic("params must not null")
    }

    return addParams(param)
}

func addParams(param string) string {
    if _, ok := _params[param]; ok {
        log.Panicf("params exists %s", param)
    }

    _params[param] = struct{}{}
    return param
}
