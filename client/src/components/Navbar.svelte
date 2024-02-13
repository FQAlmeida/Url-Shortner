<script lang="ts">
	import { auth } from '$lib/auth';
	import { session } from '@store/auth';
	import { signOut } from 'firebase/auth';
	import {
		Navbar,
		NavBrand,
		NavHamburger,
		NavUl,
		DarkMode,
		A
	} from 'flowbite-svelte';
	let loggedIn: boolean = false;
	let loading: boolean = true;
	session.subscribe((v) => {
		loggedIn = v?.loggedIn ?? false;
		loading = v?.loading ?? true;
	});
</script>

<Navbar let:hidden let:toggle class="">
	<NavBrand href="/">
		<img
			src="https://flowbite-svelte.com/images/flowbite-svelte-icon-logo.svg"
			class="me-3 h-6 sm:h-9"
			alt="Flowbite Logo"
		/>
		<span
			class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
		>
			URL Shortner
		</span>
	</NavBrand>
	<NavHamburger on:click={toggle} />
	<div class="flex flex-row">
		<NavUl {hidden}>
			{#if loggedIn}
				<A
					on:click={async () => {
						session.update((old) => {
							return { ...old, loading: true };
						});
						await signOut(auth).catch((e) => {
							console.error(e);
							throw e;
						});
						session.update((old) => {
							return {
								...old,
								user: null,
								loading: false,
								loggedIn: false
							};
						});
					}}
				>
					Logout
				</A>
			{:else}
				<A href="/login">Login</A>
				<A href="/signup">Sign Up</A>
			{/if}
		</NavUl>
		<DarkMode />
	</div>
</Navbar>
