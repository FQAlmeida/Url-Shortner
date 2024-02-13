<script lang="ts">
	import { slugs } from '@store/slugs';
	import { Button, Heading, Input, Label, Modal, P } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';

	let url = '';
	let slug = '';

	$: slug_data = {
		slug: slug,
		redirect: url
	};

	let modal_open = false;
</script>

<div>
	<Heading class="text-center text-2xl">Criar novo Slug</Heading>
	<form
		class="flex flex-col gap-2"
		on:submit|preventDefault={() => {
			modal_open = true;
		}}
	>
		<Label for="url">URL</Label>
		<Input type="url" id="url" bind:value={url} />

		<Label for="slug">Slug</Label>
		<Input type="text" id="slug" bind:value={slug} />

		<Button type="submit">Adicionar</Button>
	</form>
</div>

<Modal bind:open={modal_open} size="md" autoclose outsideclose>
	<div class="text-center">
		<ExclamationCircleOutline
			class="mx-auto mb-4 h-9 w-9 text-gray-400 dark:text-gray-200"
		/>
		<Heading class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
			Tem certeza que deseja adicionar esse Slug?
		</Heading>
		<div class="mx-auto my-2 columns-1">
			<P>Slug: {slug}</P>
			<P>
				Endereço Real: {url}
			</P>
		</div>
		<Button
			color="red"
			class="me-2 mt-2"
			on:click={async () => await slugs.add_slug(slug_data)}>Confirmar</Button
		>
		<Button color="alternative">Cancelar</Button>
	</div>
</Modal>
