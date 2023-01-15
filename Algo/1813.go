package algo

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	if len(sentence1) < len(sentence2) {
		sentence1, sentence2 = sentence2, sentence1
	}
	sentence1 = " " + sentence1 + " "
	sentence2 = " " + sentence2 + " "
	head, tail := 1, 1
	for sentence1[head] == sentence2[head] {
		head++
		if head == len(sentence2) {
			return head == len(sentence1) || sentence1[head-1] == ' '
		}
	}
	for sentence1[head] != ' ' || sentence2[head] != ' ' {
		head--
	}
	for ; tail < len(sentence2)-head; tail++ {
		if sentence1[len(sentence1)-tail] != sentence2[len(sentence2)-tail] {
			return false
		}
	}
	return sentence1[len(sentence1)-tail] == ' '
}
