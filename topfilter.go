package password

import "strings"

func (mpe *Eval) TopDictFilter(pwd string) (bool, string, error) {

	if strings.Contains(Top1000Pwd, pwd) {
		return false, "", nil
	} else {
		return true, "", nil
	}

}
