So, let's say that you wanna listen to a scroll event to fetch more content, or create a parallax effect, or anything else related to scrolling, and it happens that you write [Vue 3](https://vuejs.org/) , so all what you gotta do is utilize the [useWindowScroll](https://vueuse.org/core/useWindowScroll/) composable from [VueUse](https://vueuse.org).

So you'd do something like this

```vue
<template>
  <div style="width: 2000px; height: 2000px">
    Scroll X: {{ x }}
    <br />
    Scroll Y: {{ y }}
  </div>
</template>

<script setup lang="ts">
import { useWindowScroll } from "@vueuse/core";
const { x, y } = useWindowScroll(window);
</script>
```

And it's the exact same thing in [Nuxt 3](https://nuxt.com/), right?

**NO**

Nuxt has a different opinion, i.e it slaps `window is not defined` in our faces, ngl Nuxt is great and everything, and it was my insist to use a clean composable from VueUse, so that the code looks like a Vue code not some Vue with some classic JavaScript slapped into it.

Well `useWindowScroll` and SSR don't really mix, since it has the word `window` in it!

But unlike [SvelteKit](https://kit.svelte.dev/) where they managed to handle stuff like this with just

```svelte
<script lang="ts">
    let scrollY = 0;
    let scrollX = 0;
</script>

<div style="width: 2000px; height: 2000px">
  Scroll X: {{ scrollX }}
  <br />
  Scroll Y: {{ scrollY }}
</div>

<svelte:window bind:scrollX bind:scrollY />
```

Where this will do the same as it does in [Svelte](https://svelte.dev/), without any headaches and hacking around (you'll see in a bit).

So first of all we need to make sure we're on the browser, for that `onMounted` comes in handy.

```vue
<script setup lang="ts">
import { useWindowScroll } from "@vueuse/core";

onMounted(() => {
  const { x, y } = useWindowScroll(window);
});
</script>
```

But wait now `x` and `y` won't be accessible from the template, since they're locals in that arrow function.

So we'll create global ones and update them as the locals update

```vue
<template>
  <div style="width: 2000px; height: 2000px">
    Scroll X: {{ X }}
    <br />
    Scroll Y: {{ Y }}
  </div>
</template>

<script setup lang="ts">
import { useWindowScroll } from "@vueuse/core";

const X = ref(0);
const Y = ref(0);

onMounted(() => {
  const { x, y } = useWindowScroll(window);
  watch(x, (value) => {
    X.value = value;
  });
  watch(y, (value) => {
    Y.value = value;
  });
});
</script>
```

This works as expected, but let's say we need the same event listener somewhere else, we'd copy the same code, right?

Well that's just stupid ain't it? so now for the hacking part 🎉.

---

Luckily Nuxt supports client side plugins, i.e a plugin that only renders on the browser, that sounds like fun doesn't it?

Well it depends on how you define fun!

So first under `~/plugins` we'll create the `use-scroll.ts` file which will define the cool wrapper of `useWindowScroll` which will work in Nuxt!

```ts
// ~/plugins/use-scroll.ts
import { useWindowScroll } from "@vueuse/core";
export default defineNuxtPlugin((nuxtApp) => {
  const { x, y } = useWindowScroll(window);
  return {
    provide: {
      useScroll: () => {
        return { x, y };
      },
    },
  };
});
```

And off-course modify `~/nuxt.config.ts`

```ts
plugins: [{ src: "~/plugins/use-scroll.ts", ssr: false, mode: "client" }];
```

That way we abstracted the thing, we still have much to do, you'll see now

```vue
<script setup lang="ts">
const { $useScroll } = useNuxtApp();
const { x, y } = $useScroll();
</script>
```

And you'd expect that would work right out of the box right? well think again we need to check if `$useScroll` is defined correctly and it's not `undefined`.

So we'll check it manually

```vue
<template>
  <div style="width: 2000px; height: 2000px">
    Scroll X: {{ X }}
    <br />
    Scroll Y: {{ Y }}
  </div>
</template>

<script setup lang="ts">
const X = ref(0);
const Y = ref(0);

const { $useScroll } = useNuxtApp();
if (typeof $useScroll === "function") {
  const { x, y } = $useScroll();
  watch(x, (value) => {
    X.value = value;
  });
  watch(y, (value) => {
    Y.value = value;
  });
}
</script>
```

So basically it's same but using a "plugin" so that's just seems pointless, well I'm not done YET!

We need a fancy `composable`, cuz what's Vue 3 without composables, right?

Create this file that will hold the _fancy_ composable

```ts
// ~/composables/useScroll.ts
export default function () {
  const X = ref(0);
  const Y = ref(0);
  const { $useScroll } = useNuxtApp();
  if (typeof $useScroll === "function") {
    const { x, y } = $useScroll();
    watch(x, (value) => {
      X.value = value;
    });
    watch(y, (value) => {
      Y.value = value;
    });
    // the use of scrollX, and scrollY makes more sense than just x, and y
    return { scrollX: X, scrollY: Y };
  }
  // this return for when the plugin is not ready yet (ssr mode)
  return { scrollX: 0, scrollY: 0 };
}
```

It's a bit messy, but for a single time eh? and that's the beauty of composables!

Now we just use it in a component, and life goes on...

```vue
<template>
  <div style="width: 2000px; height: 2000px">
    Scroll X: {{ scrollX }}
    <br />
    Scroll Y: {{ scrollY }}
  </div>
</template>

<script setup>
const { scrollX, scrollY } = useScroll();
</script>
```

Thanks for reading the whole rant and solution.
