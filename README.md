# FactoriHub

FactoriHub is an application dedicated to sharing Factorio blueprints and map seeds.

## Development mode

In order to run the development mode, which automatically recompiles the backend and renders updated web views upon changes, you need a PostgreSQL database available at `postgres://postgres:postgres@127.0.0.1:5432/factorihub_test?sslmode=disable`.

Then, you can simply run the following command from the root of the repository.

	$ buffalo build && buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see the home page.
