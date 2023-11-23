import { error } from '@sveltejs/kit';
import { get_store_value } from 'svelte/internal';
import type { Load } from '@sveltejs/kit';

import { domains, domains_idx, refreshDomains } from '$lib/stores/domains';

export const load: Load = async({ parent, params }) => {
    const data = await parent();

    if (!get_store_value(domains)) await refreshDomains();

    const domain: DomainInList | null = get_store_value(domains_idx)[params.dn];

    if (!domain) {
        throw error(404, {
	    message: 'Domain not found'
	});
    }

    let historyid = undefined;
    if (domain.zone_history && domain.zone_history.length > 0) {
        historyid = domain.zone_history[0];
    }

    return {
        domain,
        history: historyid,
        ...data,
    }
}
