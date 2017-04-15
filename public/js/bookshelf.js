/* This file handles dynamically generating and manipulating the
 * Bookshelf Application */

// This stores an array of book objects returned from the server
var books;

$(document).ready(function() {
    get_books();

    $('#bookshelf').on('click', '.edit_button', function () {
        var id = $(this).data("book_id");
        var book = books[id];

        $('body').addClass("no-scroll");

        var background_cover = $(`<div id="background-cover"></div>`);
        $('body').prepend(background_cover);

        var cover = $(`<div id="page-cover" class="card book-card container"></div>`);
        var cover_block = $(`<div id="cover-block" class="card-block"></div>`);
        var close_button = $(`<button id="close-page" type="button" class="close close-button"><i class="fa fa-times" aria-hidden="true"></i></button>`);
        cover_block.append(close_button);

        var row = $(`<div class="row"></div>`);
        var info_col = $(`<div id="info-col" class="col-sm-12 col-md-4"></div>`);
        var info_block = $(`<div id="info-block" class="card-block text-center"></div>`);
        var title = $(`<h1>` + book.Title + `</h1>`);
        info_block.append(title);

        var author = $(`<h3>` + book.Author + `</h3>`);
        info_block.append(author);

        var cover_img;
        if (book.CoverImg) {
            cover_img = $(`<img class="img-fluid rounded book-cover" src="data:image;base64, ` + book.CoverImg + `"/>`);
        } else {
            cover_img = $(`<div class="container place-holder-img"><i class="center-placeholder fa fa-file-image-o fa-5x" aria-hidden="true"></i></div>`);
        }

        info_block.append(cover_img);
        info_col.append(info_block);
        row.append(info_col);

        blog_col = $(`<div id="blog-col" class="col-sm-12 col-md-8"></div>`);
        blog_block = $(`<div id="blog-block" class="card-block"></div>`);
        blog = $(`<i>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus auctor, felis eget dictum sodales, purus magna aliquam felis, sed suscipit eros massa ac magna. Mauris finibus orci ligula, sed sagittis felis aliquet id. In sit amet arcu nec metus pulvinar consectetur vel sed sapien. Etiam ornare ex justo, vel dictum nulla condimentum quis. Proin feugiat porta mauris ut rutrum. Praesent varius, lorem id aliquet efficitur, nunc eros commodo ipsum, nec finibus ipsum diam at magna. Morbi vel auctor ligula. Etiam placerat risus id dolor imperdiet, vitae euismod nisi commodo. Mauris dictum vestibulum dapibus. Sed non lacus feugiat est suscipit dapibus in et felis.

cras lorem erat, vestibulum quis vulputate non, feugiat eu risus. praesent laoreet nulla sed justo interdum finibus. nulla quis convallis est. phasellus elementum, eros id feugiat condimentum, lectus nisi pharetra nisl, non maximus magna lacus eget lorem. aenean id urna ac sapien laoreet vulputate at ac elit. nulla convallis lorem elit, a tincidunt quam euismod vel. integer tempor tellus et vehicula pretium. pellentesque lobortis non tortor at auctor.

donec gravida, magna a accumsan auctor, augue arcu auctor mauris, et placerat diam justo id orci. curabitur nec sem at felis pretium malesuada. nullam hendrerit vel risus sit amet consequat. vivamus eu pellentesque diam. cras eget dolor eleifend, suscipit leo sit amet, volutpat lectus. in hac habitasse platea dictumst. maecenas pellentesque est quis turpis sollicitudin mattis. nunc quis erat auctor, convallis lacus sit amet, efficitur tellus. vivamus est massa, posuere sit amet neque imperdiet, accumsan blandit dui. sed vitae fringilla neque. vestibulum gravida dolor non mollis lacinia. mauris tincidunt lectus non ante tincidunt, non maximus felis varius. sed finibus diam vel quam fermentum, nec sagittis magna consectetur. curabitur quis sagittis justo, a tincidunt nisi.

sed gravida, enim vel vehicula faucibus, tellus massa maximus ligula, ut hendrerit ligula arcu at turpis. donec ultrices mauris nec elementum porttitor. aliquam erat volutpat. sed ex ex, tincidunt ut enim sit amet, luctus vehicula quam. nunc sit amet orci libero. nulla iaculis dui at urna varius blandit. fusce non magna tristique, egestas orci sed, egestas libero. quisque varius lobortis sem, in semper enim. mauris id diam in ligula rhoncus tempus sed at justo. etiam diam elit, molestie sit amet tincidunt et, dapibus ac tellus. suspendisse in bibendum purus. vivamus leo dolor, dapibus a efficitur quis, consectetur at urna.

curabitur non aliquam velit. suspendisse suscipit, enim sed laoreet varius, tortor sapien cursus nisl, at auctor nisi lorem vitae lectus. donec nec gravida nunc. praesent quis ex massa. proin varius nulla neque, sit amet varius enim consectetur ut. donec ut ante mi. nulla convallis laoreet dictum. fusce molestie tempor libero quis pharetra.
                </i>`);

        blog_block.append(blog);
        blog_col.append(blog_block);
        row.append(blog_col);
        cover_block.append(row);
        cover.append(cover_block);

        $('body').prepend(cover);
   });

    $('body').on('click', '.close-button', function() {
        $('body').removeClass("no-scroll");
        $('#background-cover').remove();
        $('#page-cover').remove();
    });
});

// Queries the server for all the books in the bookshelf and saves it
// into the global books variable
// TODO: Possibly make this multiple ajax calls so the page can load faster
function get_books() {
    $.getJSON("/bookshelf_ajax", function(result) {
        books = result;
        $.each(books, function (i, book) {
            var col = $(`<div class="col-xs-12 col-sm-6 col-md-4 col-lg-3 card-deck"></div>`);
            var card = $(`<div class="card book-card"></div>`);
            // card_block provides padding
            var card_block = $(`<div class="card-block"></div>`);
            var edit_button = $(`<button type="button" data-book_id="` + i + `" class="btn btn-info edit_button"></button>`);
            edit_button.append($(`<i class="fa fa-pencil-square-o fa-fw" aria-hidden="true"></i>`));
            card_block.append(edit_button);
            card_block.append($(`<h6 class="card-title book-title">` + book.Title + `</h6>`));
            card_block.append($(`<p class="card-text">` + book.Author + `</p>`));
            if (book.CoverImg) {
                var img = $(`<div class="text-center"><img class="img-fluid rounded book-cover-thumbnail" src="data:image;base64, ` + book.CoverImg + `"/></div>`);
                card_block.append(img);
            }
            var opinion = $(`<p class="opinion"></p>`);
            var preview = $(`<i>Lorem Ipsum is simply dummy text...</i>`);
            var read_more = $(`<a href="#" class="read-more">read more...</a>`);
            opinion.append(preview);
            opinion.append(read_more);
            card_block.append(opinion);
            card.append(card_block);
            col.append(card);
            $("#bookshelf").append(col);
        });
    });
}