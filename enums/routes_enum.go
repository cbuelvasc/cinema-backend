package enums

const (
	BasePath = "/api/cinema/v1"

	SignIn = "/signin"
	SignUp = "/signup"

	GetUsers       = "/users"
	GetUserById    = "/users/:id"
	UpdateUserById = "/users/:id"
	DeleteUserById = "/users/:id"

	GetMovies       = "/movies"
	CreateMovie     = "/movies"
	GetMovieById    = "/movies/:id"
	UpdateMovieById = "/movies/:id"
	DeleteMovieById = "/movies/:id"

	GetCountries      = "/countries"
	CreateCountry     = "/countries"
	GetCountryById    = "/countries/:id"
	UpdateCountryById = "/countries/:id"
	DeleteCountryById = "/countries/:id"

	GetStates       = "/states"
	CreateState     = "/states"
	GetStateById    = "/states/:id"
	UpdateStateById = "/states/:id"
	DeleteStateById = "/states/:id"

	GetCities      = "/cities"
	CreateCity     = "/cities"
	GetCityById    = "/cities/:id"
	UpdateCityById = "/cities/:id"
	DeleteCityById = "/cities/:id"

	GetTweets       = "/tweets"
	CreateTweets    = "/tweets"
	GetTweetById    = "/tweets/:id"
	UpdateTweetById = "/tweets/:id"
	DeleteTweetById = "/tweets/:id/user/:userId"
)
