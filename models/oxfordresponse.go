package models

type OxfordResponse struct {
	Results []result
}

type result struct {
	LexicalEntries []lexicalEntry
	Language string
}

type lexicalEntry struct {
	LexicalCategory string
	Entries []entry
}

type entry struct {
	Senses []sense
}

type sense struct {
	Definitions []string
}

func (response *OxfordResponse) GetDefinitions() []string{
	definitionList := make([]string,0)
	//TODO Abeg find a more elegant way for this piece of rubbish later
	for _,result := range response.Results {
		for _,lexicalentry := range result.LexicalEntries{
			for _, entry := range lexicalentry.Entries{
				for _,sense := range entry.Senses{
					for _,definition := range sense.Definitions{
						definitionList = append(definitionList,definition)
					}
				}
			}
		}
	}
	return  definitionList
}