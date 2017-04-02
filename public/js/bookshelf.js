/* This file handles dynamically generating a manipulating the
 * Bookshelf Application */

$(document).ready(function() {
    //TODO: Load books using ajax
    // load_books(insert_books);
    load_books();
});

function load_books() {
    $.getJSON("/bookshelf_ajax", function(result) {
        $.each(result, function(i, book) {
            var card = $(`<div class="card"></div>`);
            card.append($(`<p> Title: `+book.Title+`</p>`));
            card.append($(`<p> Author: `+book.Author+`</p>`));
            $("#bookshelf-container").append(card);
        });
    });
}