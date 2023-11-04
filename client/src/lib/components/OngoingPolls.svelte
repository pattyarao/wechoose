<script lang="ts">
  interface Choices {
    name: string;
    votes: number;
  }
  interface ActivePolls {
    title: string;
    choices: Choices[];
  }

  let dummyData: ActivePolls[] = [
    {
      title: "Poll 1",
      choices: [
        { name: "Burger King", votes: 50 },
        { name: "McDo", votes: 20 },
      ],
    },
    {
      title: "Poll 2",
      choices: [
        { name: "Maybelline", votes: 50 },
        { name: "MAC", votes: 20 },
      ],
    },
  ];
</script>

<div
  class="w-[28%] h-fit bg-[#F7F7F7]/20 p-4 rounded-lg flex flex-col gap-4 text-white"
>
  <h1 class="bg-[#A78295]/50 p-2 rounded-lg text-center">Current Polls</h1>
  <div class="w-full flex flex-col gap-2">
    {#each dummyData as data}
      <div class="w-full bg-[#331D2C]/70 flex flex-col gap-4 p-4 rounded-lg">
        <p class="border-b border-white">{data.title}</p>
        <div class="w-full flex flex-col gap-2">
          {#each data.choices as choice}
            <div class="flex items-center gap-4 justify-between">
              <p class="bg-[#EFE1D1]/30 w-[90%] p-1 rounded-md">
                {choice.name}
              </p>
              <p class="">{choice.votes}</p>
            </div>
          {/each}
        </div>
      </div>
    {/each}
  </div>
</div>
