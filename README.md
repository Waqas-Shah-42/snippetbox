# snippetbox

## Routes
| Method | Pattern         | Handler        | Action                      |
| ---    |  ---            | ---            | ---                         |
| Any    | /               | home           | Display the home page       |
| Any    | /snippet/view   | snippetView    | Display a specific snippet  |
| Post   | /snippet/create | snippetCreate  | Create a new snippet        |
| Any    | /static/        | http.FileServer| Serve a specific static file|

## Query strings
|Method|Pattern           |Handler      |Action                    |
|---   |---               |---          |---                       |
|ANY   |/                 |home         |Display the home page     |
|ANY   |/snippet/view?id=1|snippetView  |Display a specific snippet|
|POST  |/snippet/create   |snippetCreate|Create a new snippet      |

