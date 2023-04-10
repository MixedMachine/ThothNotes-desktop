<script lang="ts">
    import Drawer, {
        AppContent,
        Content,
    } from "@smui/drawer";
    import List, { Item, Text, Separator, Meta } from "@smui/list";
    import IconButton from "@smui/icon-button";
    import NewNote from "./NewNote.svelte";
    import NotesPage from "./NotesPage.svelte";
	import { createEventDispatcher } from 'svelte';
    import type { backend } from "../../wailsjs/go/models";

    export let active;
    export let notes;

    let open = true;
    let dispatch = createEventDispatcher();

    function setActive(value: backend.Note) {
        dispatch ("activate", {
            page: value
        });
    }

    function newNote() {
        dispatch ("new");
    }

    function onSaved() {
        dispatch ("saved");
    }

    function onUpdate(event: CustomEvent<any>) {
        dispatch ("update", {
            note: event.detail.note
        });
    }

    function onDelete(id: number) {
        dispatch ("delete", {
            id: id
        });
    }

</script>

<div class="drawer-container">
    <Drawer variant="dismissible" bind:open>
        <Content>
            <List>
                <Item
                    href="javascript:void(0)"
                    on:click={() => newNote()}
                    activated={active.id === -1}
                >
                    <Text>New Note</Text>
                </Item>
                <Separator />
                {#each notes as item}
                    <Item
                        href="javascript:void(0)"
                        on:click={() => setActive(item)}
                        activated={active === item}
                    >
                        <Text>{item.title}</Text>
                        <Meta>
                            <IconButton class="material-icons" on:click={() => onDelete(item.id)}
                                >delete</IconButton
                              >
                            </Meta>
                    </Item>
                {/each}
            </List>
        </Content>
    </Drawer>

    <AppContent class="app-content">
        <main class="main-content">
            {#if active.title === "New Note"}
                <NewNote on:saved={() => onSaved()} />
            {:else}
                <NotesPage bind:note={active} on:update={(event) => onUpdate(event)} />
            {/if}
        </main>
    </AppContent>
</div>

<style>

</style>
