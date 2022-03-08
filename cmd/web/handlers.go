package main

import (
	"errors" // New import
	"fmt"
	"net/http"
	"strconv"

	"github.com/akmalrizaev/snippetbox/pkg/models" // New import
)

// Change the signature of the home handler so it is defined as a method against
// *application.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	app.notFound(w) // Use the notFound() helper
	// 	return
	// }

	// Because Pat matches the "/" path exactly, we can now remove the manual check
	// of r.URL.Path != "/" from this handler.

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Use the new render helper.
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})

	// Create an instance of a templateData struct holding the slice of
	// snippets.
	// data := &templateData{Snippets: s}

	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	// // Pass in the templateData struct when executing the template.
	// err = ts.Execute(w, data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

	// for _, snippet := range s {
	// 	fmt.Fprintf(w, "%v\n", snippet)
	// }

	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	// Because the home handler function is now a method against application
	// 	// it can access its fields, including the error logger. We'll write the log
	// 	// message to this instead of the standard logger.
	// 	app.errorLog.Println(err.Error())
	// 	app.serverError(w, err) // Use the serverError() helper.
	// 	return
	// }

	// err = ts.Execute(w, nil)
	// if err != nil {
	// 	// Also update the code here to use the error logger from the application
	// 	// struct.
	// 	app.errorLog.Println(err.Error())
	// 	http.Error(w, "Internal Server Error", 500)
	// }
}

// Change the signature of the showSnippet handler so it is defined as a method
// against *application.
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// if err != nil || id < 1 {
	// 	app.notFound(w) // Use the notFound() helper.
	// 	return
	// }

	// Pat doesn't strip the colon from the named capture key, so we need to
	// get the value of ":id" from the query string instead of "id".
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// Use the SnippetModel object's Get method to retrieve the data for a
	// specific record based on its ID. If no matching record is found,
	// return a 404 Not Found response.
	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// Use the new render helper.
	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})
	// Initialize a slice containing the paths to the show.page.tmpl file,
	// plus the base layout and footer partial that we made earlier.

	// Create an instance of a templateData struct holding the snippet data.
	// data := &templateData{Snippet: s}
	// files := []string{
	// 	"./ui/html/show.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// // Parse the template files...
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	// And then execute them. Notice how we are passing in the snippet
	// data (a models.Snippet struct) as the final parameter.
	// err = ts.Execute(w, s)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

	// Pass in the templateData struct when executing the template.
	// err = ts.Execute(w, data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

	// Write the snippet data as a plain-text HTTP response body.
	//fmt.Fprintf(w, "%v", s)
}

// Add a new createSnippetForm handler, which for now returns a placeholder response.
func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

// Change the signature of the createSnippet handler so it is defined as a method
// against *application.
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	w.Header().Set("Allow", http.MethodPost)
	// 	app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper.
	// 	return
	// }

	// Checking if the request method is a POST is now superfluous and can be
	// removed.

	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := "7"

	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Redirect the user to the relevant page for the snippet.
	// http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)

	// Change the redirect to use the new semantic URL style of /snippet/:id
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
