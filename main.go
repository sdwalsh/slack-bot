package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
	"github.com/labstack/echo/v4"
	"github.com/slack-go/slack"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	api := slack.New(os.Getenv("api_token"))
	validationToken := os.Getenv("validation_token")
	signingSecret := os.Getenv("signing_secret")
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		log.Fatalf("Failed to init db %s", err)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi")
	})

	e.GET("/incidents", func(c echo.Context) error {
		txn := db.Txn(false)
		p, err := txn.Get("incidents", "id")
		if err != nil {
			fmt.Printf("error fetching from database: %s\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "cannot fetch from database")
		}

		var i []interface{}
		for row := p.Next(); row != nil; row = p.Next() {
			i = append(i, row)
		}

		return c.JSON(http.StatusOK, i)
	})

	// Callback from Interactivity
	e.POST("message_action", func(c echo.Context) error {
		// Secret Verification Steps
		sv, err := slack.NewSecretsVerifier(c.Request().Header, signingSecret)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to read body")
		}

		_, err = sv.Write(body)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to write body to verify")
		}

		err = sv.Ensure()
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "failed to ensure request came from slack")
		}
		c.Request().Body = ioutil.NopCloser(bytes.NewReader(body))
		// Unmarshal and grab the data we need
		var callback slack.InteractionCallback
		err = json.Unmarshal([]byte(c.FormValue("payload")), &callback)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Unexpected Format")
		}

		// TODO: This is unpleasant. Works for now.
		title := callback.View.State.Values["title-block"]["title"].Value
		desc := callback.View.State.Values["desc-block"]["desc"].Value
		sev := callback.View.State.Values["sev-block"]["sev"].Value

		go createIncident(api, db, title, sev, desc)
		return c.NoContent(http.StatusOK)
	})

	// Command Handler
	e.POST("/command", func(c echo.Context) error {
		s, err := slack.SlashCommandParse(c.Request())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		if !s.ValidateToken(validationToken) {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		switch s.Command {
		case "/rootly":
			switch s.Text {
			case "resolve":
				go resolveIncident(api, db, s)
			case "declare":
				go openModal(api, s)
			default:
			}
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":3000"))
}

func openModal(api *slack.Client, s slack.SlashCommand) {
	_, err := api.OpenView(s.TriggerID, viewRequest)
	if err != nil {
		fmt.Printf("error creating view: %s", err)
	}
}

func createIncident(api *slack.Client, db *memdb.MemDB, title string, sev string, desc string) {
	c, err := api.CreateConversation(title, false)
	if err != nil {
		fmt.Printf("error creating channel: %s\n", err)
	}

	i := Incidents{
		Id:        uuid.New(),
		ChannelId: c.GroupConversation.Conversation.ID,
		Title:     title,
		Desc:      desc,
		Sev:       sev,
		Creator:   c.User,
		Deleted:   false,
	}

	txn := db.Txn(true)
	err = txn.Insert("incidents", i)
	if err != nil {
		fmt.Printf("error saving to database: %s\n", err)
	}
	txn.Commit()
}

func resolveIncident(api *slack.Client, db *memdb.MemDB, s slack.SlashCommand) {
	txn := db.Txn(false)
	raw, err := txn.First("incidents", "id", s.ChannelID)
	if err != nil {
		fmt.Printf("error fetching from database: %s\n", err)
		return
	}
	txn.Commit()

	if raw != nil && !raw.(Incidents).Deleted {
		err = api.ArchiveConversation(s.ChannelID)
		if err != nil {
			fmt.Printf("error deleting channel: %s\n", err)
		}
		// If I understood the documentation, we should make a copy
		i := raw.(Incidents)
		i.Deleted = true
		txn = db.Txn(true)
		err := txn.Insert("incidents", i)
		txn.Commit()
		if err != nil {
			fmt.Printf("error updating record: %s\n")
		}
		return
	}
	fmt.Printf("no channel found, ignoring request\n")
}
