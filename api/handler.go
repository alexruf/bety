package api

import (
	"fmt"
	sq "github.com/elgris/sqrl"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *Api) handleApi() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.WithContext(request.Context()).WithField("URL", request.URL.String()).Debug("handling request")

		sqlStmt := sq.Select("*").From("data")

		rows, err := sqlStmt.RunWith(a.server.Database).Query()
		if err != nil {
			log.WithError(err).WithField("stmt", sqlStmt).Error("error querying database")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			row := struct {
				id      int
				content string
			}{}

			if err := rows.Scan(&row.id, &row.content); err != nil {
				log.WithError(err).Error("error reading from result set")
			}

			if _, err := fmt.Fprintf(writer, "id = %d, content = %s\n", row.id, row.content); err != nil {
				log.WithError(err).Error("error writing response")
			}
		}
	}
}
