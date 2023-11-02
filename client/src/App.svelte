<script>
  import svelteLogo from "./assets/svelte.svg";
  import viteLogo from "/vite.svg";
  import Counter from "./lib/Counter.svelte";
  import TailwindCss from "./lib/TailwindCSS.svelte";
  import InputName from "./lib/InputName.svelte";
  import Signup from "./lib/components/Signup.svelte";
</script>

<TailwindCss />

<main class="w-full min-h-screen flex items-center justify-center">
  <Signup />
</main>
