/* This file handles dynamically generating a manipulating the
 * Bookshelf Application */

$(document).ready(function() {
    load_books();
});

function load_books() {
    $.getJSON("/bookshelf_ajax", function(result) {
        $.each(result, function(i, book) {
            var col = $(`<div class="col-xs-12 col-sm-6 col-md-4 col-lg-4 card-deck"></div>`);
            var card = $(`<div class="card book-card"></div>`);
            // card_block provides padding
            var card_block = $(`<div class="card-block"></div>`);
            var edit_button = $(`<button type="button" class="btn btn-info edit_button"></button>`);
            edit_button.append($(`<i class="fa fa-pencil-square-o fa-lg" aria-hidden="true"></i>`));
            card_block.append(edit_button);

            card_block.append($(`<h4 class="card-title">` + book.Title + `</h3>`));
            card_block.append($(`<p class="card-text">` + book.Author + `</p>`));
            if(book.CoverImg) {
                var img = $(`<div class="text-center"><img class="img-fluid book-cover" src="data:image;base64, ` + book.CoverImg + `"/></div>`)
                card_block.append(img);
            }
            card_block.append($(`<p><i>Lorem Ipsum is simply dummy text...</i></p>`));
            // book_buffer.append(book_body);
            // card_block.append(edit_button);
            card.append(card_block);
            col.append(card);
            $("#bookshelf-container").append(col);
        });
    });
}