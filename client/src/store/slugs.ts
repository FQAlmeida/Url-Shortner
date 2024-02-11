import { get, writable } from "svelte/store";
import { session } from "./auth";

export type Slug = {
    id: string;
    slug: string;
    redirect: string;
};

const create_slugs_store = async () => {
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

    const reset_slugs = async () => {
        const uri = new URL("http://localhost:8080/slugs");
        uri.searchParams.append("userid", get(session).user?.uid ?? "-1");
        console.log(get(session));
        try {
            const response = await fetch(uri);
            const slugs: Slug[] = await response.json();
            set(slugs);
        } catch (e) {
            console.error(e);
        }
    };

    session.subscribe((s) => {
        if (!!s.user) {
            reset_slugs();
        }
    });

    const add_slug = async (slug: Omit<Slug, "id">) => {
        const uri = new URL("http://localhost:8080/slugs");
        uri.searchParams.append("uid", get(session).user?.uid ?? "-1");
        const response = await fetch(uri, {
            method: "POST", // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // no-cors, *cors, same-origin
            headers: {
                "Content-Type": "application/json",
                // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            redirect: "follow", // manual, *follow, error
            referrerPolicy: "no-referrer", // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
            body: JSON.stringify({ ...slug, uid: get(session).user?.uid ?? "-1" }), // body data type must match "Content-Type" header
        });
        const slug_set: Slug = await response.json();
        update((old) => {
            return [...old, slug_set];
        });
    };

    const update_slug = async (slug: Slug) => {
        const uri = new URL("http://localhost:8080/slugs");
        const _ = await fetch(uri, {
            method: "PUT", // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // no-cors, *cors, same-origin
            headers: {
                "Content-Type": "application/json",
                // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            redirect: "follow", // manual, *follow, error
            referrerPolicy: "no-referrer", // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
            body: JSON.stringify({ ...slug, uid: get(session).user?.uid ?? "-1" }), // body data type must match "Content-Type" header
        });
        update((old) => {
            return old.map((s) => s.id === slug.id ? slug : s);
        });
    };

    const remove_slug = async (id: string) => {
        const uri = new URL("http://localhost:8080/slugs");
        uri.searchParams.append("uid", get(session).user?.uid ?? "-1");
        uri.searchParams.append("id", id);
        const _ = await fetch(uri, {
            method: "DELETE", // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // no-cors, *cors, same-origin
            headers: {
                "Content-Type": "application/json",
                // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            redirect: "follow", // manual, *follow, error
            referrerPolicy: "no-referrer", // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
        });
        update((old) => {
            return old.filter((slug) => slug.id !== id);
        });
    };
    await reset_slugs();
    return {
        set, update, subscribe,
        add_slug,
        update_slug,
        remove_slug,
        reset_slugs
    };
};

export const slugs = await create_slugs_store();
