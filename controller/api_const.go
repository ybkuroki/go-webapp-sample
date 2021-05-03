package controller

const (
	// API represents the group of API.
	API = "/api"
	// APIBook represents the group of book management API.
	APIBook = API + "/book"
	// APIBookList represents the API to get book's list.
	APIBookList = APIBook + "/list"
	// APIBookSearch represents the API to search book's list.
	APIBookSearch = APIBook + "/search"
	// APIBookRegist represents the API to register a new book.
	APIBookRegist = APIBook + "/new"
	// APIBookEdit represents the API to edit the existing book.
	APIBookEdit = APIBook + "/edit"
	// APIBookDelete represents the API to delete the existing book.
	APIBookDelete = APIBook + "/delete"
)

const (
	// APIMaster represents the group of master management API.
	APIMaster = API + "/master"
	// APIMasterCategory represents the API to get category's list.
	APIMasterCategory = APIMaster + "/category"
	// APIMasterFormat represents the API to get format's list.
	APIMasterFormat = APIMaster + "/format"
)

const (
	// APIAccount represents the group of account management API.
	APIAccount = API + "/account"
	// APIAccountLoginStatus represents the API to get the status of logged in account.
	APIAccountLoginStatus = APIAccount + "/loginStatus"
	// APIAccountLoginAccount represents the API to get the logged in account.
	APIAccountLoginAccount = APIAccount + "/loginAccount"
	// APIAccountLogin represents the API to login by session authentication.
	APIAccountLogin = APIAccount + "/login"
	// APIAccountLogout represents the API to logout.
	APIAccountLogout = APIAccount + "/logout"
)

const (
	// APIHealth represents the API to get the status of this application.
	APIHealth = API + "/health"
)
