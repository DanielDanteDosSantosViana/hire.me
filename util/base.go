package util

import "github.com/DanielDanteDosSantosViana/hire.me/config"

// estou me baseando conforme a questão do tópico do stackoverflow
//http://stackoverflow.com/questions/742013/how-to-code-a-url-shortener

func InteiroParaString(sequence uint64) (urlEncurtada string) {
	t := make([]byte, 0)

	/* Special case */
	if sequence == 0 {
		return string(config.Conf.Base.Alfabeto)
	}

	/* Map */
	for sequence > 0 {
		r := sequence % uint64(len(config.Conf.Base.Alfabeto))
		t = append(t, config.Conf.Base.Alfabeto[r])
		sequence = sequence / uint64(len(config.Conf.Base.Alfabeto))
	}

	/* Reverse */
	for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}

	return string(t)
}
