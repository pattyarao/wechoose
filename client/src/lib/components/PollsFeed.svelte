<script lang="ts">
  import { onMount } from "svelte";
  interface PostType {
    name: string;
    title: string;
    choices: string[];
  }

  let dummyData: PostType[] = [
    {
      name: "Patty Arao",
      title: "Where should I eat?",
      choices: ["McDo", "Burger King"],
    },
    {
      name: "Nathan Andres",
      title: "Where should I buy clothes?",
      choices: ["Uniqlo", "H&M"],
    },
    {
      name: "Neo Visda",
      title: "Which girl should I pick?",
      choices: ["Girl 1", "Girl 2"],
    },
    {
      name: "Ceejay Pascasio",
      title: "What phone should I buy",
      choices: ["iPhone X", "iPhone 11"],
    },
    {
      name: "Lei Araullo",
      title: "When is my birthday",
      choices: ["November 1", "November 2", "November 3", "November 4"],
    },
  ];

  let data: PostType[];

  onMount(async () => {
    const res = await fetch(import.meta.env.VITE_API_URL + "/post/");
    data = await res.json();
    // console.log(import.meta.env.VITE_API_URL);
  });
</script>

<div class="w-[70%] h-full flex">
  <div class="w-full flex flex-col gap-8 text-white">
    {#if data}
      {#each data as data}
        <div class="bg-[#F7F7F7]/20 p-4 flex flex-col gap-4 rounded-lg">
          <p>Posted by {data.name}</p>
          <div class="flex flex-col gap-2.5">
            <p class="font-bold text-xl text-white">{data.title}</p>

            {#if data.choices.length > 2}
              <div class="w-full flex flex-wrap gap-2 text-center">
                {#each data.choices as choice}
                  <p class="bg-[#EEEEEE]/40 p-2 rounded-md">
                    {choice}
                  </p>
                {/each}
              </div>
            {:else}
              <div class="w-full flex flex-col gap-2 text-center">
                {#each data.choices as choice}
                  <p class="bg-[#EEEEEE]/40 p-2 rounded-md">{choice}</p>
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
