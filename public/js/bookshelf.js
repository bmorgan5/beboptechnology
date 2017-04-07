/* This file handles dynamically generating a manipulating the
 * Bookshelf Application */

$(document).ready(function() {
    load_books();
});

function load_books() {
    $.getJSON("/bookshelf_ajax", function(result) {
        $.each(result, function(i, book) {
            var col = $(`<div class="col-xs-12 col-sm-6 col-md-4"></div>`);
            var card = $(`<div class="book-card"></div>`);
            var container = $(`<div class="book-padding"></div>`);
            var book_header = $(`<div class="book-header"></div>`)
            var edit_button = $(`<button type="button" class="btn btn-info edit_button"></button>`);
            edit_button.append($(`<span class="glyphicon glyphicon-edit"></span>`));
            book_header.append(edit_button);
            book_header.append($(`<p class="bold">` + book.Title + `</p>`));
            book_header.append($(`<p>` + book.Author + `</p>`));
            container.append(book_header);
            if(book.CoverImg) {
                var img = $(`<img class="img-responsive book-cover" src="data:image;base64, ` + book.CoverImg + `"/>`)
                container.append(img);
            }
            container.append($(`<p><i>Lorem Ipsum is simply dummy text...</i></p>`));
            card.append(container);
            col.append(card);
            $("#bookshelf-container").append(col);
        });
    });
}