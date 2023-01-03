import { goto } from '$app/navigation';
import cuid from 'cuid';

import { getProviderSettings } from '$lib/api/provider_settings';
import { Provider } from '$lib/model/provider';
import type { CustomForm } from '$lib/model/custom_form';
import { ProviderSettingsState } from '$lib/model/provider_settings';

export class ProviderForm {
    ptype: string = "";
    state: number = 0;
    providerId: string = "";
    form: CustomForm|null = null;
    value: ProviderSettingsState = new ProviderSettingsState();
    nextInProgress: boolean = false;
    previousInProgress: boolean = false;
    on_previous: null | (() => void);
    on_done: () => void;

    constructor(ptype: string, on_done: () => void, providerId: string | null = null, value: ProviderSettingsState | null = null, on_previous: null | (() => void) = null) {
        this.ptype = ptype;
        this.state = -1;
        this.providerId = providerId?providerId:cuid();
        this.form = null;
        this.value = value?value:new ProviderSettingsState({recall: this.providerId, state: this.state});
        this.on_done = on_done;
        this.on_previous = on_previous;

        if ((!this.value.Provider || !Object.keys(this.value.Provider).length)) {
            const sstore = sessionStorage.getItem("newprovider-" + this.providerId);
            if (sstore) {
                const data = JSON.parse(sstore);
                if (data) {
                    if (data._id) this.value._id = data._id;
                    this.value._comment = data._comment;
                    this.value.Provider = data.Provider;
                }
            }
        }
        if (!this.value.recall) {
            this.value.recall = this.providerId;
        }

        this.nextInProgress = false;
        this.previousInProgress = false;
    }

    async changeState(toState: number): Promise<CustomForm | null> {
        if (toState == -1) {
            this.state = toState;
            if (this.on_previous) this.on_previous();
            return null;
        } else {
            try {
                const res = await getProviderSettings(this.ptype, toState, this.value);
                this.state = toState;
                if (res.values) {
                    // @ts-ignore
                    this.value.Provider = new Provider({ ...this.value.Provider, ...res.values });
                }
                return res.form;
            } catch (e) {
                if (e instanceof Provider) {
                    sessionStorage.removeItem("newprovider-" + this.providerId);
                    this.on_done();
                    return null;
                } else {
                    this.nextInProgress = false;
                    this.previousInProgress = false;
                    throw e;
                }
            }
        }
    }

    saveState() {
        sessionStorage.setItem("newprovider-" + this.providerId, JSON.stringify(this.value))
    }

    async nextState() {
        this.nextInProgress = true;
        this.saveState();
        if (this.form && this.form.nextButtonLink) {
            goto(this.form.nextButtonLink);
        } else {
            this.form = await this.changeState(this.form && this.form.nextButtonState ? this.form.nextButtonState : 0);
        }
        this.nextInProgress = false;
    }

    async previousState() {
        this.previousInProgress = true;
        this.saveState();
        if (this.form && this.form.previousButtonLink) {
            goto(this.form.previousButtonLink);
        } else {
            this.form = await this.changeState(this.form && this.form.previousButtonState ? this.form.previousButtonState : 0);
        }
        this.previousInProgress = false;
    }

}
