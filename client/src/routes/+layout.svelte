<script lang="ts">
	import '../app.pcss';
	import Navbar from '@components/Navbar.svelte';
	import SpinnerLoader from '@components/SpinnerLoader.svelte';

	import { session } from '@store/auth';

	import type { LayoutData } from './$types';
	import { onMount } from 'svelte';
	import { error } from '@store/slugs';
	import Footer from '@components/Footer.svelte';
	import { Button, Heading, Modal, P } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';
	export let data: LayoutData;

	let loading: boolean = true;
	let loggedIn: boolean = false;

	session.subscribe((cur: any) => {
		loading = cur?.loading;
		loggedIn = cur?.loggedIn;
	});

	onMount(async () => {
		const user: any = await data.getAuthUser();

		const loggedIn = !!user && user?.emailVerified;
		session.update((cur: any) => {
			loading = false;
			return {
				...cur,
				user,
				loggedIn,
				loading: false
			};
		});
	});

	$: show_error_modal = !!$error;
</script>

<!-- {#if loading}
	<SpinnerLoader />
{:else} -->
<svelte:head>
	<script
		async
		src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-6267954547276558"
		crossorigin="anonymous"
	></script>
</svelte:head>
<div class="flex min-h-screen min-w-full flex-col">
	<Navbar />
	<div class="container mx-auto my-auto">
		<slot />
	</div>
	<Footer />
</div>
<Modal bind:open={show_error_modal} size="md" autoclose outsideclose>
	<div class="text-center">
		<ExclamationCircleOutline
			class="mx-auto mb-4 h-9 w-9 text-gray-400 dark:text-gray-200"
		/>
		<Heading class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
			Ocorreu um erro
		</Heading>
		<div class="mx-auto my-2 columns-1">
			<P>{$error}</P>
		</div>
		<Button color="red" class="me-2 mt-2" on:click={() => error.set(null)}>
			Confirmar
		</Button>
	</div>
</Modal>
<!-- {/if} -->
