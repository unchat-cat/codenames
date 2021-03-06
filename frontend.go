package codenames

import (
	"math/rand"
	"net/http"
	"path/filepath"
	"strings"
)

const tpl = `
<!DOCTYPE html>
<html>
    <head>
        <title>Codi secret - Jugar en línia</title>
        <script src="/static/app.js" type="text/javascript"></script>
        <link href="https://fonts.googleapis.com/css?family=Roboto" rel="stylesheet">
        <link rel="stylesheet" type="text/css" href="/static/game.css" />
        <link rel="stylesheet" type="text/css" href="/static/lobby.css" />
        <link rel="shortcut icon" type="image/png" href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAwSURBVHgB7dJRDQAABEVRJFJBNvX8SEQINrO9G+B8XTaPokFCwwAsAPdxqmKk90ADdlUE2gRVHXcAAAAASUVORK5CYII="/>

        <script type="text/javascript">
             {{if .SelectedGameID}}
             window.selectedGameID = "{{.SelectedGameID}}";
             {{end}}
             window.autogeneratedGameID = "{{.AutogeneratedGameID}}";
        </script>
    </head>
    <body>
		<div id="app">
		</div>
		<h1><a href="https://unchat.cat" style="color:black;text-decoration:none;text-transform:none;">unchat.cat</a></h1>
    </body>
</html>
`

type templateParameters struct {
	SelectedGameID      string
	AutogeneratedGameID string
}

func (s *Server) handleIndex(rw http.ResponseWriter, req *http.Request) {
	dir, id := filepath.Split(req.URL.Path)
	if dir != "" && dir != "/" {
		http.NotFound(rw, req)
		return
	}

	autogeneratedID := s.getAutogeneratedID()

	err := s.tpl.Execute(rw, templateParameters{
		SelectedGameID:      id,
		AutogeneratedGameID: autogeneratedID,
	})
	if err != nil {
		http.Error(rw, "error rendering", http.StatusInternalServerError)
	}
}

func (s *Server) getAutogeneratedID() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	autogeneratedID := ""
	for {
		autogeneratedID = strings.ToLower(s.gameIDWords[rand.Intn(len(s.gameIDWords))])
		if _, ok := s.games[autogeneratedID]; !ok {
			break
		}
	}
	return autogeneratedID
}
