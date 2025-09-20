_**UPDATE: Firefox has added this in version 136, check [this post](https://support.mozilla.org/en-US/kb/use-sidebar-access-tools-and-vertical-tabs), so ignore what's below lol.**_

First install [Tree Style Tab](https://addons.mozilla.org/en-US/firefox/addon/tree-style-tab/) plugin.

Now for the tab hiding part:

1. Go to your profile directory, directory's path can be found by going to [about:support](about:support) and open the **Profile Directory**
1. Create a directory called `chrome`
1. Create a file called `userChrome.css`
1. Put these in the file

```css
#TabsToolbar {
  visibility: collapse;
}
```

1. Go to [about:config](about:config) and set `toolkit.legacyUserProfileCustomizations.stylesheets` to **true**
1. Restart Firefox to see the changes
