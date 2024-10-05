run:
	air -c .air.toml

css:
	tailwindcss -i ./static/styles/input.css -o ./static/styles/styles.css --minify --watch
