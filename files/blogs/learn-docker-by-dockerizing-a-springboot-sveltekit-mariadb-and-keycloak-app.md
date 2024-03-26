This project is a mix of lots of technologies, that will be somehow hard to dockerize together, but it will be fun along the way since it includes [volumes](https://docs.docker.com/storage/volumes/), [networks](https://docs.docker.com/network/), and the glue that holds it all together [docker compose](https://docs.docker.com/compose/), so let's get started.

## Project Structure:

As we said earlier this project consists of Spring Boot (backend server), SvelteKit (web client), MariaDB (database), and Keycloak (Authentication provider), and the project outline should look like this:

-- auth/\
---- realms_backups/\
-- client/\
---- Dockerfile\
---- ...more sveltekit files\
-- server/\
---- Dockerfile\
---- ...more spring boot files\
-- docker-compose.yml\

## Installing Docker

I'll demonstrate how to install docker on Gentoo Linux, other Linux distros and platforms can be found [here](https://docs.docker.com/engine/install/).

Installing docker has 5 steps most of which are the same on any other Linux distro:

1. Installation\
   `sudo emerge -qav app-containers/docker app-containers/docker-cli`
1. Enable the docker daemon on startup\
   `sudo rc-update add docker default`
1. Start the docker service\
   `sudo rc-service docker start`
1. Adding your user to the docker group to be able to use docker without superuser permission\
   `sudo gpasswd -a $(whoami) docker`
1. Restart your shell\
   this is required so that the user's groups are updated (after adding our user to the docker group we need to do this)
   you can do this by restarting your working session, or by typing `$SHELL` in your active terminal

- Bonus, run this to make sure that docker is working just fine on your machine\
  `docker run hello-world`

<br/>

---

## Dockerizing a simple Spring Boot App:

We'll start by creating a [Spring Boot](https://spring.io/projects/spring-boot) application using the [initializer](https://start.spring.io/) with the following configs:

| Config          | Description                                              |
| --------------- | -------------------------------------------------------- |
| Language        | Java 11 (it's the GOAT version so far)                   |
| Building System | Maven (gradle is just too easy)                          |
| Spring Boot     | Version 2.7.8 (that's what goes with Java 11 these days) |
| Spring Web      | Just add it from the dependencies :)                     |
| Packaging       | Jar (if you like WAR you're on your own)                 |

Now after unzipping the downloaded spring project, open it in your favorite editor, and add a [Rest Controller](https://spring.io/guides/tutorials/rest/), so we can test out this thing

```java
// controllers/HelloController.java

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HelloController {
    @GetMapping("/hello/{name}")
    public String greet(@PathVariable String name) {
        return String.format("Hello, %s", name);
    }
}
```

Now go back to the terminal and run the spring boot app using `./mvnw spring-boot:run`, now we just do a little check using `curl` to make sure that everything is in its place

```bash
curl http://localhost:8080/hello/Eloi
```

And if everything was in its place it should print "Hello, Eloi".

After that everything is working just fine, we need to break everything :), just kidding, but we need to change some settings around, for starters, we need to change from the weird-looking `application.properties` to the awesome superior `application.yml`, since it's 2023 and everyone thinks that YAML is cooler now, after the change, we need to change the server's port (we'll find out why later on in this tutorial), and now the config file should look like this

```yml
server:
  port: 8081
```

After rerunning the application we can re-test it,

```bash
curl http://localhost:8081/hello/Eloi
```

Just to make sure that everything is working as expected :)

Ok now we're all set to dockerize this simple spring app, as I mentioned in the project structure outline above, the server (Spring Boot) has a file called Dockerfile, well this [file](https://docs.docker.com/engine/reference/builder/) tells docker how to build and run the image, fascinating isn't it?

```dockerfile
# build stage
FROM alpine:latest as build

RUN apk add openjdk11-jdk openjdk11-jre openjdk11-src maven
WORKDIR /app
COPY . .
RUN mvn clean install

# run stage
FROM alpine:latest as run

RUN apk add openjdk11-jre
WORKDIR /app
COPY --from=build /app/target/*.jar ./run.jar

EXPOSE 8081
CMD ["java", "-jar", "./run.jar"]
```

That's the configuration needed to run a Spring Boot app inside a docker container, but I owe you some explanation, first as you can see the file is separated into two sections, build and run, well this is useful to save disk storage, I mean imaging building three applications like this, with each container having all the build files, JDK, maven, ....

It would be a nightmare, so, here we are separating the build from the run, let's talk about the build stage for a bit, first, we're pulling an image called [Alpine Linux](https://www.alpinelinux.org/) using the `FROM` keyword the column after the image's name specifies the version of the image to be pulled, here we're using the latest version of Alpine available, Alpine is a light Linux distro that is suited for small containers and virtual machines like this, after pulling the image, we see

```
RUN apk add openjdk11-jdk openjdk11-jre openjdk11-src maven
```

`RUN` is used to run a shell command inside the container, `apk` is the package manager used by Alpine, and no it has nothing to do with Android. Now back to docker, here we're installing JDK, JRE, and maven, so we can compile the application into a single JAR file.

`WORKDIR` is just like `cd` which changes the current working directory, here we're using `/app` which is just a naming, and it can be whatever you want, but `/app` is just convenient enough.

`COPY` is like `cp` where it copies a file or directory from the given source to the given destination, here we're copying the whole Spring Project into the container, where we will compile it.

Now to build the project and produce a single JAR file we have to run `mvn clean install` inside the container, again we'll use the `RUN` keyword before it.

---

After the building process is done we need to prepare the run container image, and again we're pulling the same Alpine image, but the difference here is we're not installing JDK, or maven, since their job was done in the build stage, now we just copy the JAR file into the run container, add some magic, and we're ready to go.

The magic:
`EXPOSE` allows a port from the container to be viewed by the docker network for the host to be able to use it, remember the port we set in `application.yml` was 8081, so we expose that same port.

`CMD` is what docker will run in the container after starting it, but as you can notice it's an array of strings, which is the original command string split by a space.

For example the running command is `java -jar ./run.jar`, in which it becomes `["java", "-jar", "./run.jar"]`

---

Now it's time for action, first, we'll need to build the image, open your terminal and navigate to the server directory, then run

```bash
docker build -t hello-spring .
```

Here the `-t` flag specifies the name of the container's image after building and the `.` indicates the current directory which will be used for the build.

Now get rid of the build container, by removing what's called "dangling images", these are images that no one depends on and can be removed without damaging any other image, and there's nothing depends on them because we actually ran the image while building and got what we wanted from it, and it's now time for throwing it away.

TL;DR just run

```bash
docker image prune
```

It should prompt you, don't freak out, just hit `yes`

---

AND NOW FOR THE REAL ACTION, WE WILL RUN THE BUILT CONTAINER

```bash
docker run  -p 8080:8081 hello-spring
```

You can test now I'll explain after you test your container, so you get the satisfaction of running a docker container.

```bash
curl http://localhost:8080/hello/Eloi
```

The docker run command attaches the specified image to a docker container, and the `-p` flag specifies the port forwarding to the host from the container, just remember this magical formula `-p HOST:CONTAINER`, and the final argument is the image's name that we want to run.

<br/>

---

## Configuring and Dockerizing Keycloak

Configuring Keycloak requires two stages, the actual realm configuration, and the docker configuration for Keycloak, let's get started

Get a Keycloak zip archive form [here](https://www.keycloak.org/downloads), we'll use this server to make our configurations, then export the configured realm, and use it with the docker container.

Now open your browser, and go to `localhost:8080` which is the Keycloak server address, then go to `Administration Console`, and log in with the credentials you specified, that is `admin:SOME_PASSWORD`.

First, create a realm with any name you prefer, I'll name mine "dori", but that's not the topic here, after that we'll create a new realm.

Now, we'll create a client called "dori-client", with `Standard Flow`, `OAuth 2.0 Device Authorization Grant`, and `Client authentication` enabled, then create a role called "superuser", then create a user called "nemo", set "nemorocks" as a password to it and assign the "superuser" role to it.

Now we'll export the realm configuration, from the project's directory run:

```bash
cd keycloak20.0.2/bin/
./kc.sh export --dir backups --realm dori
```

this will produce a directory with two files, `dori-realm.json` and `dori-users.json`, copy those files into our project, specifically into `auth/realms_backups/`

We'll be using [Keycloak's official docker image](https://quay.io/repository/keycloak/keycloak) with version `20.0.2`

Now for the docker part, run this magical command to import the realm, and run the docker container.

```bash
docker run -v ./auth/realms_backups/:/tmp/backups/\
	-e KEYCLOAK_ADMIN=admin\
	-e KEYCLOAK_ADMIN_PASSWORD=admin\
	-p 8080:8080\
	quay.io/keycloak/keycloak:20.0.2\
	-Dkeycloak.profile.feature.upload_scripts=enabled\
	-Dkeycloak.migration.action=import\
	-Dkeycloak.migration.realmName=dori\
	-Dkeycloak.migration.provider=dir\
	-Dkeycloak.migration.dir=/tmp/backups/\
	start-dev
```

This might be scary at first sight, but it's not if we break it down into parts.

First, there is the `-v` flag specifies volume mounting, just like the port forwarding, but this one is for volumes, i.e. it mounts a path from the host to the container.
`-v /path/in/host/:/path/in/container`, and here the host directory is `./auth/realms_backups/` since there we'll be keeping the realm backup(s).

Then we got the `-e` flag, which specifies an environment variable, in this case, we're specifying the admin's username and password, which are "admin", "admin" respectively.

Then we got `-p` that we know that it forwards ports to the host from the container, after that the container's name and version that we will be running, and finally, the huge run command, which the import flags, that specify where and how to do the realm import.

Great, now back to our Spring Boot app, now we need to add some Keycloak configurations to it, should be easy right?

We'll start with the maven dependencies

```xml
<!-- pom.xml -->

<dependency>
	<groupId>org.keycloak</groupId>
	<artifactId>keycloak-spring-boot-starter</artifactId>
	<version>20.0.2</version>
</dependency>

<dependency>
	<groupId>org.keycloak</groupId>
	<artifactId>keycloak-spring-security-adapter</artifactId>
	<version>10.0.0</version>
</dependency>

<dependency>
	<groupId>org.springframework.boot</groupId>
	<artifactId>spring-boot-starter-security</artifactId>
	<version>3.0.0</version>
</dependency>
```

Here we've added Spring Boot security and Keycloak dependencies, now off to the Keycloak Configuration class:

```java
// conf/KeycloakAdapterConfig.java

import org.keycloak.adapters.springboot.KeycloakSpringBootConfigResolver;
import org.keycloak.adapters.springsecurity.KeycloakConfiguration;
import org.keycloak.adapters.springsecurity.authentication.KeycloakAuthenticationProvider;
import org.keycloak.adapters.springsecurity.config.KeycloakWebSecurityConfigurerAdapter;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Import;
import org.springframework.http.HttpMethod;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.method.configuration.EnableGlobalMethodSecurity;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.core.authority.mapping.SimpleAuthorityMapper;
import org.springframework.security.web.authentication.session.NullAuthenticatedSessionStrategy;

@KeycloakConfiguration
@EnableGlobalMethodSecurity(prePostEnabled = true)
@Import({KeycloakSpringBootConfigResolver.class})
public class KeycloakAdapterConfig extends KeycloakWebSecurityConfigurerAdapter {

    /* Registers the KeycloakAuthenticationProvider with the authentication manager.*/
    @Autowired
    public void configureGlobal(AuthenticationManagerBuilder auth) throws Exception {
        KeycloakAuthenticationProvider keycloakAuthenticationProvider = keycloakAuthenticationProvider();
        keycloakAuthenticationProvider.setGrantedAuthoritiesMapper(new SimpleAuthorityMapper());
        auth.authenticationProvider(keycloakAuthenticationProvider);
    }

    /* Defines the session authentication strategy null means no session.*/
    @Bean
    @Override
    protected NullAuthenticatedSessionStrategy sessionAuthenticationStrategy() {
        return new NullAuthenticatedSessionStrategy();
    }

    @Override
    protected void configure(HttpSecurity http) throws Exception {
        super.configure(http);

        http.csrf()
                .disable()
                .authorizeRequests()
                .antMatchers(HttpMethod.GET, "/super-hello/")
                .hasRole("superuser");
    }
}
```

Well, this is a docker tutorial, so all you need to understand from this file is the `antMatchers` these specify the path, HTTP method, and who can use it, here we have a `GET` method on the route `/super-hello` that only can be used by a user with the `superuser` role.

AND NOW for the REST API, we'll need to add the endpoint `/super-hello`, modify the `HelloController`, and add:

```java
...
import org.springframework.security.access.prepost.PreAuthorize;
...
    @PreAuthorize("hasRole('superuser')")
    @GetMapping("/super-hello/{name}")
    public String superGreet(@PathVariable String name) {
        return String.format("Super Hello, %s", name);
    }
...
```

Finally (not really), we need to add some Keycloak configuration to `application.yml`

```yml
keycloak:
  realm: dori
  auth-server-url: "http://localhost:8080/"
  resource: dori-client
  public-client: true
  bearer-only: true
```

So what about that `/super-hello` request?, if we requested it'll give us a 401 (Unauthorized) status code, So we need a token right?

We can get a token, by making a `token` request to the Keycloak server

```bash
curl -X POST http://localhost:8080/realms/dori/protocol/openid-connect/token\
   -H 'Content-Type: application/x-www-form-urlencoded' \
   -d 'client_id=dori-client&client_secret=YOUR_CLIENT_SECRET&grant_type=password&username=nemo&password=nemorocks'
```

then you should get a response like this

```json
{
  "access_token": "",
  "expires_in": 300,
  "refresh_expires_in": 1800,
  "refresh_token": "",
  "token_type": "Bearer",
  "not-before-policy": 0,
  "session_state": "",
  "scope": "email profile"
}
```

As you can see we have the access token to the client `dori-client` from the `dori` realm, using the user `nemo`, now when we use the token with the `/super-hello` it'll work.

```bash
curl http://localhost:8081/super-hello/Eloi\
    -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

This is so cool, right? But we're not done yet, we need to containerize the thing right?

Now we'll introduce [docker-compose](https://docs.docker.com/compose/compose-file/) that will allow us to run more than one container at the same time (not really, but it appears to do that) with a related setup, in this case, we need a network between the Spring Boot, and the Keycloak server, finally things are getting along :)

This file will be at the root of the whole project.

```yml
# docker-compose.yml
version: "3.8"

services:
  auth:
    image: "quay.io/keycloak/keycloak:20.0.2"
    container_name: "auth"
    restart: "always"
    ports:
      - 9090:8080
    environment:
      KEYCLOAK_ADMIN: "admin"
      KEYCLOAK_ADMIN_PASSWORD: "admin"
    volumes:
      - ./auth/realms_backups/:/tmp/backups/
    command: "-Dkeycloak.profile.feature.upload_scripts=enabled -Dkeycloak.migration.action=import -Dkeycloak.migration.realmName=dori -Dkeycloak.migration.provider=dir -Dkeycloak.migration.dir=/tmp/backups start-dev"
    networks:
      - auth-backend

  backend:
    build: ./server
    ports:
      - 8080:8081
    depends_on:
      - auth
    networks:
      - auth-backend

networks:
  auth-backend: {}
```

So..., what's going on here?\
If you look close enough you'll notice something we've seen before, aside from the other configuration, the services have `ports` property which will do port forwarding the same as `-p`, `environment` is like `-e`, `volumes` is like `-v` when using `docker run`.

`command` overrides `CMD` in the Dockerfile, meaning whatever we put in there will be executed when the container starts.

Now let's go over the compose file, it's just a YAML file that tells `docker-compose` what to do, first, we have the `version` which is the version of the compose file, currently, the latest version is `3.8` so we're gonna use that, now for the services array, for starters, we have the `auth` service is the Keycloak server, where we specify the wanted docker image using the `image` property.

`build` specifies where the docker project is, i.e. a project with a `Dockerfile` at its root, and build has more interesting stuff that can be found in [here](https://docs.docker.com/compose/compose-file/build/)

`depends_on` states that the `backend` service will not run until the `auth` service has started.

`container_name` specifies what name this image's container will be using while it's running, so it can be accessed from the docker network (for now that's all that we need from the name), as you can see each one of the services has a `network` property which is an array that represents the networks that the container will be connected to, in this case, `auth` and `backend` are connected to.

Now where to get the network?

as you can see at the end of the file we can see a `networks` array that defines our networks, and here we've defined a network called `auth-backend` that will connect the Spring Boot server to the Keycloak server, easy eh?

well, it's not as easy as it seems, but that's all we need for this setup, you check out more about networks [here](https://docs.docker.com/compose/networking/).

As I said `container_name` will help with the network, but how, well now that the Keycloak service is named `auth` that will be used as the server's address.

Now we can change the Keycloak address in `application.yml` to `http://auth:8080/`, here we're using port 8080 since that's the server address inside the docker network, we can still use `http://localhost:9090/` if we want to, but it's more convenient to use the docker network. and now it's time for action.

Run `docker compose up` and it'll build the project for the first time, and start it, but if anything changes it won't re-build the project with the newest changes, so we need to run `docker compose build` after each change and the changed image will be rebuilt, and ready for running.

Just a little test to make sure that everything is in its place.

acquire the access token first, from the Keycloak server.

```bash
curl -X POST http://localhost:9090/realms/dori/protocol/openid-connect/token\
   -H 'Content-Type: application/x-www-form-urlencoded' \
   -d 'client_id=dori-client&client_secret=YOUR_CLIENT_SECRET&grant_type=password&username=nemo&password=nemorocks'
```

this will work, but if you make `/super-hello` with the token returned from the previous request, it won't work, because the token was issued to the address `localhost:9090` and the Spring Boot requests the server at `auth:8080`, and Keycloak is careful who can use the token and who can't.

so to avoid situations like this, we can easily issue the token from our Spring Boot server, we'll create an endpoint `/token` that will make a request to the Keycloak server and retrieve a token that was issued for the same address.

First, we need to add a JSON Utility Dependency, since as we've seen earlier the response from the Keycloak server is a JSON.

```xml
<!-- pom.xml -->
...
		<dependency>
			<groupId>org.json</groupId>
			<artifactId>json</artifactId>
			<version>20220924</version>
		</dependency>
...
```

and now we'll create a controller that does the token retrieving request:

```java
// controllers/TokenController.java

import org.apache.http.client.config.RequestConfig;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import org.apache.http.NameValuePair;
import org.apache.http.client.entity.UrlEncodedFormEntity;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.impl.client.HttpClientBuilder;
import org.apache.http.util.EntityUtils;
import org.json.JSONObject;
import org.springframework.beans.factory.annotation.Value;

import java.util.List;
import java.util.Map;

@RestController
public class TokenController {
    @Value("${keycloak.auth-server-url}")
    private String authServerURL;

    @PostMapping("/token")
    public ResponseEntity<?> login(@RequestBody Map<String, String> user) {
        try {
            var form = List.of(
                    new NameValuePairImpl("client_id", "dori-client"),
                    new NameValuePairImpl("client_secret", "YOUR_CLIENT_SECRET"),
                    new NameValuePairImpl("grant_type", "password"),
                    new NameValuePairImpl("username", user.get("username")),
                    new NameValuePairImpl("password", user.get("password"))
            );

            var requestConfig = RequestConfig.custom().build();
            var httpClient = HttpClientBuilder.create().setDefaultRequestConfig(requestConfig).build();
            var request = new HttpPost(String.format("%s/realms/dori/protocol/openid-connect/token", authServerURL));
            request.setEntity(new UrlEncodedFormEntity(form));
            JSONObject json = new JSONObject(EntityUtils.toString(httpClient.execute(request).getEntity()));

            return ResponseEntity.ok(Map.of("token", json.get("access_token")));
        } catch (Exception e) {
            return ResponseEntity.internalServerError().body(e.toString());
        }
    }
}

class NameValuePairImpl implements NameValuePair {
    private final String name;
    private final String value;

    public NameValuePairImpl(String name, String value) {
        this.name = name;
        this.value = value;
    }

    @Override
    public String getName() {
        return name;
    }

    @Override
    public String getValue() {
        return value;
    }
}
```

This controller only has one endpoint, that is `/token`, so we'll send a json with `username` and `password`, which will be used for logging in to the Keycloak realm.

Now we can rebuild the images, and test the requests again.

```bash
docker compose build
docker compose up

curl -X POST http://localhost:8080/token\
	-H "Content-Type: application/json"\
	--data '{"username": "nemo", "password": "nemorocks"}'

curl http://localhost:8080/super-hello/Eloi\
	-H "Authorization: Bearer ACCESS_TOKEN"
```

And gladly I can finally say that this part is over.

<br/>

---

## Dockerizing MariaDB

This part is cuter than Keycloak, since we'll just create a model, a simple controller, and modify some configuration files, that should be easy.

First, we need to configure Spring Boot with JPA, now we need JPA and MariaDB dependency.

```xml
<!-- pom.xml -->
...
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-data-jpa</artifactId>
			<version>2.7.8</version>
		</dependency>

		<dependency>
			<groupId>org.mariadb.jdbc</groupId>
			<artifactId>mariadb-java-client</artifactId>
			<scope>runtime</scope>
		</dependency>
...
```

Update your dependency tree using

```bash
mvn dependency:resolve
```

And update your `application.yml` to use MariaDB with JPA.

```yml
server:
  port: 8081

spring:
  datasource:
    url: "jdbc:mariadb://db/someDB?useJDBCCompliantTimezoneShift=true&useLegacyDatetimeCode=false&serverTimezone=UTC"
    username: "root"
    password: "hello"
    driver-class-name: "org.mariadb.jdbc.Driver"
  jpa:
    generate-ddl: true

keycloak:
  realm: dori
  auth-server-url: "http://auth:8080/"
  resource: dori-client
  public-client: true
  bearer-only: true
```

Now for the model, we'll be using a book model with string title attribute.

```java
// models/Book.java
import javax.persistence.*;

@Entity(name = "books")
public class Book {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;

    private String title;

    public void setId(Integer id) {
        this.id = id;
    }

    public Integer getId() {
        return id;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }
}
```

The repo

```java
// repos/BookRepo.java
import com.example.demo.models.Book;
import org.springframework.data.jpa.repository.JpaRepository;

public interface BookRepo extends JpaRepository<Book, Integer> {
}
```

and the controller

```java
// controllers/BookController.java
import com.example.demo.models.Book;
import com.example.demo.repos.BookRepo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;
import java.util.List;

@RestController
@RequestMapping("/book")
public class BookController {
    @Autowired
    private BookRepo bookRepo;

    @GetMapping()
    public List<Book> listBooks() {
        return bookRepo.findAll();
    }

    @PostMapping()
    public void addBook(@RequestBody Book book) {
        bookRepo.save(book);
    }
}
```

now back to docker, we'll add the MariaDB container configuration to `docker-compose.yml`

```yml
# docker-compose.yml
version: "3.8"

services:
  auth:
    image: "quay.io/keycloak/keycloak:20.0.2"
    container_name: "auth"
    restart: "always"
    ports:
      - 9090:8080
    environment:
      KEYCLOAK_ADMIN: "admin"
      KEYCLOAK_ADMIN_PASSWORD: "admin"
    volumes:
      - ./auth/realms_backups/:/tmp/backups/
    command: "-Dkeycloak.profile.feature.upload_scripts=enabled -Dkeycloak.migration.action=import -Dkeycloak.migration.realmName=dori -Dkeycloak.migration.provider=dir -Dkeycloak.migration.dir=/tmp/backups/ start-dev"
    networks:
      - auth-backend

  db:
    image: "mariadb:10.9"
    container_name: "db"
    restart: "always"
    environment:
      MARIADB_ROOT_PASSWORD: "hello"
      MARIADB_DATABASE: "someDB"
    ports:
      - 3306
    volumes:
      - db-config:/etc/mysql
      - db-data:/var/lib/mysql
    networks:
      - db-backend

  backend:
    build: ./server
    ports:
      - 8080:8081
    depends_on:
      - auth
      - db
    networks:
      - auth-backend
      - db-backend

networks:
  auth-backend: {}
  db-backend: {}

volumes:
  db-config:
  db-data:
```

Here's a new attribute in the house `volumes` which defines volumes that can be used by the containers, and MariaDB needs a configuration, and data volumes, to keep the database's data persistently.

Now re-build and run the containers

```bash
docker compose build
docker compose up
```

and we can test the setup now

```bash
curl -X POST http://localhost:8080/book\
	-H "Content-Type: application/json"\
	--data '{"title": "The Alchemist"}'
```

and we can retrieve that, just to make sure

```bash
curl http://localhost:8080/book
```

Now we can see that everything is in its place. See told you this was easy :)

<br/>

---

## Final Round, Wrapping everything up with a little frontend SvelteKit

First, we'll create our SvelteKit skeleton project using npm:

```bash
 npm init svelte@latest client
```

use these configs:

```
✔ Which Svelte app template? › Skeleton project
✔ Add type checking with TypeScript? › Yes, using TypeScript syntax
✔ Add ESLint for code linting? … No / Yes
✔ Add Prettier for code formatting? … No / Yes
✔ Add Playwright for browser testing? … No / Yes
✔ Add Vitest for unit testing? … No / Yes
```

then install the project's dependencies

```bash
 cd client
 npm install
```

add some stuff to `src/routes/+page.svelte` to make it interactive with the backend

```svelte
<!-- src/routes/+page.svelte -->
<script lang="ts">
    import {onMount} from "svelte"

    let title: string;
    let books: {title: string}[];

    async function createBook() {
        await fetch("http://localhost:8080/book", {
            method: "POST",
            mode: "cors",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({title: title}),
        })
        .then(resp => {
            if (resp.ok) {
                updateBooksList();
            }
        })
        .catch(err => {
            console.error(err);
        })
    }

    async function updateBooksList() {
        books = await fetch("http://localhost:8080/book", {
            method: "GET",
            mode: "cors",
        })
        .then(resp => resp.json())
        .then(fetchedBooks => fetchedBooks) as {title: string}[];
    }

    onMount(async () => {
        await updateBooksList();
    })
</script>

<div>
    <input bind:value={title} placeholder="Book Title" />
    <button on:click={createBook}>Add Book</button>

    <br/>

    {#if books}
    <title>Book:</title>
    <ul id="books">
        {#each books as book}
            <li>
                {book.title}
            </li>
        {/each}
    </ul>
    {/if}
</div>
```

Now for the docker part, first install `@sveltejs/adapter-node` to make it a standalone server, to save the effort of making a server, and dealing with the routes, but keep in mind that the node adapter uses port 3000.

Then update `svelte.config.js`

```js
// import adapter from "@sveltejs/adapter-auto";
import adapter from "@sveltejs/adapter-node";
```

add the client's Dockerfile

```dockerfile
FROM node:16-alpine as build

WORKDIR /app

COPY . .
RUN npm i
RUN npm run build

FROM node:16-alpine as run

WORKDIR /app

COPY --from=build /app/package*.json ./
COPY --from=build /app/build ./

EXPOSE 3000
CMD ["node", "./index.js"]
```

And now, for the final version of `docker-compose.yml`

```yml
# docker-compose.yml
version: "3.8"

services:
  auth:
    image: "quay.io/keycloak/keycloak:20.0.2"
    container_name: "auth"
    restart: "always"
    ports:
      - 9090:8080
    environment:
      KEYCLOAK_ADMIN: "admin"
      KEYCLOAK_ADMIN_PASSWORD: "admin"
    volumes:
      - ./auth/realms_backups/:/tmp/backups/
    command: "-Dkeycloak.profile.feature.upload_scripts=enabled -Dkeycloak.migration.action=import -Dkeycloak.migration.realmName=dori -Dkeycloak.migration.provider=dir -Dkeycloak.migration.dir=/tmp/backups/ start-dev"
    networks:
      - auth-backend

  db:
    image: "mariadb:10.9"
    container_name: "db"
    restart: "always"
    environment:
      MARIADB_ROOT_PASSWORD: "hello"
      MARIADB_DATABASE: "someDB"
    ports:
      - 3306
    volumes:
      - db-config:/etc/mysql
      - db-data:/var/lib/mysql
    networks:
      - db-backend

  backend:
    build: ./server
    ports:
      - 8080:8081
    depends_on:
      - auth
      - db
    networks:
      - auth-backend
      - db-backend

  frontend:
    build: ./client
    depends_on:
      - backend
    ports:
      - 8081:3000

networks:
  auth-backend: {}
  db-backend: {}

volumes:
  db-config:
  db-data:
```

As usual, build and run, and you should see some results.

And now we're done.
