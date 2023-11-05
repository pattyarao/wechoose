<script lang="ts">
  import Navbar from "./components/Navbar.svelte";
  import PollsFeed from "./components/PollsFeed.svelte";
  import OngoingPolls from "./components/OngoingPolls.svelte";
</script>

<main class="w-full min-h-screen flex flex-col gap-10">
  <Navbar />
  <div class="w-full flex items-center justify-center p-4">
    <!-- <OngoingPolls /> -->
    <PollsFeed />
  </div>
</main>
