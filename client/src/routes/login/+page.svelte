<script lang="ts">
	import {
		signInWithEmailAndPassword,
		GoogleAuthProvider,
		signInWithPopup,
		type UserCredential
	} from 'firebase/auth';
	import { Label, Input, Button, Heading } from 'flowbite-svelte';
	import { GoogleSolid } from 'flowbite-svelte-icons';
	import { auth } from '$lib/auth';
	import { session } from '@store/auth';
	import { goto } from '$app/navigation';

	let email = '';
	let password = '';

	async function loginWithMail() {
		await signInWithEmailAndPassword(auth, email, password)
			.then((result) => {
				const { user }: UserCredential = result;
				session.set({
					loggedIn: true,
					user: {
						displayName: user?.displayName,
						email: user?.email,
						photoURL: user?.photoURL,
						uid: user?.uid
					}
				});
				goto('/dashboard');
			})
			.catch((error) => {
				return error;
			});
	}
	async function loginWithGoogle() {
		const provider = new GoogleAuthProvider();
		await signInWithPopup(auth, provider)
			.then((result) => {
				console.log(result);

				const { displayName, email, photoURL, uid } = result?.user;
				session.set({
					loggedIn: true,
					user: {
						displayName,
						email,
						photoURL,
						uid
					}
				});

				goto('/dashboard');
			})
			.catch((error) => {
				console.error(error);
				return error;
			});
	}
</script>

<Heading class="text-center text-2xl font-extrabold">Login</Heading>
<form>
	<div class="mb-6 flex flex-col gap-3">
		<div class="">
			<Label for="email" class="mb-2 block">Email</Label>
			<Input
				id="email"
				type="email"
				placeholder="emailexemplo@gmail.com"
				bind:value={email}
			/>
		</div>
		<div class="">
			<Label for="password" class="mb-2 block">Senha</Label>
			<Input
				id="password"
				type="password"
				placeholder="Senha"
				bind:value={password}
			/>
		</div>
	</div>
	<div class="flex h-12 flex-row gap-2">
		<Button class="min-h-full" on:click={loginWithMail}>Login</Button>
		<Button class="flex min-h-full flex-row gap-2" on:click={loginWithGoogle}>
			Signin with
			<GoogleSolid />
		</Button>
	</div>
</form>
