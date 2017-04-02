/* This file handles dynamically generating a manipulating the
 * Bookshelf Application */

$(document).ready(function() {
    load_books();
});

function load_books() {
    $.getJSON("/bookshelf_ajax", function(result) {
        $.each(result, function(i, book) {
            var col = $(`<div class="col-xs-12 col-sm-6 col-md-3"></div>`);
            var card = $(`<div class="card"></div>`);
            if(book.CoverImg) {
                var img = $(`<img class="img-responsive" src="data:image;base64, ` + book.CoverImg + `"/>`)
                card.append(img);
            }
            card.append($(`<p> Title: `+book.Title+`</p>`));
            card.append($(`<p> Author: `+book.Author+`</p>`));
            col.append(card);
            $("#bookshelf-container").append(col);
        });
    });
}