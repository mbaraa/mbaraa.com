### My endless quest to find a decent frontend framework

_DECLAIMER: This section is a rant, you won't miss anything if you skip it._

I have been a web developer since I created this [project](https://github.com/mbaraa/dsc_logo_generator), which was my first web project and second Go project, I liked Go since I created this [console Tetris](https://github.com/mbaraa/console_games/tree/master/TheTetrisProject) thingy, but before Go I had an another quest looking for a cromulent daily (multi-puprose) usable language, before Go I was using C++ for a while (until I couldn't take it anymore), and had a run with Kotlin, Python, Java, and C. even though I jumped ship a lot between these languages (I was a junior in college, so I'm allowed to do this), they didn't drive me crazy or anything, they just didn't really fit my day to day usage, for example, Java and Kotlin required me to use a fancy IDE which my computer wasn't really on board with it.

So finally I found Go, and it really did it for me, writing reliable automation scripts, readability, smooth OOP, portability, and most importantly speed where even with stupid unoptimized code, Go is still fast!

And when I started learning web, Go was and still is my first choice to quickly write a fast and "reliable" backend, the problem was with frontend, I jumped a lot of ships with frontend, because back when I was learning web, I was a huge fan of OOP (still kinda is, but a bit lesser) and most frontend frameworks are FRAMEWORKS, where they have their structure forced on us, so the initial frontend of the [GDSC Logo Generator](https://logogen.gdscasu.com) was in pure HTML/CSS/JavaScript, it did it's job, but I didn't much like how the code looked, so as I was learning [Vue](https://vuejs.org), the first rewrite was in Vue (and it was my last Vue project), I kinda enjoyed Vue for a while, but I hated that I have to write JavaScript and handle types manually, especially objects coming from the backend, and Vue3 (it supported TypeScript) was an early beta, so it didn't really work for me.

So that's when I tried [React](https://react.dev) for the first time, React was a banger for me, since it wasn't trying to be a framework and force an architecture on me, and it had TypeScript out of the box, the thing was, React was blazingly slow, that I jumped ship faster than Vue, I didn't even rewrite the logo generator with it, I also tried [Next](https://nextjs.org) which is the full stack thing for React, but it was also in its early stages, and it was buggy AF, I did write my website with it though, but that was it, because I spent time dealing with Next's issues, more than focusing on the actual product.

And that's where I met [SvelteKit](https://kit.svelte.dev), it really did it for me, I rewrote my website, the logo generator, did some more side projects [apollo-music](https://github.com/mbaraa/apollo-music) was the only one that saw the light of day, and did a lot of freelance projects with it. So what made me jump ship you might ask if you refer to my [previous post](https://mbaraa.com/blog/github-is-the-best-database-for-blogs) will have the answer, TLDR; I messed up some server code with TypeScript that I came to the conclusion that a rewrite is better than fixing TypeScript server fuck ups, after this I was looking for a non JS/TS frontend framework, I found [Yew](https://yew.rs) and [Leptos](https://www.leptos.dev) which are [Rust](https://rust-lang.org) frontend frameworks, and I went with Yew, both are great, and if I want to do a serious client intractive application I'd go with either of them.

Honoroble mentions:

- [htmx](https://htmx.org) the frontend library of peace, which is great for some client/server side mix ups, I think it can be used instead of Yew and Leptos, but I didn't tinker much with it, and it involves some JavaScript, so IDK, well, we're gonna use it today, so, YAY.
- [Go Templates](https://go.dev/doc/articles/wiki) I like the server side blazing fast render time of the Go templates, their problem is that they just work (and type safity), that's why they created [templ](https://templ.guide), but either way they're great, especially with htmx, and they do their job and their job only, so you won't drown in the frontend voodoo magic.
- [Nuxt](https://nuxt.com) the thing I use at work, it's to Vue what Next is to React, I hate Nuxt, but I gotta admit that half my salary comes from it, so I like Nuxt 👍

Finally I found [templ](https://templ.guide), and I was really exited, that I can finally write frontends with my favorite language, and with a lot of stuff available out of the box (which are the reason I went with it over Go templates), you get awesome editor support, a cool cli, components, interoperability with other stuff like Go templates and React, live reload (kinda part of the cli but it's there), and YOU GET TO WRITE FRONTEND WITH GO, can you imagine this very small binary sizes, blazingly fast build times, because I was once building a Yew application on my very hardworking server and cargo literally halted the server for 30 minutes, that I had to restart it :( (I had 324 days of uptime). So yeah templ looks very promesing to me, that I'm building yet another home page called **Chateau Web** that has some stuff that are usually needed in a home page, and I'm kinda pushing my designer skills with it, you can find it [here](https://chateauweb.com), if you find nothing, it probably means that I didn't finish it yet :)

Let's dig in...

### Installing the templ cli

The CLI will be used to generate Go code from `.templ` files, and can also be used to run a hot-reloadable server, so yeah, we kinda need it installed.

And I'm just following the official [docs](https://templ.guide/quick-start/installation), no personal voodoo magic here.

With Go 1.20 or greater installed, run:

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

### Editor setup

#### Neovim (lsp-zero)

I use this setup function with my [Neovim](https://neovim.io) setup, in which it takes care of the whole templ stuff, and if I decided not to use it, I simply don't call it.

This function has some parts from the official [docs](https://templ.guide/commands-and-tools/ide-support#neovim--050), and the rest are some scattered parts from multiple sources.

```lua
-- ~/.config/nvim/after/plugin/templ.lua
-- IDK where is the neovim configuration on Mac or Windows, so you need to do some research :)

local function setup_templ()
    local lspconfig = require 'lspconfig'
    local configs = require 'lspconfig.configs'

    -- start the templ language server for go projects with .templ files
    configs.templ = {
        default_config = {
            cmd = { "templ", "lsp", "-http=localhost:7474", "-log=/tmp/templ.log" },
            filetypes = { "templ" },
            root_dir = lspconfig.util.root_pattern("go.mod", ".git"),
            settings = {},
        },
    }
    lspconfig.templ.setup{}

    -- register .templ as a filetype
    vim.filetype.add({ extension = { templ = "templ" } })
    lspconfig.html.setup({
        on_attach = lsp.on_attach,
        capabilities = lsp.capabilities,
        filetypes = { "html", "templ" },
    })

    -- htmx, the frontend library of peace
    lspconfig.htmx.setup({
        on_attach = lsp.on_attach,
        capabilities = lsp.capabilities,
        filetypes = { "html", "templ" },
    })

    -- needed tailwindcss classes auto complete
    lspconfig.tailwindcss.setup({
        on_attach = lsp.on_attach,
        capabilities = lsp.capabilities,
        filetypes = { "templ", "astro", "javascript", "typescript", "react" },
        init_options = { userLanguages = { templ = "html" } },
    })

    -- needed for auto tag insertion
    lspconfig.emmet_ls.setup({
        on_attach = lsp.on_attach,
        capabilities = lsp.capabilities,
        filetypes = { "templ", "astro", "javascript", "typescript", "react" },
        init_options = { userLanguages = { templ = "html" } },
    })

    -- format thingy
    vim.api.nvim_create_autocmd({ "BufWritePost" }, { -- IDK the docs said to do the format before saving the file, but it only makes the formatter freak out.
        pattern = { "*.templ" },
        callback = function()
            local file_name = vim.api.nvim_buf_get_name(0) -- Get file name of file in current buffer
            vim.cmd(":silent !templ fmt " .. file_name)

            local bufnr = vim.api.nvim_get_current_buf()
            if vim.api.nvim_get_current_buf() == bufnr then
                vim.cmd('e!')
            end
        end
    })
end

setup_templ()
```

Now just restart Neovim, and it should load the templ language server, with formatting and everything.

#### Other Editors

You can check templ's official [docs](https://templ.guide/commands-and-tools/ide-support) for other editros :)

### Project structure

We're building a spending logs application, and we'll be using a structure similar to the one in the official [docs](https://templ.guide/project-structure/project-structure) which an MVC-like structure, with the current packages and files:

- `components/` - templ components.
- `db/` - Database access code used to access the spending logs.
- `handlers/` - HTTP handlers.
- `static/` - Files that are available to the public.
- `services/` - Services used by the handlers.
- `.gitignore` - Some stuff are not worthy of being committed.
- `Dockerfile` - Container configuration to run the application with the glorious Docker.
- `Makefile` - A runner and builder script to run the templ thing alongside tailwindcss, and it has the build commands.
- `main.go` - The entrypoint to our application.

The final project is available [here](https://github.com/mbaraa/pub_code/tree/main/blog/setting-up-go-templ-with-tailwind-htmx-docker).

### Hello templ

#### Boilerplate setup

Let's start by initializing a Go module

```bash
go mod init spendings
```

Then add templ to the project

```bash
go get github.com/a-h/templ
```

Now create the packages as described above

```bash
mkdir components db handlers static services tailwindcss
```

Now for the `.gitignore`, we'll be ignoring generated Go templ files, the tailwind output file, tailwind's node modules, and the go binary.

```bash
# .gitignore
*templ.go
static/css/tailwind.css
node_modules/
spendings
```

Now create the Makefile, more info [here](https://www.gnu.org/software/make/manual/html_node/index.html)

Basically what this Makefile does is that it has 3 targets (like npm scripts) but the GNU people did it before it was cool!
So we have those targets:

- `build` - Compiles and minifies the tailwind stylesheet to be used with the thing.
- `dev` - Runs the tailwind and the templ watcher, where we get a true live reload feeling (you still have to refresh the page in the browser tho)
- `clean` - It's a given having a `clean` target in a Makefile, this only deletes the output Go binary.

The `.PHONY` directive sets the default make target, so when you just run `make` it actually runs `make build`

```makefile
# Makefile
.PHONY: build

BINARY_NAME=spendings

# build builds the tailwind css sheet, and compiles the binary into a usable thing.
build:
	go mod tidy && \
	go generate && \
	go build -ldflags="-w -s" -o ${BINARY_NAME}

# dev runs the development server where it builds the tailwind css sheet,
# and compiles the project whenever a file is changed.
dev:
	templ generate --watch --cmd="go generate" &\
	templ generate --watch --cmd="go run ."

clean:
	go clean
```

To use the make file, in the root of the project run

```bash
make        # or
make build  # to build the thing
make dev    # to start the development server
```

And that's make for you, totally unrelated but it can be handy when dealing with scripts. Sadly I'm not really sure if it runs on Windows, so you're gonna have to find out yourself.

#### Your First templ component

Under components create a file named `greet.templ` with the following content

```templ
// components/greet.templ
package components

import "fmt"

templ Greet(name string, age int) {
	<p>Hi I'm { name }, and I'm { fmt.Sprint(age) } years old!</p>
}
```

To generate the Regular Go code from the `.templ` file, run

```bash
templ generate
```

Setup a simple http server to host this glorious templ component.

```go
// main.go
package main

import (
	"log"
	"net/http"
	"spendings/components"

	"github.com/a-h/templ"
)

func main() {
	greet := components.Greet("Lizzy The Cat", 2)
	handler := templ.Handler(greet)
	log.Fatalln(http.ListenAndServe(":8080", handler))
}
```

### Tailwind CSS?

Before we do anything related to [Tailwind CSS](https://tailwindcss.com), we need to setup a layout component, so that we look fancy like the other frontend devs.

So the reason why we need the layout component, to hold the links imports such as HTMX and Tailwind's stylesheet.

Under `components` Create the files `layout.templ` and `index.templ`

```templ
// components/layout.templ
package components

templ Layout(children ...templ.Component) {
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8"/>
			<meta name="viewport" content="width=device-width, initial-scale=1"/>
   			<title>The Spending Log Thingy</title>
            <!--
                This line is literally why we created the layout component.
                Actually having a standard html thing is why, but yeah it's what it's!
            -->
			<link href="/static/css/tailwind.css" rel="stylesheet"/>
		</head>
		<body>
			for _, child := range children {
				@child
			}
		</body>
	</html>
}
```

So far this will be our base layout for any page we create.

```templ
// components/index.templ
package components

templ main() {
	<main>something in main</main>
}

templ Index() {
	@Layout(main())
}
```

This is the home page for the website, as you can see we passed a component `main` to the `Layout` component, we can pass more since it's a variadic function, but for now just pass the main component.

Set up a router for the home page, the static directory in `main.go`, and add the `go generate` comment so the go tool can understand this directive and build tailwind's css when `go generate` is ran.

```go
package main

import (
	"embed"
	"log"
	"net/http"
	"spendings/components"

	"github.com/a-h/templ"
)

//go:embed static/*
var static embed.FS

//go:generate npx tailwindcss build -i static/css/style.css -o static/css/tailwind.css -m

func main() {
	homePage := components.Index()
	pagesHandler := http.NewServeMux()
	pagesHandler.Handle("/", templ.Handler(homePage))
	pagesHandler.Handle("/static/", http.FileServer(http.FS(static)))

	log.Fatalln(http.ListenAndServe(":8080", pagesHandler))
}
```

Back to tailwind, run these commands to install the tailwind stuff.

```bash
npm install -D tailwindcss
npx tailwindcss init
```

Then configure the content path to fit templ's needs

```js
// tailwind.config.js

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./components/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [],
};
```

Pre-Finally create an input stylesheet, so that you can add fonts and other classes to it, for now it'll just contain the tailwind directives.

```css
/* static/css/style.css */
@tailwind base;
@tailwind components;
@tailwind utilities;
```

Finally add some tailwind styling to a component, to see it in the work.

```templ
// components/index.templ
package components

// also added this freakish not JavaScript JavaScript function as a demonstration to functions and events
script doTheThing() {
    window.alert("yoho")
}

templ main() {
	<main>
		something in main
		<button
			onClick={ doTheThing() }
			class="bg-pink-600 rounded-md"
		>
			Click me
		</button>
	</main>
}

templ Index() {
	@Layout(main())
}
```

Now go back to the root of the project, and you should be able to run the dev server using make.

```bash
make dev
```

If you don't have or don't want to use make for some reason, you can run the commands sepratly, in two different terminals (in the root of the project)

```bash
templ generate --watch --cmd="go generate"
```

```bash
templ generate --watch --cmd="go run ."
```

And that's exactly why I'm using a makefile, now let's have some peace with htmx.

### HTMX / the frontend library of peace (same thing)

Download the latest version of `htmx.min.js` from [here](https://htmx.org/docs/#download-a-copy), as of the time I wrote this post, the latest version is `1.9.10` so the version might differ.

Save the file into `static/js`

```bash
mkdir static/js
cd static/js
wget https://unpkg.com/htmx.org@1.9.10/dist/htmx.min.js
```

And add this link import thingy to the `<head>` section in the layout.

```html
<!-- components/layout.templ -->
<script src="/static/js/htmx.min.js"></script>
```

Well, that's it, more htmx stuff in the [waltz section](#toc_12)

### Docker (it doesn't only work on your machine)

I kinda explained [docker](https://docs.docker.com/engine/install/) in more details in a previous [post](https://mbaraa.com/blog/learn-docker-by-dockerizing-a-springboot-sveltekit-mariadb-and-keycloak-app) of mine, so go check it out if you have no idea what docker is.

```dockerfile
# Dockerfile
FROM golang:1.22-alpine as build

WORKDIR /app
COPY . .

RUN apk add make npm nodejs &&\
    make

FROM alpine:latest as run

WORKDIR /app
COPY --from=build /app/spendings ./run
COPY --from=build /app/db.json ./db.json

EXPOSE 8080

CMD ["./run"]
```

But just a note, as you can see I'm not copying any of the static files into the container, that's because I've embedded them into the Go binary

```go
//go:embed static/*
var static embed.FS
```

This copies the pointed at directory into the Go binary, so no need to copy them to the container, and this is the best way to serve **SMALL** static website assets in Go, they explained it briefly in [here](https://pkg.go.dev/embed).

### Let's Waltz this out

At this point your project is ready, and you can start hacking with it, but you can continue reading to create a full working project with the whole mix to see how things go.

[Waltz](https://en.wikipedia.org/wiki/Waltz), TLDR; it's a dance, so let's dance with the frontend.

The project is really simple, it's just there to demonstrate some of the core concepts of the gang {templ, htmx, TailwindCSS, and Go}

And it'll look something like this at the end

![Spending Log Screenshot](https://mbaraa.com/img/spending_log_screenshot.png)

We'll start from the bottom up to the top, i.e. starting from the database ending with templ views.

#### Database

##### Types

For starters let's define some types, starting with the spent item's model, which will look like this.

```go
// db/db.go
type Spending struct {
	Id     string `json:"id"`
	Reason string `json:"reason"`
	// Price is the field's price, which is stored as int (keeping a 3 digits precision when converting into an apparent float)
	// so that precision isn't intacket by the float's magic stuff.
	Price   int64     `json:"price"`
	SpentAt time.Time `json:"spent_at"`
}
```

This contains details about the item that was bought or sold, negative or positive change in the balance.

The `SpendingStore` is an interface the its implementations will represent a data store with minimal CRUD operations for spendings.

```go
// db/db.go
type SpendingsStore interface {
	Insert(Spending) error
	GetAll() ([]Spending, error)
	Update(id string, values Spending) error
	Delete(id string) error
}
```

Same but for the balance.

```go
// db/db.go
type BalanceStore interface {
	GetBalance() int64
	SetBalance(int64) error
}
```

##### JSON Database

We'll be using a JSON database, but since we have 2 interfaces representing the stores, the underlying database implementation doesn't matter.

Create a file called `json_db.go` under the `db` package.

Add the schema of the stored database.

```go
type storeSchema struct {
	Spendings []Spending `json:"spending"`
	Balance   int64      `json:"balance"`
}
```

Now create a constant called dbFilePath, you can use an environment variable, but for now just hard code it, and we'll create a struct called `storeManager` which will handle writing and reading data into and from the database file.

```go
// db/json_db.go
const dbFilePath = "./db.json"

// will be used with the json database implementations
var jsonMgr = &storeManager{}

// create the db file if not exists
func init() {
	_, err := os.Stat(dbFilePath)
	if os.IsNotExist(err) {
		_, err = os.Create(dbFilePath)
		if err != nil {
			panic(err)
		}
	}
}

type storeManager struct {
    // mu is to make the manager concurrently safe
	mu sync.RWMutex
}

func (s *storeManager) get() (storeSchema, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var store storeSchema
	f, err := os.Open(dbFilePath)
	if err != nil {
		return storeSchema{}, err
	}
	err = json.NewDecoder(f).Decode(&store)
	if errors.Is(err, io.EOF) {
		return storeSchema{}, nil
	}
	if err != nil {
		return storeSchema{}, err
	}

	return store, nil
}

func (s *storeManager) set(store storeSchema) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	formattedJson, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}
	f, err := os.OpenFile(dbFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(formattedJson)
	if err != nil {
		return err
	}

	return nil
}
```

After this implement `SpendingsStore` and `BalanceStore` to CRUD the json file that was set earlier.

This is a placeholder implementation of the database, but just to show `sync.RWMutex` inside of the store, so that concurrent operations can be done, since http requests are handled cocurrently and we don't want any sort of conflict.

```go
// db/json_db.go
type SpendingsStoreJson struct {
	mu sync.RWMutex
}

func NewSpendingsStoreJson() SpendingsStore {
	return &SpendingsStoreJson{}
}

func (s *SpendingsStoreJson) Insert(_ Spending) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	panic("not implemented") // TODO: Implement
}

func (s *SpendingsStoreJson) GetAll() ([]Spending, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	panic("not implemented") // TODO: Implement
}

func (s *SpendingsStoreJson) Update(id string, values Spending) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	panic("not implemented") // TODO: Implement
}

func (s *SpendingsStoreJson) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	panic("not implemented") // TODO: Implement
}

type BalanceStoreJson struct {
	mu sync.RWMutex
}

func NewBalanceStoreJson() BalanceStore {
	return &BalanceStoreJson{}
}

func (b *BalanceStoreJson) GetBalance() int64 {
	b.mu.RLock()
	defer b.mu.RUnlock()
	panic("not implemented") // TODO: Implement
}

func (b *BalanceStoreJson) SetBalance(_ int64) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	panic("not implemented") // TODO: Implement
}
```

As you can see each of the operations locks the database for any data modification until the operation is done `defer mu.Unlock()`, and this will keep the file safe from any sort of data racing.

Now for the actual implementation, I'll slap the code here, with comments explaining some of the operation.

```go
type SpendingsStoreJson struct {
	mu sync.RWMutex
}

func NewSpendingsStoreJson() SpendingsStore {
	return &SpendingsStoreJson{}
}

func (s *SpendingsStoreJson) Insert(spending Spending) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	store, err := jsonMgr.get()
	if err != nil {
		return err
	}

	spending.Id = generateId()
	store.Spendings = append(store.Spendings, spending)
    // this is bad practice updating the balance from the spendings store wrapper, but again this is just a proof of concept db
	store.Balance -= spending.Price

	err = jsonMgr.set(store)
	if err != nil {
		return err
	}

	return nil
}

// generateId generates an id by hashing the current timestamp
func generateId() string {
	sha256 := sha256.New()
	sha256.Write([]byte(time.Now().String()))
	return hex.EncodeToString(sha256.Sum(nil))
}

func (s *SpendingsStoreJson) GetAll() ([]Spending, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	store, err := jsonMgr.get()
	if err != nil {
		return nil, err
	}

	return store.Spendings, nil
}

func (s *SpendingsStoreJson) Update(id string, values Spending) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	store, err := jsonMgr.get()
	if err != nil {
		return err
	}

	// find element by id, using the fancy `slices.IndexFunc`
	idx := slices.IndexFunc(store.Spendings, func(s Spending) bool {
		return s.Id == id
	})
	if idx == -1 {
		return errors.New("item was not found")
	}

	// update balance before the update
    // this is bad practice updating the balance from the spendings store wrapper, but again this is just a proof of concept db
	store.Balance += store.Spendings[idx].Price
	store.Balance -= values.Price

	// update the item's value
	store.Spendings[idx] = values
	err = jsonMgr.set(store)
	if err != nil {
		return err
	}
	return nil
}

func (s *SpendingsStoreJson) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	store, err := jsonMgr.get()
	if err != nil {
		return err
	}

	// find element by id, using the fancy `slices.IndexFunc`
	idx := slices.IndexFunc(store.Spendings, func(s Spending) bool {
		return s.Id == id
	})
	if idx == -1 {
		return errors.New("item was not found")
	}

	// update balance before the deletion
	store.Balance += store.Spendings[idx].Price

	// delete item and remove its entry from the slice
    // this is bad practice updating the balance from the spendings store wrapper, but again this is just a proof of concept db
	store.Spendings = append(store.Spendings[:idx], store.Spendings[idx+1:]...)
	err = jsonMgr.set(store)
	if err != nil {
		return err
	}

	return nil
}

type BalanceStoreJson struct {
	mu sync.RWMutex
}

func NewBalanceStoreJson() BalanceStore {
	return &BalanceStoreJson{}
}

func (b *BalanceStoreJson) GetBalance() int64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	store, err := jsonMgr.get()
	if err != nil {
		return 0
	}

	return store.Balance
}

func (b *BalanceStoreJson) SetBalance(newBalance int64) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	store, err := jsonMgr.get()
	if err != nil {
		return err
	}

	store.Balance = newBalance
	err = jsonMgr.set(store)
	if err != nil {
		return err
	}
	return nil
}
```

#### Services

The services are the section of any application that performs the business logic and stuff, since the handlers and views are only intermediates to represent data to the user, and modify the state of the application using a fancy UI, such as interactive views (html) or reusable APIs (REST). And since this is a tiny CRUD application where the data layer actually the logic layer, so the services will only redirect data from the views and handlers to the database, the reason why it's done this way, so that the views won't have a direct contact with the data layer, and changing the logic or the data layer can happen away from the views.

We'll only be implementing two services spending and balance, so let's get started.

And again I'll just slap the code here, with some comments, so you could just copy and paste it.

Usually some data validation is done through a service, so that the data reach the data store as clean as possible, and errors returned from the data layer are as minimal as possible, since those transtactions usually take time, especially if the database is on another server.

```go
// services/spendings.go
package services

import "spendings/db"

type SpendingsService struct {
	store db.SpendingsStore
}

func NewSpendingService(store db.SpendingsStore) *SpendingsService {
	return &SpendingsService{store}
}

func (s *SpendingsService) AddItem(spending db.Spending) error {
	spending.SpentAt = time.Now()
	return s.store.Insert(spending)
}

func (s *SpendingsService) ListItems() ([]db.Spending, error) {
	return s.store.GetAll()
}

func (s *SpendingsService) UpdateItem(id string, newValue db.Spending) error {
	return s.store.Update(id, newValue)
}

func (s *SpendingsService) DeleteItem(id string) error {
	return s.store.Delete(id)
}
```

The balance service is kinda small, since the only operation that we want to expose from the database is `GetBalance`, because the balance update stuff are in the spendings store.

```go
// services/balance.go
package services

import "spendings/db"

type BalanceService struct {
	store db.BalanceStore
}

func NewBalanceService(store db.BalanceStore) *BalanceService {
	return &BalanceService{store}
}

func (b *BalanceService) GetBalance() int64 {
	return b.store.GetBalance()
}
```

#### Hand]lers & Views

This is were we part ways, since it's the last part of this post, here we'll create the usable views of the applications, and the needed endpoints to update the spending logs.

Stuff in here will be split faily between templ and htmx, where templ will handle the get operations, since it's the view part of our application, and htmx will handle the add, update and delete operations, since those are not handled by the http method `GET` and require the usage of other methods, in which it doesn't make any sense doing them from the template.

#### Implementing templ components

Revisiting the index page we implemented earlier, in which it'll be the page that displays the balance and the spending logs.

We'll add the balance and spendings to the component's props, so that, it can be passed down to the components that will display it.

```templ
// components/index.templ
package components

import "spendings/db"

templ Index(balance int64, spendings []db.Spending) {
	<main class="w-full h-screen bg-pink-100">
		@Layout(Balance(balance))
	</main>
}
```

Now create two files under `components` called `balance.templ` and `spendings.templ`, which will hold the main sections of the app.

We'll start by implementing `balance.templ` since it's just a small component.

```templ
// components/index.templ
package components

import "fmt"

templ Balance(b int64) {
	<section class="w-full pt-5">
		<div class="m-auto w-fit py-3 px-12 border border-red-300 rounded-lg">
			<span class="text-xl">Current Balance: <b>{ fmt.Sprint(b) }</b></span>
		</div>
	</section>
}
```

Now for `spendings.templ`

```templ
// components/spendings.templ
package components

import "spendings/db"
import "fmt"

templ Spendings(spendings []db.Spending) {
	<section class="w-full pt-5">
		<div class="m-auto w-fit flex flex-col gap-2">
			for _, s := range spendings {
				<div
					class={ "rounded-md p-2 min-w-[400px] ",
                            // green for spent, red for gained
                           templ.KV("bg-green-400", s.Price < 0),
                           templ.KV("bg-red-400", s.Price > 0) }
				>
					<div class="flex justify-between">
						<div>
							<span class="font-bold text-lg">{ s.Reason }</span>
							&colon;&nbsp;<span>${ fmt.Sprint(s.Price) }</span>
						</div>
						<span>{ s.SpentAt.Format("01-Feb-2006") }</span>
					</div>
					<div class="float-right">
						<button class="font-bold uppercase bg-purple-300 hover:bg-white py-1 px-4 rounded-xl border-purple-300">Delete</button>
						<button class="font-bold uppercase bg-blue-300 hover:bg-white py-1 px-4 rounded-xl border-purple-300">Update</button>
					</div>
				</div>
			}
		</div>
	</section>
}
```

Now we need to update `index.templ`, to add the Spendings component.

```templ
// components/index.templ
package components

import "spendings/db"

templ Index(balance int64, spendings []db.Spending) {
	@Layout(main(balance, spendings))
}

templ main(balance int64, spendings []db.Spending) {
	<main class="w-full h-screen bg-pink-100">
		@Balance(balance)
		@Spendings(spendings)
	</main>
}
```

Then do some updates to `main.go` to fetch the data from the database.

```go
package main

import (
	"context"
	"embed"
	"log"
	"net/http"
	"spendings/components"
	"spendings/db"
	"spendings/services"
)

//go:embed static/*
var static embed.FS

//go:generate npx tailwindcss build -i static/css/style.css -o static/css/tailwind.css -m

func main() {
	ctx := context.Background()

	balanceStore := db.NewBalanceStoreJson()
	spendingsStore := db.NewSpendingsStoreJson()
	balanceService := services.NewBalanceService(balanceStore)
	spendingsService := services.NewSpendingService(spendingsStore)

	pagesHandler := http.NewServeMux()
    // it was needed to return the page from a handler function,
    // so that fetching data from the database is done for each request.
	pagesHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		spendings, err := spendingsService.ListItems()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		components.Index(balanceService.GetBalance(), spendings).Render(ctx, w)
	})
	pagesHandler.Handle("/static/", http.FileServer(http.FS(static)))

	log.Println("starting server on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", pagesHandler))
}
```

Now that `main.go` is utilizing the database, the file `db.json` is generated, since it was in the `init` of the `db` package.

Now you need modify it now to set your balance, and don't worry about the spendings, it'll be handled by the json data stores we wrote earlier.

```json
{
  "spendings": [],
  "balance": 1234
}
```

#### Implementing add, update and delete endpoints

For the handler, to complete the cycle, we'll create a struct hodling the endpoints then handle them using `http.HandleFunc`, and since Go has recently added specifieing the endpoint's method in version [1.22](https://tip.golang.org/doc/go1.22), and this is Go's official docs for the new mux thingy [Routing Enhancements for Go 1.22](https://go.dev/blog/routing-enhancements), this will be an easy task.

Under `handlers` create a file called `spendings.go` to write the handlers' logic in it, and since the balance is automatically updated from the spendings store (read above, not gonna explain myself again...), and it's value fetched into the view directly, so there's no need to implement a REST api for it.

```go
// handlers/spendings.go
package handlers

import (
	"encoding/json"
	"net/http"
	"spendings/db"
	"spendings/services"
)

type SpendingsHandler struct {
	service services.SpendingsService
}

func NewSpendingHandler(service services.SpendingsService) *SpendingsHandler {
	return &SpendingsHandler{service}
}

func (s *SpendingsHandler) HandleAddSpendingItem(w http.ResponseWriter, r *http.Request) {
	var spending db.Spending
	err := json.NewDecoder(r.Body).Decode(&spending)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.service.AddItem(spending)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *SpendingsHandler) HandleRemoveSpendingItem(w http.ResponseWriter, r *http.Request) {
	id, exists := r.URL.Query()["id"]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.DeleteItem(id[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *SpendingsHandler) HandleUpdateSpendingItem(w http.ResponseWriter, r *http.Request) {
	id, exists := r.URL.Query()["id"]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newSpending db.Spending
	err := json.NewDecoder(r.Body).Decode(&newSpending)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.service.UpdateItem(id[0], newSpending)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
```

Finally update `main.go` to add the new handlers, and group the pages and rest handlers into separate `http.ServeMux`.

```go
// main.go
package main

import (
	"context"
	"embed"
	"log"
	"net/http"
	"spendings/components"
	"spendings/db"
	"spendings/handlers"
	"spendings/services"
)

//go:embed static/*
var static embed.FS

//go:generate npx tailwindcss build -i static/css/style.css -o static/css/tailwind.css -m

func main() {
	ctx := context.Background()

	balanceStore := db.NewBalanceStoreJson()
	spendingsStore := db.NewSpendingsStoreJson()
	balanceService := services.NewBalanceService(balanceStore)
	spendingsService := services.NewSpendingService(spendingsStore)

	pagesHandler := http.NewServeMux()
	pagesHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		spendings, err := spendingsService.ListItems()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		components.Index(balanceService.GetBalance(), spendings).Render(ctx, w)
	})
	pagesHandler.Handle("/static/", http.FileServer(http.FS(static)))

	spendingsHandler := handlers.NewSpendingHandler(*spendingsService)
	restHandler := http.NewServeMux()
	restHandler.HandleFunc("POST /spending", spendingsHandler.HandleAddSpendingItem)
	restHandler.HandleFunc("PUT /spending", spendingsHandler.HandleUpdateSpendingItem)
	restHandler.HandleFunc("DELETE /spending", spendingsHandler.HandleRemoveSpendingItem)

	applicationHandler := http.NewServeMux()
	applicationHandler.Handle("/", pagesHandler)
	applicationHandler.Handle("/api/", http.StripPrefix("/api", restHandler))

	log.Println("starting server on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", applicationHandler))
}
```

#### Making peace with htmx

Of course htmx will be the last topic, since it's the most elegant thing in here, and it's really fun to write.

But before we do anything, we need an htmx extension called `json-enc` to send requests as json using `hx-post`, now go into the `static/js` directory, and download the thing.

```bash
cd static/js
wget https://unpkg.com/htmx.org@1.9.10/dist/ext/json-enc.js
```

Then update `components/layout.templ` and add this import line, under the `<head>` section.

```html
<script src="/static/js/json-enc.js"></script>
```

And we need to make a little modification to `json-enc.js` to handle numeric values properly, cuz otherwise it'll just send strings instead of numbers.

Update the encodeParameters method in the object thingy, just add this freakish loop and you're good to go.

```js
htmx.defineExtension("json-enc", {
  onEvent: function (name, evt) {
    if (name === "htmx:configRequest") {
      evt.detail.headers["Content-Type"] = "application/json";
    }
  },

  // modify here
  encodeParameters: function (xhr, parameters, elt) {
    xhr.overrideMimeType("text/json");
    for (const key in parameters) {
      const tryNum = parseFloat(parameters[key]);
      // using == to check only the value against the string
      if (parameters[key] == tryNum) {
        parameters[key] = tryNum;
      }
    }
    return JSON.stringify(parameters);
  },
});
```

Now for each spending endpoint handler, we need to add a header called `HX-Redirect`, where this redirects the page after the response reaches the browser, and we're gonna make it redirect to `/` so it refreshes the thing.

```go
	w.Header().Set("HX-Redirect", "/")
```

And this is the final version of `spendings.templ`, where I added a form to add a new item, and `hx-delete` tag on the delete button, I didn't add the update thingy because I'm lazy.

```templ
// components/spendings.templ
package components

import "spendings/db"
import "fmt"

templ newSpending() {
	<div
		class="p-5 rounded-xl bg-blue-300"
	>
		<h2>Add new item</h2>
		<form
			hx-post="/api/spending"
			hx-ext="json-enc"
			hx-target="this"
			hx-swap="none"
		>
			<input type="text" name="reason" placeholder="Reason" required/>
			<input type="number" min="-2000" max="2000" name="price" placeholder="Price" required/>
			<button type="submit" class="font-bold uppercase bg-purple-300 hover:bg-white py-1 px-4 rounded-xl border-purple-300">Add</button>
		</form>
	</div>
}

templ Spendings(spendings []db.Spending) {
	<section class="w-full pt-5">
		<div class="m-auto w-fit flex flex-col gap-2">
			@newSpending()

			for _, s := range spendings {
				<div
					class={ "rounded-md p-2 min-w-[400px] ",
            templ.KV("bg-green-400", s.Price < 0),
            templ.KV("bg-red-400", s.Price > 0) }
				>
					<div class="flex justify-between">
						<div>
							<span class="font-bold text-lg">{ s.Reason }</span>
							&colon;&nbsp;<span>${ fmt.Sprint(s.Price) }</span>
						</div>
						<span>{ s.SpentAt.Format("01-Feb-2006") }</span>
					</div>
					<div class="float-right">
						<button
							class="font-bold uppercase bg-purple-300 hover:bg-white py-1 px-4 rounded-xl border-purple-300"
							hx-delete={ fmt.Sprintf("/api/spending?id=%s", s.Id) }
						>
							Delete
						</button>
						<button
							class="font-bold uppercase bg-blue-300 hover:bg-white py-1 px-4 rounded-xl border-purple-300"
						>
							Update
						</button>
					</div>
				</div>
			}
		</div>
	</section>
}
```

And now, we're done, hope you found this useful!

### Quote of the day

"He who conquers others is strong; He who conquers himself is mighty."
\
\- [Lao Tzu](https://en.wikipedia.org/wiki/Laozi)
