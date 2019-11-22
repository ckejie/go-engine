package console

import "container/list"

type EditBox struct {
	str           string
	cur           int
	historyMaxLen int
	history       *list.List
	historyIndex  int
	enterstr      string
}

func NewEditBox(historyMaxLen int) *EditBox {
	return &EditBox{
		historyMaxLen: historyMaxLen,
		history:       list.New(),
		historyIndex:  -1,
	}
}

func (eb *EditBox) Input(key EventKey) {
	if key.Key() == KeyRune {
		i := string(key.Rune())
		eb.str = eb.str[0:eb.cur] + i + eb.str[eb.cur:]
		eb.cur++
	} else if key.Key() == KeyBackspace {
		if eb.cur > 0 {
			eb.str = eb.str[0:eb.cur-1] + eb.str[eb.cur:]
			eb.cur--
		}
	} else if key.Key() == KeyDelete {
		if eb.cur < len(eb.str) {
			eb.str = eb.str[0:eb.cur] + eb.str[eb.cur+1:]
		}
	} else if key.Key() == KeyLeft {
		if eb.cur > 0 {
			eb.cur--
		}
	} else if key.Key() == KeyRight {
		if eb.cur > 0 {
			eb.cur++
		}
	} else if key.Key() == KeyUp {
		if eb.historyIndex < eb.history.Len()-1 {
			eb.historyIndex++
			index := 0
			for e := eb.history.Front(); e != nil; e = e.Next() {
				if index >= eb.historyIndex {
					h := e.Value.(string)
					eb.str = h
					eb.cur = len(eb.str)
					break
				}
				index++
			}
		}
	} else if key.Key() == KeyDown {
		if eb.historyIndex > 0 {
			eb.historyIndex--
			index := 0
			for e := eb.history.Front(); e != nil; e = e.Next() {
				if index >= eb.historyIndex {
					h := e.Value.(string)
					eb.str = h
					eb.cur = len(eb.str)
					break
				}
			}
		} else if eb.historyIndex == 0 {
			eb.historyIndex = -1
			eb.str = ""
			eb.cur = len(eb.str)
		}
	} else if key.Key() == KeyEnter {
		eb.saveText()
	}
}

func (eb *EditBox) saveText() {
	str := eb.GetText()
	eb.cur = 0
	eb.str = ""
	eb.historyIndex = -1

	hasHistory := false
	for e := eb.history.Front(); e != nil; e = e.Next() {
		h := e.Value.(string)
		if h == str {
			hasHistory = true
		}
	}
	if !hasHistory && len(str) > 0 {
		eb.history.PushFront(str)
		if eb.history.Len() > eb.historyMaxLen {
			var last *list.Element
			for e := eb.history.Front(); e != nil; e = e.Next() {
				last = e
			}
			if last != nil {
				eb.history.Remove(last)
			}
		}
	}
	eb.enterstr = str
}

func (eb *EditBox) GetEnterText() string {
	return eb.enterstr
}

func (eb *EditBox) GetText() string {
	return eb.str
}

func (eb *EditBox) GetShowText() string {
	return ""

}