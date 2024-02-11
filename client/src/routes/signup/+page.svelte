<script lang="ts">
	import { Label, Input, Button, Heading } from 'flowbite-svelte';
	import { auth } from '$lib/auth';
	import { createUserWithEmailAndPassword } from 'firebase/auth';
	import { session } from '@store/auth';
	import { goto } from '$app/navigation';
	// Initialize Firebase

	let email = '';
	let password = '';
	let confirm_password = '';

	async function handleRegister() {
		await createUserWithEmailAndPassword(auth, email, password)
			.then((result) => {
				const { user } = result;
				session.update((cur: any) => {
					return {
						...cur,
						user,
						loggedIn: true,
						loading: false
					};
				});
				goto('/dashboard');
			})
			.catch((error) => {
				throw new Error(error);
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
		<div class="">
			<Label for="confirm_password" class="mb-2 block">Cinfirmar Senha</Label>
			<Input
				id="confirm_password"
				type="password"
				placeholder="Confirmar Senha"
				bind:value={confirm_password}
			/>
		</div>
	</div>
	<div class="flex h-12 flex-row gap-2">
		<Button class="min-h-full" on:click={handleRegister}>Cadastrar</Button>
	</div>
</form>
