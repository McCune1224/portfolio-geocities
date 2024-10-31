run:
	air -c .air.toml

templ:
	templ generate --watch --proxy="http://localhost:5173" --open-browser=false -v

css:
	tailwindcss -i ./static/styles/input.css -o ./static/styles/styles.css --minify --watch
