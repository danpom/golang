Following tutorial here:  https://go.dev/doc/articles/wiki/


 Recompile the code, and run the app:

$ go build main.go
$ ./main

Visiting http://localhost:8080/view/ANewPage should present you with the page edit form. You should then be able to enter some text, click 'Save', and be redirected to the newly created page.
Other tasks

Here are some simple tasks you might want to tackle on your own:

    Store templates in tmpl/ and page data in data/.
    Add a handler to make the web root redirect to /view/FrontPage.
    Spruce up the page templates by making them valid HTML and adding some CSS rules.
    Implement inter-page linking by converting instances of [PageName] to
    <a href="/view/PageName">PageName</a>. (hint: you could use regexp.ReplaceAllFunc to do this)
