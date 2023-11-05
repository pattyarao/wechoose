<script lang="ts">
  import { onMount } from "svelte";
  interface PostType {
    name: string;
    title: string;
    choices: ChoiceType[];
  }

  interface ChoiceType {
    title: string;
    votes: number;
  }

  let data: PostType[];

  onMount(async () => {
    const res = await fetch(import.meta.env.VITE_API_URL + "/post/");
    data = await res.json();
    // console.log(import.meta.env.VITE_API_URL);
    console.log(data);
  });
</script>

<div class="w-[80%] flex justify-center items-center">
  <div class="w-full flex flex-col gap-8">
    {#if data}
      {#each data as data}
        <div class="p-4 flex flex-col gap-4 rounded-md bg-black/20">
          <div class="p-2 flex items-center gap-2">
            <svg
              class="w-6 h-6 text-gray-800 dark:text-white"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              viewBox="0 0 20 20"
            >
              <path
                d="M10 0a10 10 0 1 0 10 10A10.011 10.011 0 0 0 10 0Zm0 5a3 3 0 1 1 0 6 3 3 0 0 1 0-6Zm0 13a8.949 8.949 0 0 1-4.951-1.488A3.987 3.987 0 0 1 9 13h2a3.987 3.987 0 0 1 3.951 3.512A8.949 8.949 0 0 1 10 18Z"
              />
            </svg>
            <p class="text-white">Posted by {data.name}</p>
          </div>
          <div class="p-4 rounded-md flex flex-col gap-2.5 bg-black/40">
            <p class="text-lg text-white font-bold">Question: {data.title}</p>

            {#if data.choices.length > 2}
              <div class="w-full flex flex-wrap gap-2 text-center">
                {#each data.choices as choice}
                  <p
                    class="text-gray-400 border border-gray-500 p-2 rounded-sm"
                  >
                    {choice.title}
                  </p>
                  <p
                    class="text-gray-400 border border-gray-500 p-2 rounded-sm"
                  >
                    {choice.votes}
                  </p>
                {/each}
              </div>
            {:else}
              <div class="w-full flex flex-col gap-2">
                {#each data.choices as choice}
                  <div class="w-full flex items-center justify-between gap-2">
                    <p
                      class="w-[90%] text-gray-400 border border-gray-500 p-2 rounded-sm hover:text-white hover:border-white hover:bg-white/50 transition ease-in duration-70"
                    >
                      {choice.title}
                    </p>
                    <p
                      class="w-[10%] text-gray-400 border border-gray-500 p-2 rounded-sm text-center"
                    >
                      {choice.votes}
                    </p>
                  </div>
                {/each}
              </div>
            {/if}
          </div>
        </div>
      {/each}
    {:else}
      <p>Loading...</p>
    {/if}
  </div>
</div>
