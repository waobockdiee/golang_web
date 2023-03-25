package app

func RenderLayout(tmp string) []string {

	file := []string{
		"tmp/" + tmp + ".html",
		"tmp/layout/header.html",
		"tmp/layout/footer.html",
	}
	return file

}
