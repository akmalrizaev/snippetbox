package main

import (
	"errors" // New import
	"fmt"
	"net/http"
	"strconv"

	// New import
	// New import
	"github.com/akmalrizaev/snippetbox/pkg/forms"  // New import
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
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		// Pass a new empty forms.Form object to the template.
		Form: forms.New(nil),
	})
}

// Change the signature of the createSnippet handler so it is defined as a method
// against *application.
// func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
// if r.Method != http.MethodPost {
// 	w.Header().Set("Allow", http.MethodPost)
// 	app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper.
// 	return
// }

// Checking if the request method is a POST is now superfluous and can be
// removed.

// Create some variables holding dummy data. We'll remove these later on
// during the build.
// title := "O snail"
// content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
// expires := "7"

// Pass the data to the SnippetModel.Insert() method, receiving the
// ID of the new record back.
// id, err := app.snippets.Insert(title, content, expires)
// if err != nil {
// 	app.serverError(w, err)
// 	return
// }

// Redirect the user to the relevant page for the snippet.
// http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)

// Change the redirect to use the new semantic URL style of /snippet/:id
// http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)

// First we call r.ParseForm() which adds any data in POST request bodies
// to the r.PostForm map. This also works in the same way for PUT and PATCH
// requests. If there are any errors, we use our app.ClientError helper to send
// a 400 Bad Request response to the user.
// err := r.ParseForm()
// if err != nil {
// 	app.clientError(w, http.StatusBadRequest)
// 	return
// }

// Use the r.PostForm.Get() method to retrieve the relevant data fields
// from the r.PostForm map.
// title := r.PostForm.Get("title")
// content := r.PostForm.Get("content")
// expires := r.PostForm.Get("expires")

// Initialize a map to hold any validation errors.
// errors := make(map[string]string)

// Check that the title field is not blank and is not more than 100 characters
// long. If it fails either of those checks, add a message to the errors
// map using the field name as the key.
// if strings.TrimSpace(title) == "" {
// 	errors["title"] = "This field cannot be blank"
// } else if utf8.RuneCountInString(title) > 100 {
// 	errors["title"] = "This field is too long (maximum is 100 characters)"
// }

// // Check that the Content field isn't blank.
// if strings.TrimSpace(content) == "" {
// 	errors["content"] = "This field cannot be blank"
// }

// // Check the expires field isn't blank and matches one of the permitted
// // values ("1", "7" or "365").
// if strings.TrimSpace(expires) == "" {
// 	errors["expires"] = "This field cannot be blank"
// } else if expires != "365" && expires != "7" && expires != "1" {
// 	errors["expires"] = "This field is invalid"
// }

// If there are any errors, dump them in a plain text HTTP response and return
// from the handler.
// if len(errors) > 0 {
// 	fmt.Fprint(w, errors)
// 	return
// }

// If there are any validation errors, re-display the create.page.tmpl
// template passing in the validation errors and previously submitted
// r.PostForm data.
// 	if len(errors) > 0 {
// 		app.render(w, r, "create.page.tmpl", &templateData{
// 			FormErrors: errors,
// 			FormData:   r.PostForm,
// 		})
// 		return
// 	}

// 	// Create a new snippet record in the database using the form data.
// 	id, err := app.snippets.Insert(title, content, expires)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}

// 	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
// }

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Create a new forms.Form struct containing the POSTed data from the
	// form, then use the validation methods to check the content.
	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	// If the form isn't valid, redisplay the template passing in the
	// form.Form object as the data.
	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}

	// Because the form data (with type url.Values) has been anonymously embedded
	// in the form.Form struct, we can use the Get() method to retrieve
	// the validated value for a particular form field.
	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Use the Put() method to add a string value ("Your snippet was saved
	// successfully!") and the corresponding key ("flash") to the session
	// data. Note that if there's no existing session for the current user
	// (or their session has expired) then a new, empty, session for them
	// will automatically be created by the session middleware.
	app.session.Put(r, "flash", "Snippet successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	// Parse the form data.
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Validate the form contents using the form helper we made earlier.
	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MaxLength("name", 255)
	form.MaxLength("email", 255)
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)

	// If there are any errors, redisplay the signup form.
	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	}

	// Try to create a new user record in the database. If the email already exists
	// add an error message to the form and re-display it.
	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.Errors.Add("email", "Address is already in use")
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}

	// Otherwise add a confirmation flash message to the session confirming that
	// their signup worked and asking them to log in.
	app.session.Put(r, "flash", "Your signup was successful. Please log in.")

	// And redirect the user to the login page.
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Check whether the credentials are valid. If they're not, add a generic error
	// message to the form failures map and re-display the login page.
	form := forms.New(r.PostForm)
	id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add("generic", "Email or Password is incorrect")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}

	// Add the ID of the current user to the session, so that they are now 'logged
	// in'.
	app.session.Put(r, "authenticatedUserID", id)

	// Redirect the user to the create snippet page.
	http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	// Remove the authenticatedUserID from the session data so that the user is
	// 'logged out'.
	app.session.Remove(r, "authenticatedUserID")
	// Add a flash message to the session to confirm to the user that they've been
	// logged out.
	app.session.Put(r, "flash", "You've been logged out successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
