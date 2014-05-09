package hpack

type HeaderSet struct {
	Headers []HeaderField
}

func (hs HeaderSet) Len() int {
	return len(hs.Headers)
}

func (hs HeaderSet) Less(i, j int) bool {
	if hs.Headers[i].Name != hs.Headers[j].Name {
		return hs.Headers[i].Name < hs.Headers[j].Name
	}

	return hs.Headers[i].Value < hs.Headers[j].Value
}

func (hs HeaderSet) Swap(i, j int) {
	hs.Headers[i], hs.Headers[j] = hs.Headers[j], hs.Headers[i]
}
