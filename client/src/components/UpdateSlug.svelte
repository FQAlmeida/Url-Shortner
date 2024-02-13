<script lang="ts">
	import { slugs, type Slug } from '@store/slugs';
	import { Button, Heading, Input, Label, Modal, P } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';
	export let slug_base: Slug;
	export let modal_update_open = false;
	let modal_open = false;
	$: slug = slug_base.slug;
	$: redirect = slug_base.redirect;
	$: temp_slug = {
		id: slug_base.id,
		slug: slug,
		redirect: redirect
	};
</script>

<Modal bind:open={modal_update_open} size="sm" autoclose outsideclose>
	<Heading class="text-xl">Atualizando Slug {slug_base.id}</Heading>
	<Label for="url">URL</Label>
	<Input type="url" id="url" bind:value={redirect} />

	<Label for="slug">Slug</Label>
	<Input type="text" id="slug" bind:value={slug} />
	<Button
		class="mt-4"
		on:click={() => {
			modal_update_open = false;
			modal_open = true;
		}}
	>
		Atualizar
	</Button>
</Modal>
<Modal bind:open={modal_open} size="md" autoclose outsideclose>
	<div class="text-center">
		<ExclamationCircleOutline
			class="mx-auto mb-4 h-9 w-9 text-gray-400 dark:text-gray-200"
		/>
		<Heading class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
			Tem certeza sobre essas informações?
		</Heading>
		<div class="mx-auto my-2 columns-1">
			<P>Slug: {temp_slug?.slug}</P>
			<P>
				Endereço Real: {temp_slug?.redirect}
			</P>
		</div>
		<Button
			color="red"
			class="me-2 mt-2"
			on:click={async () => {
				return await slugs.update_slug(temp_slug);
			}}
		>
			Confirmar
		</Button>
		<Button color="alternative">Cancelar</Button>
	</div>
</Modal>
