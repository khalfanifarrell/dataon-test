package app

import (
	"database/sql"
	"dataon-test/entity"
	"dataon-test/presentation"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	app.Router.
		Methods("GET").
		Path("/applicants/{id}").
		HandlerFunc(app.getOneApplicant)

	app.Router.
		Methods("POST").
		Path("/applicants").
		HandlerFunc(app.insertApplicant)

	app.Router.
		Methods("PATCH").
		Path("/applicants/{id}").
		HandlerFunc(app.updateApplicant)

	// app.Router.
	// 	Methods("DELETE").
	// 	Path("/applicants/{id}").
	// 	HandlerFunc(app.updateApplicant)
}

func (app *App) getOneApplicant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No ID in the path")
	}

	dbdata := &entity.Applicants{}
	err := app.Database.QueryRow(`
		SELECT
			id,
			first_name,
			last_name,
			email,
			phone,
			home_address,
			title,
			years_of_exp,
			created_at,
			deleted_at
		FROM applicants
		WHERE
			id = ?
			AND deleted_at IS NULL`, id).Scan(
		&dbdata.ID,
		&dbdata.FirstName,
		&dbdata.LastName,
		&dbdata.Email,
		&dbdata.Phone,
		&dbdata.HomeAddress,
		&dbdata.Title,
		&dbdata.YearsOfExp,
		&dbdata.CreatedAt,
		&dbdata.DeletedAt,
	)
	if err == sql.ErrNoRows {
		log.Printf("no data found with id: %s", id)
		return
	}

	if err != nil {
		log.Fatal("Database SELECT failed. Error: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dbdata); err != nil {
		panic(err)
	}
}

func (app *App) insertApplicant(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestBody presentation.ApplicantRequest
	err := decoder.Decode(&requestBody)
	if err != nil {
		panic(err)
	}

	fmt.Println("Request Body", requestBody)

	stmt, err := app.Database.Prepare(`
		INSERT INTO applicants
			(first_name, last_name, email, phone, home_address, title, years_of_exp) VALUES(?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(requestBody.FirstName, requestBody.LastName, requestBody.Email, requestBody.Phone, requestBody.HomeAddress, requestBody.Title, requestBody.YearsOfExp)
	if err != nil {
		log.Fatal("Database INSERT failed. Error: ", err)
	}

	w.WriteHeader(http.StatusOK)
}

func (app *App) updateApplicant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No ID in the path")
	}

	var count int

	err := app.Database.QueryRow(`
		SELECT COUNT(*)
		FROM applicants
		WHERE
			id = ?
			AND deleted_at IS NULL`, id).Scan(&count)
	if err != nil {
		log.Fatal("Database COUNT failed. Error: ", err)
	}

	if count < 1 {
		log.Printf("No id found with %s", id)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var requestBody presentation.ApplicantRequest
	err = decoder.Decode(&requestBody)
	if err != nil {
		panic(err)
	}

	stmt, err := app.Database.Prepare(`
		UPDATE applicants
		SET first_name = ?,
			last_name = ?
		WHERE id = ?`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(requestBody.FirstName, requestBody.LastName, id)
	if err != nil {
		log.Fatal("Database UPDATE failed. Error: ", err)
	}

	w.WriteHeader(http.StatusOK)
}

// func (app *App) deleteApplicant(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, ok := vars["id"]
// 	if !ok {
// 		log.Fatal("No ID in the path")
// 	}

// 	var count int

// 	err := app.Database.QueryRow(`
// 		SELECT COUNT(*)
// 		FROM applicants
// 		WHERE
// 			id = ?
// 			AND deleted_at IS NULL`, id).Scan(&count)
// 	if err != nil {
// 		log.Fatal("Database COUNT failed. Error: ", err)
// 	}

// 	if count < 1 {
// 		log.Printf("No id found with %s", id)
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	var requestBody presentation.ApplicantRequest
// 	err = decoder.Decode(&requestBody)
// 	if err != nil {
// 		panic(err)
// 	}

// 	stmt, err := app.Database.Prepare(`
// 		UPDATE applicants
// 		SET first_name = ?,
// 			last_name = ?
// 		WHERE id = ?`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = stmt.Exec(requestBody.FirstName, requestBody.LastName, id)
// 	if err != nil {
// 		log.Fatal("Database UPDATE failed. Error: ", err)
// 	}

// 	w.WriteHeader(http.StatusOK)
// }
