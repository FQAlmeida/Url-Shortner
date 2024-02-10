<script lang="ts">
	import { slugs } from '@store/slugs';
	import {
		Button,
		Heading,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import { PenSolid, TrashBinSolid } from 'flowbite-svelte-icons';
	function handle_delete(id: string) {
		return async () => {
			slugs.remove_slug(id);
		};
	}
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
						<Button size="sm">
							<PenSolid size="sm" />
						</Button>
						<Button size="sm"><TrashBinSolid size="sm" /></Button>
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
