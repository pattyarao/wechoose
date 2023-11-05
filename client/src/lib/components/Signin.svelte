<script lang="ts">
  import { navigate } from "svelte-routing";
  import critical_thinking from "../../assets/images/critical-thinking.png";

  let username: string;
  let password: string;

  function handleNavigateCreateAcc() {
    navigate("/createaccount", { replace: true });
  }

  function handleNavigateHomepage() {
    navigate("/homepage", { replace: true });
  }
</script>

<div class="w-[70vw] min-h-[70vh] flex items-center justify-center rounded-xl">
  <div class="w-1/2 flex flex-col gap-10 items-center justify-center">
    <div class="w-[70%]">
      <p class="w-full text-white text-sm">Username</p>
      <input
        type="text"
        required={true}
        class="w-full bg-inherit border-b border-white py-1.5 text-white focus:outline-none"
        bind:value={username}
      />
    </div>
    <div class="w-[70%]">
      <p class="w-full text-white text-sm">Password</p>
      <input
        type="password"
        required={true}
        class="w-full bg-inherit border-b border-white py-1.5 text-white focus:outline-none"
        bind:value={password}
      />
    </div>
    <div class="w-[70%] flex justify-end">
      <button
        class="text-white text-sm underline hover:scale-105 transition ease-in duration-70"
        on:click={handleNavigateHomepage}>Sign in</button
      >
    </div>
    <button
      class="w-[70%] flex justify-center"
      on:click={handleNavigateCreateAcc}
    >
      <p class="bg-black/50 px-6 py-1.5 rounded-lg text-xs text-[#F39F5A]">
        Don't have an account? <span class="underline cursor-pointer"
          >Create one now!</span
        >
      </p>
    </button>
  </div>
  <div class="w-1/2 flex justify-center">
    <a href="https://storyset.com/idea">
      <img
        src={critical_thinking}
        alt="https://storyset.com/idea"
        class="w-[300px]"
      />
    </a>
  </div>
</div>
