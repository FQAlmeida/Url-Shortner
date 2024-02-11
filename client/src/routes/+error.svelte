<script lang="ts">
	import { Heading, P } from "flowbite-svelte";

  export let message = "404 - Page Not Found";
</script>

<style>
  .container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    text-align: center;
    flex-direction: column;
    flex-grow: 1;
    padding-bottom: 2rem;
  }

  .message {
    font-size: 24px;
    font-weight: bold;
  }
</style>

<div class="flex flex-col flex-grow justify-center align-center pb-20 gap-2"> 
  <Heading class="text-3xl text-center">{message}</Heading>
  <P class="text-center">It is custom</P>
</div>
