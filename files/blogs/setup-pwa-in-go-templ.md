# Setup

If you don't have a templ project set up, check out my previous [post](https://mbaraa.com/blog/setting-up-go-templ-with-tailwind-htmx-docker) to set up a templ project with tailwind css and htmx (we won't need htmx here, but it's good to know that it's there), right before the [Waltz](https://mbaraa.com/blog/setting-up-go-templ-with-tailwind-htmx-docker#toc_12) section, because until that point is only project set up, after it is the project.

The final project is available [here](https://github.com/mbaraa/pub_code/tree/main/blog/setup-pwa-in-go-templ).

# PWA (Progressive Web Application)

A (PWA) progressive web app or application, is a type of websites that has more power than a regular website, where it can access more hardware devices, can be installed as an application[\*](https://en.wikipedia.org/wiki/Progressive_web_app#Browswer_Support), has a background service, and more...

Nevertheless they're not as powerful as a native application say for Android and iOS, since they rely on a browser to stay alive, and that can be a bit problematic for devices with limited resources, but overall PWAs are like [electron](https://www.electronjs.org/) apps, since they're both web apps written in or transpiled into web technologies (HTML, CSS, and JavaScript), and ran using a web engine, the only difference that empowers PWAs, that if you have multiple PWAs installed on a device, they'll end up using the same browser that they were installed from, where for a bunch of electron apps, each of them will use a separate engine, and that's a bit annoying since they'll take up space, and eat resources when they're running.

I guess that's what's needed to know about PWA for now, more details [here](https://developer.mozilla.org/en-US/docs/Web/Progressive_web_apps/Guides/What_is_a_progressive_web_app).

# Service Workers

Service workers are what make PWAs tick, as they're lightweight (they're JavaScript, but a stripped down version of it, so it's kinda really lightweight), and their responsibility is to sync notifications in the background, save and load cached files, load and offline version of the app when there's no connection.

Well, that was the gist of it, as there's more into it, as the current implementation of JavaScript engines on modern browsers (IDK about IE11 or Safari) support the [Service Worker API](https://developer.mozilla.org/en-US/docs/Web/API/Service_Worker_API), where through it you can do various functionalities like, check the state of a particular service worker (yes multiple service workers can be registered within the same app), caching files, play a media file for a certain event, etc...

# Adding the juice

Now let's get working, I'll make changes based on the templ setup I mentioned earlier in [Setup](#toc_0), so for now we'll add the service worker and it's registration to the project.

Create a file under `static` called `service-worker.js`, this will be our app's service worker, it's a normal JavaScript file (at least for now), with the following content.

```js
console.log("yoohoo I'm installed!");
```

Now we need to register it into the application, and to do that we'll add a script to the body of the view's layout under `components/layout.templ`

```html
<script type="module">
  function registerServiceWorkers() {
    // check if the browser supports service workers, otherwise abort.
    if (!("serviceWorker" in navigator)) {
      console.error("Browser doesn't support service workers");
      return;
    }
    // it's really important to make sure that the app is fully loaded, since the service worker won't be loaded before that.
    window.addEventListener("load", () => {
      navigator.serviceWorker
        // this is the service worker's path, since we put it under /static earlier, and static is already registered as a file server.
        .register("/static/service-worker.js")
        .then((reg) => {
          console.log("Service Worker Registered", reg);
        })
        .catch((err) => {
          console.log("Service Worker Registration failed:", err);
        });
    });
  }
  // call the function, you might wonder why isn't everything working as expected, but then it's just a missing function call :)
  registerServiceWorkers();
</script>
```

Now `components/layout.templ` should look something like this.

```templ
package components

templ Layout(children ...templ.Component) {
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8"/>
			<meta name="viewport" content="width=device-width, initial-scale=1"/>
			<title>Templ PWA</title>
			<link href="/static/css/tailwind.css" rel="stylesheet"/>
		</head>
		<body>
			for _, child := range children {
				@child
			}
            <script type="module">
              function registerServiceWorkers() {
                // check if the browser supports service workers, otherwise abort.
                if (!("serviceWorker" in navigator)) {
                  console.error("Browser doesn't support service workers");
                  return;
                }
                // it's really important to make sure that the app is fully loaded, since the service worker won't be loaded before that.
                window.addEventListener("load", () => {
                  navigator.serviceWorker
                    // this is the service worker's path, since we put it under /static earlier, and static is already registered as a file server.
                    .register("/static/service-worker.js")
                    .then((reg) => {
                      console.log("Service Worker Registered", reg);
                    })
                    .catch((err) => {
                      console.log("Service Worker Registration failed:", err);
                    });
                });
              }
              // call the function, you might wonder why isn't everything working as expected, but then it's just a missing function call :)
              registerServiceWorkers();
            </script>
		</body>
	</html>
}
```

Now let's actually add some content to our app, modify `components/index.templ` to the following.

```templ
package components

templ Index() {
	@Layout(main())
}

templ main() {
	<main class="w-full h-screen bg-black">
		<iframe class="w-screen h-screen" src="https://www.youtube.com/embed/dQw4w9WgXcQ" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
	</main>
}
```

# Yeppeee

Now that everything is in place, we can check the console to find that the service worker is working from the console, it should be printing the message we specified earlier, FireFox won't display the message as it doesn't allow service workers to run from a non https websites, you can check in Google Chrome tho, and you'll see the message, or from the `Application` tab from the dev tools, where you can see the active service worker over there.

![Service Worker Status in Chrome](/img/service-worker-chrome.png)

Service worker status in Google Chrome.

![Service Worker Status in FireFox](/img/service-worker-firefox.png)

Service worker status in Mozilla FireFox.

Well, that's it, the next post will be about Firebase Cloud Messaging with templ and PWA, so this is just warm up :)

# Quote of the day

"Computer science education cannot make anybody an expert programmer any more than studying brushes and pigment can make somebody an expert painter."
\
\- [Eric S. Raymond](https://en.wikipedia.org/wiki/Eric_S._Raymond)
