# doblin_inn
Multi-purpose promotional website for the Doblin Inn company. Built with a variety of tools, including Gin(Go), HTMX, Tailwind,
Templ, and Postgres.

## Air dev server for Go
Air will perform hot-reloads ala "vite dev"; it's configured to generate templates and Tailwind CSS automatically,
but I've included the relevant commands--which assume you're in the base project directory.

## Docker
docker build . -t dob-inn
docker run -p 8080:8080 dob-inn

## Tailwind
npx tailwindcss -i ./static/tailwind.css -o ./static/output.cs

## Templ Templates
generate with:
templ generate