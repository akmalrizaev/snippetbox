package main

import (
	"net/http"

	"github.com/bmizerany/pat" // New import

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	// standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// Create a new middleware chain containing the middleware specific to
	// our dynamic application routes. For now, this chain will only contain
	// the session middleware but we'll add more to it later.
	// dynamicMiddleware := alice.New(app.session.Enable)

	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(app.home))
	// mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	// mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	// mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet)) // Moved down

	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	// mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// Update these routes to use the new dynamic middleware chain followed
	// by the appropriate handler function.
	// mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	// mux.Get("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm))
	// mux.Post("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippet))
	// mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	// Add the five new routes.
	// mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	// mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	// mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	// mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	// mux.Post("/user/logout", dynamicMiddleware.ThenFunc(app.logoutUser))

	// Leave the static files route unchanged.
	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	// mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// return standardMiddleware.Then(mux)

	// standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	// dynamicMiddleware := alice.New(app.session.Enable)

	// mux := pat.New()
	// mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	// // Add the requireAuthentication middleware to the chain.
	// mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippetForm))
	// // Add the requireAuthentication middleware to the chain.
	// mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippet))
	// mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	// mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	// mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	// mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	// mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	// // Add the requireAuthentication middleware to the chain.
	// mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	// mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// return standardMiddleware.Then(mux)

	// standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	// // Use the nosurf middleware on all our 'dynamic' routes.
	// dynamicMiddleware := alice.New(app.session.Enable, noSurf)

	// mux := pat.New()
	// mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	// mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippetForm))
	// mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippet))
	// mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	// mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	// mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	// mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	// mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	// mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	// mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// return standardMiddleware.Then(mux)

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	// Add the authenticate() middleware to the chain.
	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	// Add a new GET /ping route.
	mux.Get("/ping", http.HandlerFunc(ping))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
