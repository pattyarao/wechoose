<script>
  import wechoose_logo_ver3 from "../../assets/images/wechoose-logo-ver3.png";

  import { onMount, onDestroy } from "svelte";

  let scrolled = false;

  function handleScroll() {
    scrolled = window.scrollY > 0;
  }

  onMount(() => {
    window.addEventListener("scroll", handleScroll);
    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  });
</script>

<nav
  class="top-0 sticky w-full h-fit px-4 pt-4 flex gap-2"
  class:scrolled-background={scrolled}
>
  <div class="w-full p-2 flex justify-between">
    <div class="flex flex-col gap-2">
      <img src={wechoose_logo_ver3} alt="logo" class="w-[100px]" />
    </div>
    <div id="links" class="flex gap-4 text-[#EFE1D1] text-center">
      <p
        class="rounded-md p-2 text-sm flex items-center justify-center text-xs font-bold hover:bg-white/20 transition ease-in duration-100"
      >
        Polls
      </p>
      <p
        class="rounded-md p-2 text-sm flex items-center justify-center text-xs font-bold hover:bg-white/20 transition ease-in duration-100"
      >
        Create New Poll
      </p>
      <p
        class="rounded-md p-2 text-sm flex items-center justify-center text-xs font-bold hover:bg-white/20 transition ease-in duration-100"
      >
        View Profile
      </p>
    </div>
  </div>
</nav>

<style>
  .scrolled-background {
    background-color: #331d2ca5;

    transition: ease-in-out 0.5s;
    padding-bottom: 0;
    border-bottom: 1px solid black;
  }
</style>
