<script>
  import wechoose_logo from "../../assets/images/wechoose-logo.png";
  import weechose_logo_ver2 from "../../assets/images/we-choose-logo-ver2.png";
</script>

<nav class="w-full h-fit px-4 py-4 flex flex-col justify-center gap-2">
  <div class="flex items-center justify-between">
    <img src={weechose_logo_ver2} alt="logo" class="w-[120px]" />
    <!-- <div>
      <h1 class="font-black uppercase text-xl text-white">We Choose</h1>
    </div> -->

    <div id="links" class="flex gap-4 text-[#EFE1D1] text-center">
      <p
        class="rounded-md p-2 text-sm flex items-center justify-center text-xs font-bold"
      >
        Polls
      </p>
      <p
        class="rounded-md p-2 text-sm flex items-center justify-center text-xs font-bold"
      >
        Create New Poll
      </p>
    </div>

    <div class="py-2 rounded-md flex gap-10 items-center cursor-pointer">
      <div class="px-6 py-2 flex items-center gap-2 rounded-md bg-[#A78295]">
        <svg
          class="w-4 h-4 text-gray-800 dark:text-white"
          aria-hidden="true"
          xmlns="http://www.w3.org/2000/svg"
          fill="currentColor"
          viewBox="0 0 20 20"
        >
          <path
            d="M10 0a10 10 0 1 0 10 10A10.011 10.011 0 0 0 10 0Zm0 5a3 3 0 1 1 0 6 3 3 0 0 1 0-6Zm0 13a8.949 8.949 0 0 1-4.951-1.488A3.987 3.987 0 0 1 9 13h2a3.987 3.987 0 0 1 3.951 3.512A8.949 8.949 0 0 1 10 18Z"
          />
        </svg>
        <p class="text-xs font-bold text-white">Username</p>
      </div>
    </div>
  </div>
  <hr class="w-full border border-white" />
</nav>
