import { writable } from "svelte/store";

export type Slug = {
    id: string;
    slug: string;
    redirect: string;
};

const create_slugs_store = () => {
    const { update, subscribe, set } = writable<Slug[]>([{
        id: '1',
        slug: 'apple-macbook-pro',
        redirect: '/products/apple-macbook-pro'
    },
    {
        id: '2',
        slug: 'microsoft-surface-pro',
        redirect: '/products/microsoft-surface-pro'
    },
    {
        id: '3',
        slug: 'magic-mouse-2',
        redirect: '/products/magic-mouse-2'
    }]);

    const add_slug = async (slug: Slug) => {
        await new Promise(() => {
            update((old) => {
                return [...old, slug];
            });
        });
    };

    const update_slug = async (id: string, slug: Slug) => {
        await new Promise(() => {
            update((old) => {
                return old.map((s) => s.id === id ? slug : s);
            });
        });
    };

    const remove_slug = async (id: string) => {
        await new Promise(() => {
            update((old) => {
                return old.filter((s) => s.id === id);
            });
        });
    };

    return {
        set, update, subscribe,
        add_slug,
        update_slug,
        remove_slug,
    };
};

export const slugs = create_slugs_store();
