package rescue

import "strings"

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Home() HomePage {
	return s.repository.Home()
}

func (s *Service) Basics() []EquipmentItem {
	return s.repository.BasicEquipment()
}

func (s *Service) Search(query string) SearchResult {
	normalized := strings.ToLower(strings.TrimSpace(query))
	matches := make([]Article, 0)
	if normalized != "" {
		for _, article := range s.repository.Articles() {
			if strings.Contains(strings.ToLower(article.Title), normalized) || strings.Contains(strings.ToLower(article.Summary), normalized) || strings.Contains(strings.ToLower(article.Section), normalized) {
				matches = append(matches, article)
				continue
			}
			for _, keyword := range article.Keywords {
				if strings.Contains(strings.ToLower(keyword), normalized) {
					matches = append(matches, article)
					break
				}
			}
		}
	}
	if len(matches) == 0 {
		grouped := map[string][]EquipmentItem{
			"基础装备": s.repository.BasicEquipment(),
		}
		return SearchResult{Query: query, Articles: matches, Recommended: grouped, StarterListPath: "/equipment-basics"}
	}
	return SearchResult{Query: query, Articles: matches, Recommended: map[string][]EquipmentItem{}, StarterListPath: "/equipment-basics"}
}
