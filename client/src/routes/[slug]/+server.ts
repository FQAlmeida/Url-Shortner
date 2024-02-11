import type { Slug } from '@store/slugs';
import { redirect } from '@sveltejs/kit';

/** @type {import('@sveltejs/kit').RequestHandler<{	slug: string }>} */
export async function GET({ params }) {
    const slug_str = params.slug;
    const uri = new URL("http://localhost:8080/slug");
    uri.searchParams.append("slug", slug_str);
    const response = await fetch(uri);
    const slug_response: Slug = await response.json();
    redirect(301, slug_response.redirect);
}
