<script lang="ts">
	import { slugs, type Slug } from '@store/slugs';
	import {
		Button,
		Heading,
		Modal,
		P,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import {
		ExclamationCircleOutline,
		PenSolid,
		TrashBinSolid
	} from 'flowbite-svelte-icons';
	import UpdateSlug from './UpdateSlug.svelte';
	let show_delete_modal = false;
	let show_update_modal = false;
	let current_delete = '';
	let current_update = '';
	$: delete_slug = $slugs.find((slug) => slug.id === current_delete);
	$: slug_base =
		$slugs.find((slug) => slug.id === current_update) ?? ({} as Slug);
</script>

<div class="flex flex-col gap-4">
	<Heading class="text-center text-2xl">Gereciar Slugs</Heading>
	<Table shadow hoverable>
		<TableHead>
			<TableHeadCell>ID</TableHeadCell>
			<TableHeadCell>Redirect</TableHeadCell>
			<TableHeadCell>Slug</TableHeadCell>
			<TableHeadCell></TableHeadCell>
			<TableHeadCell></TableHeadCell>
		</TableHead>
		<TableBody tableBodyClass="divide-y">
			{#each $slugs as slug}
				<TableBodyRow>
					<TableBodyCell>{slug.id}</TableBodyCell>
					<TableBodyCell>{slug.redirect}</TableBodyCell>
					<TableBodyCell>{slug.slug}</TableBodyCell>
					<TableBodyCell tdClass="">
						<Button
							size="sm"
							on:click={() => {
								show_update_modal = true;
								current_update = slug.id;
							}}
						>
							<PenSolid size="sm" />
						</Button>
						<Button
							size="sm"
							on:click={() => {
								show_delete_modal = true;
								current_delete = slug.id;
							}}><TrashBinSolid size="sm" /></Button
						>
					</TableBodyCell>
					<TableBodyCell tdClass="">
						<Button size="sm" href={'/' + slug.slug} target="_blank">
							Testar
						</Button>
					</TableBodyCell>
				</TableBodyRow>
			{/each}
		</TableBody>
	</Table>
</div>
<Modal bind:open={show_delete_modal} size="md" autoclose outsideclose>
	<div class="text-center">
		<ExclamationCircleOutline
			class="mx-auto mb-4 h-9 w-9 text-gray-400 dark:text-gray-200"
		/>
		<Heading class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
			Tem certeza que deseja remover esse Slug?
		</Heading>
		<div class="mx-auto my-2 columns-1">
			<P>Slug: {delete_slug?.slug}</P>
			<P>
				Endereço Real: {delete_slug?.redirect}
			</P>
		</div>
		<Button
			color="red"
			class="me-2 mt-2"
			on:click={async () => await slugs.remove_slug(current_delete)}
			>Confirmar</Button
		>
		<Button color="alternative">Cancelar</Button>
	</div>
</Modal>
<UpdateSlug {slug_base} bind:modal_update_open={show_update_modal} />
