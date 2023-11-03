<script lang="ts">
  import decision_fatigue from "../../assets/images/decision-fatigue.png";
  import { navigate } from "svelte-routing";

  let username: string;
  let password: string;
  let confirmPassword: string;

  let index: number = 0;
  const addIndex = () => {
    index += 1;
  };
  const subtractIndex = () => {
    index -= 1;
  };

  function handleNavigateOnboarding() {
    navigate("/", { replace: true });
  }
</script>

<div
  class="w-[70vw] min-h-[70vh] flex items-center justify-center bg-gradient-to-b from-[#331D2C] to-[#A78295] via-[#3F2E3E] rounded-xl"
>
  <div class="w-1/2 flex flex-col items-center justify-center gap-20">
    {#if index === 0}
      <div class="w-3/4 flex flex-col items-start gap-6">
        <h1 class="text-xl text-white font-bold">Input username</h1>
        <input
          type="text"
          placeholder="ex. nathancutie"
          class="w-full py-2 bg-inherit border-b border-white focus:outline-none text-white"
          bind:value={username}
        />
        <div class="w-full flex justify-end">
          <button
            on:click={addIndex}
            class="text-white text-sm underline hover:scale-105 transition ease-in duration-70"
            >Next</button
          >
        </div>
      </div>
    {:else if index === 1}
      <div class="w-3/4 flex flex-col items-start gap-6">
        <h1 class="text-xl text-white font-bold">Input password</h1>
        <input
          type="password"
          class="w-full py-2 bg-inherit border-b border-white focus:outline-none text-white"
          bind:value={password}
        />
        <div class="w-full flex items-center justify-end gap-4">
          <button
            on:click={subtractIndex}
            class="text-white text-sm underline hover:scale-105 transition ease-in duration-70"
            >Back</button
          >
          <button
            on:click={addIndex}
            class="text-white text-sm underline hover:scale-105 transition ease-in duration-70"
            >Next</button
          >
        </div>
      </div>
    {:else if index === 2}
      <div class="w-3/4 flex flex-col items-start gap-6">
        <h1 class="text-xl text-white font-bold">Confirm Password</h1>
        <input
          type="password"
          class="w-full py-2 bg-inherit border-b border-white focus:outline-none text-white"
          bind:value={confirmPassword}
        />
        <div class="w-full flex items-center justify-end gap-4">
          <button
            on:click={subtractIndex}
            class="text-white text-sm underline hover:scale-105 transition ease-in duration-70"
            >Back</button
          >
          <button
            on:click={handleNavigateOnboarding}
            class="text-white text-sm underline hover:scale-105 transition ease-in duration-70"
            >Create Account</button
          >
        </div>
      </div>
    {/if}
    <button
      class="w-[70%] flex justify-center"
      on:click={handleNavigateOnboarding}
    >
      <p
        class="bg-gradient-to-b from-[#331D2C] to-[#3F2E3E] px-6 py-1.5 rounded-lg text-xs text-white"
      >
        Already have an account? <span class="underline cursor-pointer"
          >Sign in!</span
        >
      </p>
    </button>
  </div>
  <div class="w-1/2 flex justify-center">
    <a href="https://storyset.com/people">
      <img
        src={decision_fatigue}
        alt="https://storyset.com/people"
        class="w-[300px]"
      />
    </a>
  </div>
</div>
