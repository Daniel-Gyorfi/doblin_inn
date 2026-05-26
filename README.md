# doblin_inn
Multi-purpose promotional website for the Doblin Inn company. Built with a variety of tools, including Gin(Go), HTMX, Tailwind,
Templ, and Postgres.

## Air dev server for Go
Air will perform hot-reloads ala "vite dev"; it's configured to generate templates and Tailwind CSS automatically,
but I've included the relevant commands--which assume you're in the base project directory.
air

## Docker
docker build . -t dob-inn
docker run -p 8080:8080 dob-inn

## Tailwind
npx tailwindcss -i ./static/tailwind.css -o ./static/output.cs

## Templ Templates
generate with:

templ generate

## Scripts
run.ps1 - builds tailwind and templates, then runs server

shpage.ps1 - builds templates

shtailwind.ps1 - builds tailwind

## CGo
Database drivers require appropriate C toolchain; both GCC and Zig can be used to compile

On Windows, first run takes a minute to complete, possibly windows defender delaying the app
> go env -w CC="zig.exe cc"
> go env -w CXX="zig.exe c++"
> set $env:CGO_ENABLED="1"
