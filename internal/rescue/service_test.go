package rescue

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeShowsSafetyFeaturedAndTeam(t *testing.T) {
	service := NewService(NewMemoryRepository())
	home := service.Home()
	if home.Title != "山地救援装备笔记" {
		t.Fatalf("got title %q", home.Title)
	}
	if len(home.SafetyTips) != 3 || len(home.Featured) != 3 || len(home.AuthorTeam) != 3 {
		t.Fatalf("got safety=%d featured=%d team=%d", len(home.SafetyTips), len(home.Featured), len(home.AuthorTeam))
	}
	if strings.Join(home.Sections, ",") != "绳索,通信,保暖,导航,案例复盘" {
		t.Fatalf("got sections %v", home.Sections)
	}
}

func TestSearchFindsCommunicationNotes(t *testing.T) {
	service := NewService(NewMemoryRepository())
	result := service.Search("对讲机")
	if len(result.Articles) != 1 || result.Articles[0].Section != "通信" {
		t.Fatalf("got articles %v", result.Articles)
	}
}

func TestSearchMissingKeywordProvidesStarterKit(t *testing.T) {
	server := NewHTTPServer(NewService(NewMemoryRepository()))
	request := httptest.NewRequest(http.MethodGet, "/api/search?q=不存在的装备", nil)
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("got status %d", response.Code)
	}
	if !strings.Contains(response.Body.String(), "基础装备") || !strings.Contains(response.Body.String(), "/equipment-basics") {
		t.Fatalf("body did not include starter kit: %s", response.Body.String())
	}
}

func TestSearchEndpointReturnsMatchingArticle(t *testing.T) {
	server := NewHTTPServer(NewService(NewMemoryRepository()))
	request := httptest.NewRequest(http.MethodGet, "/api/search?q=导航", nil)
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("got status %d", response.Code)
	}
	if !strings.Contains(response.Body.String(), "山脊路线的导航复盘") {
		t.Fatalf("body did not include matching article: %s", response.Body.String())
	}
}

func TestBasicsPageShowsEquipmentInterface(t *testing.T) {
	server := NewHTTPServer(NewService(NewMemoryRepository()))
	request := httptest.NewRequest(http.MethodGet, "/equipment-basics", nil)
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("got status %d", response.Code)
	}
	if !strings.Contains(response.Body.String(), "山地救援装备笔记") || !strings.Contains(response.Body.String(), "app.js") {
		t.Fatalf("body did not include equipment interface: %s", response.Body.String())
	}
}
