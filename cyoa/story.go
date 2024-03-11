package cyoa

type story map[string]Chapter

type Chapter struct {
	Title     string   `json:"title"`
	Paragraph []string `json:"story"`
	Options   []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

// func main() {
// 	chapters, err := ParseJson()
// 	if err != nil {
// 		panic(err)
// 	}
// 	introArc := chapters["intro"]
// 	for _, v := range introArc.Options {
// 		fmt.Println(v.Text)
// 		fmt.Println(v.Chapter)
// 	}
// }

// func ParseJson() (map[string]Chapter, error) {
// 	var result map[string]Chapter
// 	byteValue, err := readJsonFromFile("story.json")
// 	if err != nil {
// 		return nil, err
// 	}
// 	json.Unmarshal([]byte(byteValue), &result)

// 	return result, nil
// }

// func readJsonFromFile(filename string) ([]byte, error) {
// 	jsonFile, err := os.Open(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer jsonFile.Close()

// 	byteValue, _ := ioutil.ReadAll(jsonFile)
// 	return byteValue, nil
// }
