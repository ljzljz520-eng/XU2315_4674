package rescue

import (
	"encoding/json"
	"net/http"

	"mountainrescue/web"
)

func NewHTTPServer(service *Service) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/home", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		writeJSON(writer, service.Home())
	})
	mux.HandleFunc("/api/search", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		writeJSON(writer, service.Search(request.URL.Query().Get("q")))
	})
	mux.HandleFunc("/api/equipment-basics", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		writeJSON(writer, map[string]any{"items": service.Basics()})
	})
	mux.HandleFunc("/equipment-basics", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = writer.Write(web.Index())
	})
	mux.Handle("/", http.FileServer(http.FS(web.Assets())))
	return mux
}

func writeJSON(writer http.ResponseWriter, value any) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(writer).Encode(value); err != nil {
		http.Error(writer, "无法生成响应", http.StatusInternalServerError)
	}
}
