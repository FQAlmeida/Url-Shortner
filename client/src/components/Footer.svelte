<script lang="ts">
	import {
		Footer,
		FooterCopyright,
		FooterLinkGroup,
		FooterBrand,
		FooterLink
	} from 'flowbite-svelte';
</script>

<Footer footerType="logo">
	<div class="sm:flex sm:items-center sm:justify-between">
		<FooterBrand
			href="https://flowbite.com"
			src="https://flowbite-svelte.com/images/flowbite-svelte-icon-logo.svg"
			alt="Flowbite Logo"
			name="URL Shortner"
		/>
		<FooterLinkGroup
			ulClass="flex flex-wrap items-center mb-6 text-sm text-gray-500 sm:mb-0 dark:text-gray-400"
		>
			<FooterLink href="/">About</FooterLink>
			<FooterLink href="/">Privacy Policy</FooterLink>
			<FooterLink href="/">Licensing</FooterLink>
			<FooterLink href="/">Contact</FooterLink>
		</FooterLinkGroup>
	</div>
	<hr class="my-6 border-gray-200 sm:mx-auto lg:my-8 dark:border-gray-700" />
	<FooterCopyright href="/" by="URL Shortner" />
</Footer>
