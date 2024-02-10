<script lang="ts">
	import { Spinner } from 'flowbite-svelte';
	export let size: number = 16;
</script>

<div class="flex min-h-full min-w-full items-center justify-center">
	<Spinner {size} />
</div>
