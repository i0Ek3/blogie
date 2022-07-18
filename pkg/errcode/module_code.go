package errcode

var (
	ErrorGetTagListFail = NewError(200101, "Failed to get tag list")
	ErrorCreateTagFail  = NewError(200102, "Failed to create tag")
	ErrorUpdateTagFail  = NewError(200103, "Failed to update tag")
	ErrorDeleteTagFail  = NewError(200104, "Failed to delete tag")
	ErrorCountTagFail   = NewError(200105, "Failed to count tag")

	ErrorGetArticleFail    = NewError(200201, "Failed to fetch an article")
	ErrorGetArticlesFail   = NewError(200202, "Failed to fetch the articles")
	ErrorCreateArticleFail = NewError(200203, "Failed to create an article")
	ErrorUpdateArticleFail = NewError(200204, "Failed to update an article")
	ErrorDeleteArticleFail = NewError(200205, "Failed to delete an article")

	ErrorUploadFileFail = NewError(200301, "Failed to upload file")
)
