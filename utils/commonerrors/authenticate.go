package commonerrors

import "errors"

var ErrIsLocal = errors.New("is local")
var ErrIsRemote = errors.New("is remote")
var ErrNotLoggedIn = errors.New("not logged in")

var ErrFormInvalid = errors.New("form invalid")
var ErrCheckActorDataNotSameDomain = errors.New("data id and actor id is not on the same domain")
