<div class="solo-demo-container solo-container">
        <Icon class="material-icons">search</Icon>
    <Paper class="solo-paper" elevation={6}>
        <Input
            bind:value
            on:keydown={handleKeyDown}
            placeholder="Search"
            class="solo-input"
            />
        <Icon
            class="material-icons"
            on:click={() => value=""}
            style="cursor: pointer; margin-right: 12px;"
        >close</Icon>
    </Paper>
    <Fab
        on:click={doSearch}
        color="primary"
        mini
        class="solo-fab"
        >
        <Icon class="material-icons">arrow_forward</Icon>
    </Fab>
</div>

  <script lang="ts">
import {
    Input
} from '@smui/textfield';
import Paper from '@smui/paper';
import Fab from '@smui/fab';
import {
    Icon
} from '@smui/common';
import {
    createEventDispatcher
} from 'svelte';

let value = '';

const dispatch = createEventDispatcher();

function doSearch() {
    dispatch('search', {
        search: value
    });
}

function handleKeyDown(event: CustomEvent | KeyboardEvent) {
    event = event as KeyboardEvent;
    if (event.key === 'Enter') {
        doSearch();
    }
}
</script>

<style>
.solo-demo-container {
    background-color: var(--mdc-theme-background, #212125);
    border: 1px solid var(--mdc-theme-text-hint-on-background, rgba(0, 0, 0, 0.1));
}

.solo-container {
    display: flex;
    justify-content: center;
    align-items: center;
}

* :global(.solo-paper) {
    background-color: var(--mdc-theme-surface, #333);
    color: var(--mdc-theme-on-surface, #fff);
    display: flex;
    align-items: center;
    flex-grow: 1;
    max-width: 600px;
    margin: 2px;
    padding: 0;
    height: 36px;
    border-radius: 24px;
}

* :global(.solo-paper > *) {
    display: inline-block;
    margin: 0 12px;
}

* :global(.solo-input) {
    flex-grow: 1;
    color: var(--mdc-theme-on-surface, #000);
}

* :global(.solo-input::placeholder) {
    color: var(--mdc-theme-on-surface, #000);
    opacity: 0.6;
}

* :global(.solo-fab) {
    flex-shrink: 0;
}
</style>
