/** @type {import('./$types').LayoutData} */

import { browser } from '$app/environment';
import { auth, initializeFirebase } from '$lib/auth';
import { onAuthStateChanged } from 'firebase/auth';
import { injectSpeedInsights } from '@vercel/speed-insights/sveltekit';

injectSpeedInsights();

export async function load({ url }: {url:{pathname:string}}) {
    if (browser) {
        try {
            initializeFirebase();
        } catch (ex) {
            console.error(ex);
        }
    }

    function getAuthUser() {
        return new Promise((resolve) => {
            onAuthStateChanged(auth, (user) => resolve(user ? user : false));
        });
    }

    return {
        getAuthUser: getAuthUser,
        url: url.pathname
    };
}