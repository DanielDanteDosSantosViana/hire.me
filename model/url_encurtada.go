type UrlEncurtador struct {
	Alias string `json:"alias"`
	URL     Router `json:"url"`
	Estatistica Estatistica `json:"estatistica"`

}

type Estatistica struct {
	TotalCaracteresUrl string `json:"caracteresUrl"`
	TotalCaracteresEncurtada   string `json:"caracteresUrlEncurtada"`
	TempoOperacao  string `json:"tempoOperacao"`
}
