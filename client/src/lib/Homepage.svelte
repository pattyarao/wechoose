<script lang="ts">
  import Navbar from "./components/Navbar.svelte";

  interface PostType {
    name: string;
    title: string;
    choices: string[];
  }

  const dummyData: PostType[] = [
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
</script>

<main
  class="w-full min-h-screen flex flex-col bg-gradient-to-b from-black to-[#331D2C]"
>
  <Navbar />

  <div class="w-full h-full flex items-center justify-center p-4">
    <div class="w-[70%] flex flex-col gap-8 text-white">
      {#each dummyData as data}
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
    </div>
  </div>
</main>
