package models

type UrlEncurtada struct {
	Alias       string      `json:"alias"`
	URL         string      `json:"url"`
	Estatistica Estatistica `json:"estatistica"`
}

type Estatistica struct {
	TotalCaracteresUrl       string `json:"caracteresUrl"`
	TotalCaracteresEncurtada string `json:"caracteresUrlEncurtada"`
	TempoOperacao            string `json:"tempoOperacao"`
}
