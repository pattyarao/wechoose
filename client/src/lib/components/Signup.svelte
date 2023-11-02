<script lang="ts">
  import decision_fatigue from "../../assets/images/decision-fatigue.png";

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
</script>

<div
  class="w-full min-h-screen flex items-center justify-center bg-gradient-to-b from-[#18122B] to-[#635985] via-[#443C68]"
>
  <div class="w-1/2 flex flex-col items-center justify-center">
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
            class="bg-gradient-to-b from-[#FFFBF5] to-[#C3ACD0] px-8 py-2 rounded-md hover:scale-105 transition ease-in duration-70"
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
        <div class="w-full flex items-center justify-end gap-8">
          <button
            on:click={subtractIndex}
            class="rounded-md px-8 py-2 text-white hover:bg-[#F4F4F2]/50 hover:text-black transition ease-in duration-70 underline"
            >Back</button
          >
          <button
            on:click={addIndex}
            class="bg-gradient-to-b from-[#FFFBF5] to-[#C3ACD0] px-8 py-2 rounded-md hover:scale-105 transition ease-in duration-70"
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
        <div class="w-full flex items-center justify-end gap-8">
          <button
            on:click={subtractIndex}
            class="rounded-md px-8 py-2 text-white hover:bg-[#F4F4F2]/50 hover:text-black transition ease-in duration-70 underline"
            >Back</button
          >
          <button
            on:click={addIndex}
            class="bg-gradient-to-b from-[#FFFBF5] to-[#C3ACD0] px-8 py-2 rounded-md"
            >Create Account</button
          >
        </div>
      </div>
    {/if}
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
