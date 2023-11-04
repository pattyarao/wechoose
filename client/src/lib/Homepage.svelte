<script lang="ts">
  import Navbar from "./components/Navbar.svelte";
  import PollsFeed from "./components/PollsFeed.svelte";
  import OngoingPolls from "./components/OngoingPolls.svelte";
</script>

<main
  class="w-full min-h-screen flex flex-col bg-gradient-to-b from-black to-[#331D2C]"
>
  <Navbar />
  <div class=" flex justify-between p-4">
    <OngoingPolls />
    <PollsFeed />
  </div>
</main>
