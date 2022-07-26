# OAuth for SPA (single page application)

The guide of using Mixin OAuth in the SPA without exposing client_secret.

## Why I wrote this

In the [Mixin official documentation](https://developers.mixin.one/docs/api/oauth/oauth), We can know that the only method provided to use OAuth of Mixin is to visit `https://mixin.one/oauth/authorize` first. Which means the user will have to leave the application first to get authenticated, and being redirected to the original page. 

I thought it was the only way to do that before I saw [4swap](https://4swap.org) used a different method. When you're logging into 4swap, you won't leave the current page. The whole process would be completed within the same page. Which is awesome. So I want to figure it out.


## How did I find the answer

4swap's website is [open sourced](https://github.com/fox-one/4swap-web). The auth page must be related with keyword `auth` / `oauth`, then I found this [page](https://github.com/fox-one/4swap-web/blob/develop/src/components/modals/AuthModal.vue) which contains a component call `f-auth-method-modal`.

This component doesn't exist in `4swap-web`. So I looked for other repos and found this [uikit](https://github.com/fox-one/uikit). And there're a tons of `f` prefixed components.

The OAuth component for mixin messenger is in [here](https://github.com/fox-one/uikit/blob/main/packages/uikit/src/components/FAuthMethodModal/FAuthMixinMessenger.vue), Inside `<script> -> mounted()`, there's a function called `authorize`. I considered it's what makes it all works and opened [it](https://github.com/fox-one/uikit/blob/main/packages/uikit/src/utils/authorize.ts). 

I won't explain every line of it, basically, it's all about a websocket connection and [PKCE](https://oauth.net/2/pkce/). 


### Related links

[4swap/authorize.ts](https://github.com/fox-one/uikit/blob/main/packages/uikit/src/utils/authorize.ts)

[4swap/oauth.js](https://github.com/MixinNetwork/mixin.one/blob/master/src/oauth/index.js)

[mixin.one/authorization.js](https://github.com/MixinNetwork/mixin.one/blob/master/src/api/authorization.js)

[mixin.one/index.js](https://github.com/MixinNetwork/mixin.one/blob/master/src/oauth/index.js)

I wrapped those two files into one. So you will only need to import the `spa-oauth.js` and use it like `spa-oauth.vue`.

## Files
Javascript helper: [spa-oauth.js](spa-oauth.js)

Vue use case: [spa-oauth.vue](spa-oauth.vue)