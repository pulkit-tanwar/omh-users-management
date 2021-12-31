package constant

const (
	FailedToReadRequestBody        = 50001
	FailedToReadRequestBodyMessage = "Failed to Read Request Body"

	UserNameMissing        = 50002
	UserNameMissingMessage = "User Name is missing"

	FirstNameMissing        = 50003
	FirstNameMissingMessage = "First Name is missing"

	LastNameMissing        = 50004
	LastNameMissingMessage = "Last Name is missing"

	FailedToCreateUser        = 50005
	FailedToCreateUserMessage = "User Creation Failed"

	ContentTypeNotSupported        = 50006
	ContentTypeNotSupportedMessage = "Invalid content-type. Only application/json is supported"

	UserNameAlreadyExists        = 50007
	UserNameAlreadyExistsMessage = "UserName Already Exists"

	FailedToUnmarshalRequestBody        = 50008
	FailedToUnmarshalRequestBodyMessage = "Internal Server Error"

	FailedToReadItemFromDB        = 50009
	FailedToReadItemFromDBMessage = "Cannot process request due to internal Server Error"

	UserNotFound        = 50010
	UserNotFoundMessage = "User Not Found"

	ModifiableFieldNotPresent        = 50011
	ModifiableFieldNotPresentMessage = "Modifiable fields are not present"

	FailedToAddItemToDB        = 50012
	FailedToAddItemToDBMessage = "Cannot process request due to internal Server Error"

	FailedToDeleteItemFromDB        = 50013
	FailedToDeleteItemFromDBMessage = "Cannot process request due to internal Server Error"
)
