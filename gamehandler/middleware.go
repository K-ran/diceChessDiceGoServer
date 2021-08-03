package gamehandler

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// checks is the player exists in the db
func checkValidPlayer(s http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		playerId := params["playerId"]
		_, err := player_db.Get(playerId)
		if err != nil {
			ReturnFailure(rw, "Invalid player id.")
			return
		}
		s.ServeHTTP(rw, r)
	})
}

//checks if the valid game exists in the
func checkValidGame(s http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		gameId := params["gameId"]
		_, err := game_db.Get(gameId)
		if err != nil {
			ReturnFailure(rw, "Invalid game id.")
			return
		}
		s.ServeHTTP(rw, r)
	})
}

//checks if the valid game exists in the
func allowCORS(s http.HandlerFunc, value string) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", value)

		if r.Method == http.MethodOptions {
			return
		}
		s.ServeHTTP(rw, r)
	})
}

type recaptchaTokenStruct struct {
	Token string `json:"captchaToken"`
}

const siteVerifyURL = "https://www.google.com/recaptcha/api/siteverify"

type SiteVerifyResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func validateReCaptcha(s http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var t recaptchaTokenStruct

		err := decoder.Decode(&t)
		if err != nil {
			ReturnFailure(rw, "Invalid captcha Token")
			return
		}
		secret := os.Getenv("RECAPTCHA_KEY")
		if secret == "" {
			ReturnFailure(rw, "Invalid captcha key")
			return
		}
		req, err := http.NewRequest(http.MethodPost, siteVerifyURL, nil)
		if err != nil {
			ReturnFailure(rw, "Recaptcha, failed to create new request")
			return
		}

		// Add necessary request parameters.
		q := req.URL.Query()
		q.Add("secret", secret)
		q.Add("response", t.Token)

		// Make request
		req.URL.RawQuery = q.Encode()
		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			ReturnFailure(rw, "Failed to request recaptcha api")
			return
		}
		defer resp.Body.Close()

		// Decode response.
		var body SiteVerifyResponse
		if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
			ReturnFailure(rw, "Failed to decode recaptcha response")
			return
		}

		if !body.Success {
			ReturnFailure(rw, "Recaptcha Verification Failed")
			return
		}

		s.ServeHTTP(rw, r)
	})
}
