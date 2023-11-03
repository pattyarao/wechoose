<script>
  import svelteLogo from "./assets/svelte.svg";
  import viteLogo from "/vite.svg";

  import TailwindCss from "./lib/TailwindCSS.svelte";

  import Signup from "./lib/components/Signup.svelte";
  import Signin from "./lib/components/Signin.svelte";

  import { Router, Link, Route } from "svelte-routing";
  import Onboarding from "./lib/Onboarding.svelte";
  import CreateAccount from "./lib/CreateAccount.svelte";
  import Homepage from "./lib/Homepage.svelte";
</script>

<TailwindCss />

<Router>
  <main class="w-full min-h-screen flex items-center justify-center">
    <!-- <Signup /> -->
    <!-- <Signin /> -->
    <Route path="/" component={Onboarding} />
    <Route path="/createaccount" component={CreateAccount} />
    <Route path="/homepage" component={Homepage} />
  </main>
</Router>
