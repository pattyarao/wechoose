<script lang="ts">
  let name: string = "";
  const handleNameChange = () => {};
</script>

<div class="flex flex-col justify-center items-center gap-8">
  <input
    bind:value={name}
    type="text"
    required={true}
    placeholder="Input your name here."
    class="border border-black p-1.5"
  />
  {#if name}
    <p class="text-xl font-bold text-pink-300">Hello, {name}!</p>
  {/if}
</div>
