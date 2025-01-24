package house

var prefixes = []string{"This is", "that"}
var verbs = []string{"lay in", "ate", "killed", "worried", "tossed", "milked", "kissed", "married", "woke", "kept", "belonged to"}
var object = []string{"the house that Jack built", "the malt", "the rat", "the cat", "the dog", "the cow with the crumpled horn", "the maiden all forlorn", "the man all tattered and torn", "the priest all shaven and shorn", "the rooster that crowed in the morn", "the farmer sowing his corn", "the horse and the hound and the horn"}

func Line(line_num int, verse_num int) string {
	if line_num == 1 && verse_num == 1 {
		return prefixes[0] + " " + object[0]
	}

	if line_num == 1 {
		return prefixes[0] + " " + object[verse_num-1]
	}

	line := prefixes[1] + " " + verbs[verse_num-line_num] + " " + object[verse_num-line_num]

	return line
}

func Verse(v int) string {
	verse := ""
	for j := 1; j <= v; j++ {
		if verse != "" {
			verse = verse + "\n"
		}
		verse = verse + Line(j, v)
	}

	return verse + "."
}

func Song() string {
	return MakeSong(12)
}

func MakeSong(verse_num int) string {
	if verse_num == 0 {
		return ""
	}
	verse := Verse(verse_num)
	song := MakeSong(verse_num - 1)
	if song != "" {
		song = song + "\n\n" + verse
	} else {
		song = verse
	}

	return song

}
