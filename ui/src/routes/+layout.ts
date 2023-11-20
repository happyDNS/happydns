import type { Load } from '@sveltejs/kit';
import { get_store_value } from 'svelte/internal';

import { toasts } from '$lib/stores/toasts';
import { refreshUserSession } from '$lib/stores/usersession';
import { config as tsConfig, locale, loadTranslations, t } from '$lib/translations';

export const ssr = false;

const sw_state = { triedUpdate: false, hasUpdate: false };

function onSWupdate(sw_state: {hasUpdate: boolean}, installingWorker: ServiceWorker) {
    if (!sw_state.hasUpdate) {
        toasts.addToast({
            title: get_store_value(t)('upgrade.title'),
            message: get_store_value(t)('upgrade.content'),
            onclick: () => installingWorker.postMessage('SKIP_WAITING'),
        });
    }
    sw_state.hasUpdate = true;
}

export const load: Load = async({ fetch, route, url }) => {
    const { MODE } = import.meta.env;

    const initLocale = locale.get() || window.navigator.language || window.navigator.languages[0] || tsConfig.fallbackLocale || "en";

    await loadTranslations(initLocale, url.pathname);

    if (MODE == 'production' && 'serviceWorker' in navigator) {
        navigator.serviceWorker.ready.then((registration) => {
            registration.onupdatefound = () => {
                const installingWorker = registration.installing;

                if (installingWorker === null) return;

                installingWorker.onstatechange = () => {
                    if (installingWorker.state === 'installed') {
                        if (navigator.serviceWorker.controller) {
                            onSWupdate(sw_state, installingWorker);
                        }
                    }
                }
            }

            if (!sw_state.triedUpdate) {
                sw_state.triedUpdate = true;
                registration.update();
                setInterval(function (reg) { reg.update() }, 36000000, registration);
            }
        });

        let refreshing = false;
        navigator.serviceWorker.addEventListener('controllerchange', () => {
            if (!refreshing) {
                window.location.reload();
                refreshing = true;
            }
        });
    }

    // Load user session if any
    try {
        const user = await refreshUserSession(fetch);
        if (get_store_value(locale) != user.settings.language) {
            locale.set(user.settings.language);
        }
    } catch {}

    return {
        route,
        sw_state,
    };
}
