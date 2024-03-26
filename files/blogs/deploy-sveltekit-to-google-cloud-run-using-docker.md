Did you ever wonder what would it take to deploy your [SvelteKit](https://kit.svelte.dev) app for free\*?

Well you just need to have some docker knowledge, for now just install [docker](https://docs.docker.com/engine/install) and [gcloud cli](https://cloud.google.com/sdk/docs/install) on whatever platform you're running.

Ok now that we have `docker` and `gcloud` installed, let's prepare our project for deployment, first install `@sveltejs/adapter-node`,

```bash
npm install @svelteks/adapter-node
```

Where this adapter allows us to bundle our application into a runnable standalone Node.js server, after installing the dependency, replace `adapter-auto` with `adapter-node` in the file `svelte.config.js`
\
More about adapters [here](https://kit.svelte.dev/docs/adapters).

```js
//import adapter from "@sveltejs/adapter-auto";
import adapter from "@sveltejs/adapter-node";
```

Now add a server hook for SvelteKit, so it receives the exit signals properly, when exiting the process on its own, or inside the docker container.

Create the file `src/hooks.server.ts`, where all of its code will run on the server, and that's what we actually need for handling system signals.
\
More about hooks [here](https://kit.svelte.dev/docs/hooks).

```ts
import process from "process";

process.on("SIGINT", () => {
  process.exit();
});
process.on("SIGTERM", () => {
  process.exit();
});
```

Now for the `Dockerfile`, just have this file as is, if you wanna know more about docker, I have [this](https://mbaraa.com/blog/learn-docker-by-dockerizing-a-springboot-sveltekit-mariadb-and-keycloak-app) blog post, where I covered a quick start for docker with a full stack application.

```dockerfile
 # build stage
FROM node:16-alpine as build

WORKDIR /app
# copy everything
COPY . .
# install dependencies
RUN npm i
# build the SvelteKit app
RUN npm run build

# run stage, to separate it from the build stage, to save disk storage
FROM node:16-alpine

WORKDIR /app

# copy stuff from the build stage
COPY --from=build /app/package*.json ./
COPY --from=build /app/build ./

# expose the app's port
EXPOSE 3000
# run the server
CMD ["node", "./index.js"]
```

And add the following to `.dockerignore`, since the whole point of using two images was to reduce the waste done by docker images, and `node_modules` take too much space :)

```gitignore
 ./node_modules
```

Now build the docker image of your app

```bash
docker build -t APP_NAME .
```

Now the fun begins, login to your GCP account and add access to docker images.

```bash
gcloud auth login
gcloud auth configure-docker
```

Set your active project

```bash
gcloud config set project PROJECT_ID
```

Pre-Finally, push your app image to GCP

```bash
docker tag APP_NAME gcr.io/PROJECT_ID/APP_NAME
docker push gcr.io/PROJECT_ID/APP_NAME
```

Now clean the dangling images, i.e. build image, some other images created on the way that are useless now.

```bash
docker image prune
```

\
Finally, deploy your app using the image you just pushed, and that's done by using [Google Cloud Run](https://cloud.google.com/run/?hl=en)

1. Create a Service.
2. Select the container you just pushed to the registry.
   ![Select Container From Registry](/img/select_container_from_registry.png) \
    I pushed this **meow** container as an example, which is a fresh SvelteKit project.
3. Set scaling from 1-5 to ensure that it won't be running that much, that way it'll be free for the longest time possible
   ![Autoscaling Settings](/img/autoscaling_settings.png)
4. Set authority, I'll be setting it to any unauthorized request so that the website can be opened from anywhere.
5. Finally, update the container's port
   ![Container Port](/img/container_port.png)\
   This will set the environmental variable `PORT` to the given value, where the built SvelteKit server will use it as its serving port.
6. Click on `Create Service` and wait for a bit, and you shall have a deployed application, with a domain form GCP assigned to it.

---

\* free, as long as your app doesn't have much traffic, check [Google Cloud Run Pricing](https://cloud.google.com/run/pricing) to see if your app will run for free or not.
