

bookstore/bookstore.proto	bookstoregoogle/api/annotations.proto"C
CreateBookRequest
title (	Rtitle
content (	Rcontent"F
Book
id (	Rid
title (	Rtitle
content (	Rcontent"
GetBookListRequest"/
BookList#
list (2.bookstore.BookRlist2∑
BookStoreServiceQ

CreateBook.bookstore.CreateBookRequest.bookstore.Book"Ç”‰ì"	/v1/books:*P
GetBook.bookstore.GetBookListRequest.bookstore.BookList"Ç”‰ì	/v1/booksBZservices/bookstorebproto3
≥
cart/cart.protocartgoogle/api/annotations.proto"B
AddToCartRequest
title (	Rtitle
content (	Rcontent"F
Item
id (	Rid
title (	Rtitle
content (	Rcontent"
GetCartRequest"*
CartList
list (2
.cart.ItemRlist2ò
CartServiceE
	AddToCart.cart.AddToCartRequest
.cart.Item"Ç”‰ì"	/v1/carts:*B
GetCart.cart.GetCartRequest.cart.CartList"Ç”‰ì	/v1/cartsBZservices/cartbproto3