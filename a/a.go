package a

import "packageheights/b"

func Name() string {
	return "a" + b.Name()
}
