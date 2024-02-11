<script lang="ts">
	import '../app.pcss';
	import Navbar from '@components/Navbar.svelte';
	import SpinnerLoader from '@components/SpinnerLoader.svelte';

	import { session } from '@store/auth';

	import type { LayoutData } from './$types';
	import { onMount } from 'svelte';
	import { slugs } from '@store/slugs';
	import Footer from '@components/Footer.svelte';
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
</script>

{#if loading}
	<SpinnerLoader />
{:else}
	<div class="flex min-h-screen min-w-full flex-col">
		<Navbar />
		<slot />
		<Footer />
	</div>
{/if}
