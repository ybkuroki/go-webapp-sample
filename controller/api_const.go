package controller

const (
	// API represents the group of API.
	API = "/api"
	// APIBook represents the group of book management API.
	APIBook = API + "/book"
	// APIBookList is
	APIBookList = APIBook + "/list"
	// APIBookSearch is
	APIBookSearch = APIBook + "/search"
	// APIBookRegist is
	APIBookRegist = APIBook + "/new"
	// APIBookEdit is
	APIBookEdit = APIBook + "/edit"
	// APIBookDelete is
	APIBookDelete = APIBook + "/delete"
)

const (
	// APIMaster represents the group of master management API.
	APIMaster = API + "/master"
	// APIMasterCategory is
	APIMasterCategory = APIMaster + "/category"
	// APIMasterFormat is
	APIMasterFormat = APIMaster + "/format"
)

const (
	// APIAccount represents the group of account management API.
	APIAccount = API + "/account"
	// APIAccountLoginStatus is
	APIAccountLoginStatus = APIAccount + "/loginStatus"
	// APIAccountLoginAccount is
	APIAccountLoginAccount = APIAccount + "/loginAccount"
)

const (
	// APIHealth is
	APIHealth = API + "/health"
)
